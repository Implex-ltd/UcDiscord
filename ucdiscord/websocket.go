package ucdiscord

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/dgrr/fastws"
)

func NewWebsocket(Token string, Prop *XProp) (*ClientWebsocket, error) {
	if Token == "" {
		return nil, fmt.Errorf("token invalid")
	}

	return &ClientWebsocket{
		Token: Token,
		Open:  true,
		Prop:  Prop,
	}, nil
}

func (C *ClientWebsocket) Login() (err error) {
	C.Conn, err = fastws.Dial(fmt.Sprintf("wss://gateway.discord.gg/?encoding=json&v=%d", VERSION))
	if err != nil {
		return err
	}

	var buff []byte
	_, buff, err = C.Conn.ReadMessage(buff)
	if err != nil {
		return err
	}

	var Op OpLoginResponse
	if err := json.Unmarshal(buff, &Op); err != nil {
		return err
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

	if _, err = C.Conn.Write(payload); err != nil {
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

		if _, err := C.Conn.WriteString(`{"op": 1, "d": 3}`); err != nil {
			C.Open = false
			return
		}
	}
}

func (C *ClientWebsocket) Listen() {
	for C.Open {
		var buff []byte

		_, buff, err := C.Conn.ReadMessage(buff)
		if err != nil {
			C.Open = false
			if C.Debug {
				log.Println(err, string(buff))
			}
			return
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
				C.Open = false
				if C.Debug {
					log.Println(err, string(buff))
				}
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

	if _, err = C.Conn.Write(payload); err != nil {
		return err
	}

	return nil
}

func (C *ClientWebsocket) Close() {
	defer C.Conn.Close()
	C.Open = false
}