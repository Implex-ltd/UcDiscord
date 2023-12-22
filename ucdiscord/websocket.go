package ucdiscord

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

func NewWebsocket(Token string, Prop *XProp) (*ClientWebsocket, error) {
	return &ClientWebsocket{
		Token: Token,
		Open:  true,
		Prop:  Prop,
	}, nil
}

func (C *ClientWebsocket) Login() (err error) {
	C.Conn, _, err = websocket.DefaultDialer.Dial(fmt.Sprintf("wss://gateway.discord.gg/?encoding=json&v=%d", VERSION), nil)
	if err != nil {
		return err
	}

	_, buff, err := C.Conn.ReadMessage()
	if err != nil {
		return err
	}

	var Op OpLoginResponse
	if err := json.Unmarshal(buff, &Op); err != nil {
		return err
	}

	if C.DebugRecvData && C.Debug {
		log.Println(string(buff))
	}

	go C.Heartbeat(Op.D.HeartbeatInterval)
	go C.Listen()

	Cap := 16381
	Comp := false

	payload, err := json.Marshal(&WsData{
		Op: 2,
		D: D{
			Token:        &C.Token,
			Capabilities: &Cap,
			Properties:   C.Prop,
			Presence: &Presence{
				Status:     "online",
				Since:      0,
				Activities: nil,
				Afk:        false,
			},
			Compress: &Comp,
			ClientState: &ClientState{
				HighestLastMessageID:     "0",
				ReadStateVersion:         0,
				UserGuildSettingsVersion: -1,
				UserSettingsVersion:      -1,
				PrivateChannelsVersion:   "0",
				APICodeVersion:           0,
			},
		},
	})
	if err != nil {
		return err
	}

	if err = C.Conn.WriteMessage(websocket.BinaryMessage, payload); err != nil {
		fmt.Println("cant write payload")
		return err
	}

	for !C.Ready {
		if !C.Open {
			return fmt.Errorf("cant login")
		}
		time.Sleep(time.Millisecond)
	}

	return nil
}

func (C *ClientWebsocket) Heartbeat(span int) {
	t := time.NewTicker(time.Millisecond * time.Duration(span))
	defer t.Stop()

	for range t.C {
		if !C.Open {
			return
		}

		if err := C.Conn.WriteMessage(websocket.BinaryMessage, []byte(`{"op": 1, "d": 3}`)); err != nil {
			if C.Debug {
				log.Println("cant sent heatbeat", err)
			}
			C.Open = false
			return
		}
	}
}

func (C *ClientWebsocket) Listen() {
	for C.Open {
		_, buff, err := C.Conn.ReadMessage()
		if err != nil {
			C.Open = false
			if C.Debug {
				log.Println("listen read:", err, string(buff))
			}
			return
		}

		if C.DebugRecvData && C.Debug {
			log.Println(string(buff))
		}

		var out TData
		if err := json.Unmarshal(buff, &out); err != nil {
			C.Open = false
			if C.Debug {
				log.Println(err, string(buff))
			}
			return
		}

		switch out.T {
		case "SESSIONS_REPLACE":
			continue
		case "READY_SUPPLEMENTAL":
			C.Ready = true
		case "READY":
			var out Reply
			if err := json.Unmarshal(buff, &out); err != nil {
				if C.Debug {
					log.Println("cant unmarshal ready payload:", err)
				}

				C.Open = false
				return
			}
			C.ReadyData = &out.D
		default:
			if C.LogNonImplemented {
				fmt.Println(string(buff))
			}
		}
	}
}

func (C *ClientWebsocket) UpdateStatus(Status, State string) error {
	payload, err := json.Marshal(&WsData{
		Op: 3,
		D: &Presence{
			Status: Status,
			Since:  0,
			Game: Game{
				Name: State,
				Type: 0,
			},
			/*Activities: []Activity{
				{
					Name:  "Custom Status",
					Type:  4,
					State: State,
					Timestamps: Timestamps{
						End: time.Now().Add(24 * time.Hour).Unix(),
					},
					Emoji: Emoji{
						Name: Emote,
					},
				},
			},*/
			Afk: false,
		},
	})
	if err != nil {
		return err
	}

	if err = C.Conn.WriteMessage(websocket.BinaryMessage, payload); err != nil {
		return err
	}

	return nil
}

func (C *ClientWebsocket) Close() {
	defer C.Conn.Close()
	C.Open = false
}
