package ucdiscord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"strings"
	"time"

	"github.com/Implex-ltd/cleanhttp/cleanhttp"
	http "github.com/bogdanfinn/fhttp"
	"github.com/google/uuid"
)

var (
	dUrl, _ = url.Parse("https://discord.com")
)

func NewClient(config *Config) (*Client, error) {
	if config.ApiVersion != 0 {
		VERSION = config.ApiVersion
	}

	C := &Client{
		Config:    config,
		Ws:        config.Ws,
		UserAgent: cleanhttp.ParseUserAgent(config.Http.Config.BrowserFp.Navigator.UserAgent),
	}

	if config.GetCookies {
		if _, _, err := C.GetCookies(); err != nil {
			return nil, err
		}
	}

	return C, nil
}

func (C *Client) GetCookies() (resp *Response, data *FingerprintResponse, err error) {
	token := C.Config.Token
	C.Config.Token = ""

	defer func() {
		C.Config.Token = token
	}()

	resp, err = C.Do(Request{
		Endpoint: fmt.Sprintf("%s/experiments", ENDPOINT),
		Method:   "GET",
		Header: C.GetHeader(&HeaderConfig{
			Info: &PropInfo{
				Type: PROP_TRACK,
			},
		}),
		Response: &data,
	})

	if err != nil {
		return nil, nil, err
	}

	if data == nil {
		return nil, nil, fmt.Errorf("cant scrape x-fingerprint (data nil)")
	}

	if data.Fingerprint == "" {
		return nil, nil, fmt.Errorf("cant scrape x-fingerprint (fp not found)")
	}

	C.XFingerprint = data.Fingerprint

	C.Config.Http.Client.SetCookies(dUrl, []*http.Cookie{{
		Name:  "locale",
		Value: strings.Split(C.Config.Http.Config.BrowserFp.Navigator.Language, "-")[0],
	}})

	return resp, data, err
}

func (C *Client) Register(config *Config) (resp *Response, data *RegisterResponse, err error) {
	if config.Username == "" || config.CaptchaKey == "" || config.Invite == "" {
		return nil, nil, fmt.Errorf("invalid params")
	}

	resp, err = C.Do(Request{
		Endpoint: fmt.Sprintf("%s/auth/register", ENDPOINT),
		Method:   "POST",
		Body: &Register{
			Username:    config.Username,
			Fingerprint: C.XFingerprint,
			Invite:      config.Invite,
			Consent:     true,
		},
		Header: C.GetHeader(&HeaderConfig{
			Referer:     fmt.Sprintf(`/invite/%s`, config.Invite),
			CaptchaKey:  config.CaptchaKey,
			Fingerprint: C.XFingerprint,
			Info: &PropInfo{
				Type: PROP_SUPER,
			},
		}),
		Response: &data,
	})

	if data == nil {
		return nil, nil, fmt.Errorf("/register data is nil")
	}

	if data.Token != "" {
		C.Config.Ws.Token = data.Token
		C.Config.Token = data.Token

		if C.Ws != nil {
			C.Ws.Token = data.Token
		}
	}

	return resp, data, err
}

func (C *Client) JoinGuild(config *Config) (resp *Response, data *JoinServerResponse, err error) {
	if config.Invite == "" || config.GuildID == "" || config.ChannelD == "" || C.Ws.ReadyData.SessionID == "" {
		return nil, nil, fmt.Errorf("invalid params or websocket disconnected")
	}

	resp, err = C.Do(Request{
		Endpoint: fmt.Sprintf("%s/invites/%s", ENDPOINT, config.Invite),
		Method:   "POST",
		Body: &JoinPayload{
			SessionID: C.Ws.ReadyData.SessionID,
		},
		Header: C.GetHeader(&HeaderConfig{
			Referer:   `/channels/@me`,
			GuildID:   config.GuildID,
			ChannelID: config.ChannelD,
			Join:      true,
			Info: &PropInfo{
				Type: PROP_SUPER,
			},
		}),
		Response: &data,
	})

	return resp, data, err
}

func (C *Client) PatchUser(config *Config) (resp *Response, data *SetBirthResponse, err error) {
	if config.Username == "" && config.DisplayName == "" && config.Date == "" && config.Email == "" && config.Password == "" && config.Avatar == "" {
		return nil, nil, fmt.Errorf("invalid params")
	}

	if !config.AvatarFromB64 && config.Avatar != "" {
		if config.Avatar, err = cleanhttp.ImageToBase64(config.Avatar); err != nil {
			return nil, nil, err
		}
	}

	resp, err = C.Do(Request{
		Endpoint: fmt.Sprintf("%s/users/@me", ENDPOINT),
		Method:   "PATCH",
		Body: CleanPatchUser(&PatchPayload{
			Date:        config.Date,
			Email:       config.Email,
			Password:    config.Password,
			Avatar:      config.Avatar,
			DisplayName: config.DisplayName,
			Username:    config.Username,
		}),
		Header: C.GetHeader(&HeaderConfig{
			Referer: `/channels/@me`,
			Info: &PropInfo{
				Type: PROP_SUPER,
			},
		}),
		Response: &data,
	})

	return resp, data, err
}

func (C *Client) PatchProfil(config *Config) (resp *Response, data *SetProfilResponse, err error) {
	if config.Bio == "" && config.Pronouns == "" && config.AccentColor == 0 {
		return nil, nil, fmt.Errorf("invalid params")
	}

	resp, err = C.Do(Request{
		Endpoint: fmt.Sprintf("%s/users/@me", ENDPOINT),
		Method:   "PATCH",
		Body: CleanSetProfil(&SetProfilPayload{
			AccentColor: config.AccentColor,
			Bio:         config.Bio,
			Pronouns:    config.Pronouns,
		}),
		Header: C.GetHeader(&HeaderConfig{
			Referer: `/channels/@me`,
			Info: &PropInfo{
				Type: PROP_SUPER,
			},
		}),
		Response: &data,
	})

	return resp, data, err
}

func (C *Client) JoinHypesquad(House int) (resp *Response, err error) {
	resp, err = C.Do(Request{
		Endpoint: fmt.Sprintf("%s/hypesquad/online", ENDPOINT),
		Method:   "POST",
		Body: &HypesquadPayload{
			House: House,
		},
		Header: C.GetHeader(&HeaderConfig{
			Referer: `/channels/@me`,
			Info: &PropInfo{
				Type: PROP_SUPER,
			},
		}),
	})

	return resp, err
}

func (C *Client) VerifyEmail(jwt, captcha string) (resp *Response, data *VerifyResponse, err error) {
	if jwt == "" {
		return nil, nil, fmt.Errorf("invalid params")
	}

	resp, err = C.Do(Request{
		Endpoint: fmt.Sprintf("%s/auth/verify", ENDPOINT),
		Method:   "POST",
		Body: &VerifyEmailPayload{
			Jwt: jwt,
		},
		Header: C.GetHeader(&HeaderConfig{
			Fingerprint: C.XFingerprint,
			Referer:     `/verify`,
			Info: &PropInfo{
				Type: PROP_SUPER,
			},
			CaptchaKey: captcha,
		}),
		Response: &data,
	})

	if err != nil {
		return nil, nil, err
	}

	if data.Token != "" {
		C.Config.Ws.Token = data.Token
		C.Config.Token = data.Token

		if C.Ws != nil {
			C.Ws.Token = data.Token
		}
	}

	return resp, data, err
}

func (C *Client) SupressTutorial() (resp *Response, err error) {
	resp, err = C.Do(Request{
		Endpoint: fmt.Sprintf("%s/tutorial/indicators/suppress", ENDPOINT),
		Method:   "POST",
		Header: C.GetHeader(&HeaderConfig{
			Referer: `/channels/@me`,
			Info: &PropInfo{
				Type: PROP_SUPER,
			},
		}),
	})

	return resp, err
}

func (C *Client) AddFriend(config *Config) (resp *Response, data *CaptchaSpawnResponse, err error) {
	if config.Username == "" || C.Ws.ReadyData.SessionID == "" {
		return nil, nil, fmt.Errorf("invalid params or websocket disconnected")
	}

	C.Do(Request{
		Endpoint: fmt.Sprintf("%s/science", ENDPOINT),
		Method:   "POST",
		Body: &FriendScience{
			Token: C.Ws.ReadyData.AnalyticsToken,
			Events: []ScEvent{
				{
					Type: "app_ui_viewed",
					Properties: ScProperties{
						ClientTrackTimestamp:      time.Now().UnixNano() / int64(time.Millisecond),
						ClientHeartbeatSessionID:  C.Ws.ReadyData.SessionID,
						TotalCompressedByteSize:   6613681,
						TotalUncompressedByteSize: 23571105,
						TotalTransferByteSize:     0,
						JSCompressedByteSize:      5399002,
						JSUncompressedByteSize:    20359840,
						JSTransferByteSize:        0,
						CSSCompressedByteSize:     484478,
						CSSUncompressedByteSize:   2398802,
						CSSTransferByteSize:       0,
						LoadID:                    uuid.NewString(),
						ScreenName:                "friends",
						DurationMSSinceAppOpened:  882,
						ClientPerformanceMemory:   0,
						AccessibilityFeatures:     524544,
						RenderedLocale:            "fr",
						UptimeApp:                 1,
						ClientRtcState:            "DISCONNECTED",
						ClientAppState:            "focused",
						ClientSendTimestamp:       time.Now().UnixNano() / int64(time.Millisecond),
					},
				},
				{
					Type: "ready_payload_received",
					Properties: ScProperties{
						ClientTrackTimestamp:           time.Now().UnixNano() / int64(time.Millisecond),
						ClientHeartbeatSessionID:       C.Ws.ReadyData.SessionID,
						CompressedByteSize:             14034,
						UncompressedByteSize:           51047,
						CompressionAlgorithm:           "zlib-stream",
						PackingAlgorithm:               "json",
						UnpackDurationMS:               0,
						IdentifyTotalServerDurationMS:  206,
						IdentifyAPIDurationMS:          140,
						IdentifyGuildsDurationMS:       1,
						NumGuilds:                      1,
						NumGuildChannels:               5,
						NumGuildCategoryChannels:       2,
						PresencesSize:                  2,
						UsersSize:                      248,
						ReadStatesSize:                 187,
						PrivateChannelsSize:            2,
						UserSettingsSize:               18,
						ExperimentsSize:                39916,
						UserGuildSettingsSize:          42,
						RelationshipsSize:              2,
						RemainingDataSize:              2947,
						GuildChannelsSize:              939,
						GuildMembersSize:               230,
						GuildPresencesSize:             2,
						GuildRolesSize:                 209,
						GuildEmojisSize:                4,
						GuildThreadsSize:               4,
						GuildStickersSize:              4,
						GuildEventsSize:                4,
						GuildFeaturesSize:              4,
						GuildRemainingDataSize:         3124,
						SizeMetricsDurationMS:          0,
						DurationMSSinceIdentifyStart:   305,
						DurationMSSinceConnectionStart: 306,
						DurationMSSinceEmitStart:       22,
						IsReconnect:                    false,
						IsFastConnect:                  false,
						DidForceClearGuildHashes:       false,
						IdentifyUncompressedByteSize:   832,
						IdentifyCompressedByteSize:     481,
						HadCacheAtStartup:              false,
						UsedCacheAtStartup:             false,
						AccessibilityFeatures:          524544,
						RenderedLocale:                 "fr",
						UptimeApp:                      1,
						ClientRtcState:                 "DISCONNECTED",
						ClientAppState:                 "focused",
						//AccessibilitySupportEnabled:    false,
						ClientUUID:          C.Ws.ReadyData.AuthSessionIDHash,
						ClientSendTimestamp: time.Now().UnixNano() / int64(time.Millisecond),
					},
				},
				{
					Type: "friends_list_viewed",
					Properties: ScProperties{
						ClientTrackTimestamp:     time.Now().UnixNano() / int64(time.Millisecond),
						ClientHeartbeatSessionID: C.Ws.ReadyData.SessionID,
						TabOpened:                "ADD_FRIEND",
						ClientPerformanceMemory:  0,
						AccessibilityFeatures:    524544,
						RenderedLocale:           "fr",
						UptimeApp:                1,
						ClientRtcState:           "DISCONNECTED",
						ClientAppState:           "focused",
						//AccessibilitySupportEnabled: false,
						ClientUUID:          C.Ws.ReadyData.AuthSessionIDHash,
						ClientSendTimestamp: time.Now().UnixNano() / int64(time.Millisecond),
					},
				},
			},
		},
		Header: C.GetHeader(&HeaderConfig{
			Referer: `/channels/@me`,
			Friend:  true,
			Info: &PropInfo{
				Type: PROP_SUPER,
			},
		}),
	})

	resp, err = C.Do(Request{
		Endpoint: fmt.Sprintf("%s/users/@me/relationships", ENDPOINT),
		Method:   "POST",
		Body: &AddFriendPayload{
			Username: config.Username,
		},
		Header: C.GetHeader(&HeaderConfig{
			Referer: `/channels/@me`,
			Friend:  true,
			Info: &PropInfo{
				Type: PROP_SUPER,
			},
			CaptchaKey:     config.CaptchaKey,
			CaptchaRqtoken: config.RqToken,
		}),
		Response: &data,
	})

	return resp, data, err
}

func (C *Client) IsLocked() (locked bool, status int, err error) {
	if C.Config.Token == "" {
		return false, 0, fmt.Errorf("token is missing")
	}

	var data LockedResponse
	resp, err := C.Do(Request{
		Endpoint: fmt.Sprintf("%s/users/@me/burst-credits", ENDPOINT),
		Method:   "GET",
		Header: C.GetHeader(&HeaderConfig{
			Referer: `/channels/@me`,
			Info: &PropInfo{
				Type: PROP_SUPER,
			},
		}),
		Response: &data,
	})

	if err != nil {
		return false, 0, err
	}

	if (resp.Status == 403 && data.Code == 40002) || resp.Status == 401 {
		return true, resp.Status, nil
	}

	return false, resp.Status, nil
}

func (C *Client) SendMessage(message string, tts bool, ChannelID string) (resp *Response, data *SendMessageResponse, err error) {
	if C.Config.Token == "" {
		return nil, nil, fmt.Errorf("token is missing. ")
	}
	if message == "" || ChannelID == "" {
		return nil, nil, fmt.Errorf("invalid params")
	}

	resp, err = C.Do(Request{
		Endpoint: fmt.Sprintf(`%s/channels/%s/messages`, ENDPOINT, ChannelID),
		Method:   "POST",
		Body: &SendMessagePayload{
			MobileNetworkType: "unknown",
			Content:           message,
			Nonce:             fmt.Sprint(Snowflake()),
			Tts:               tts,
			Flags:             0,
		},
		Header: C.GetHeader(&HeaderConfig{
			Referer: fmt.Sprintf(`%s/channels/%s/messages`, ENDPOINT, ChannelID),

			Info: &PropInfo{
				Type: PROP_SUPER,
			},
		}),
		Response: &data,
	})
	if err != nil {
		return nil, nil, err
	}

	return resp, data, nil
}

func (C *Client) CaptchaEvent(siteKey string) (resp *Response, err error) {
	if siteKey == "" || C.Ws.ReadyData.SessionID == "" {
		return nil, fmt.Errorf("invalid params or websocket disconnected")
	}

	resp, err = C.Do(Request{
		Endpoint: fmt.Sprintf("%s/science", ENDPOINT),
		Method:   "POST",
		Body: &FriendScience{
			Token: C.Ws.ReadyData.AnalyticsToken,
			Events: []ScEvent{
				{
					Type: "captcha_event",
					Properties: ScProperties{
						ClientTrackTimestamp:        time.Now().UnixNano() / int64(time.Millisecond),
						CaptchaEventName:            "verify",
						CaptchaService:              "hcaptcha",
						SiteKey:                     siteKey,
						CaptchaFlowKey:              "814d5ca3-0b6a-4ac5-b350-403ff03bf8af",
						ClientPerformanceMemory:     0,
						AccessibilityFeatures:       256,
						RenderedLocale:              "fr",
						AccessibilitySupportEnabled: false,
						ClientUUID:                  C.Ws.ReadyData.AuthSessionIDHash,
						ClientSendTimestamp:         time.Now().UnixNano() / int64(time.Millisecond),
					},
				},
			},
		},
		Header: C.GetHeader(&HeaderConfig{
			Info: &PropInfo{
				Type: PROP_SUPER,
			},
		}),
	})

	return resp, err
}

/*
	{
		"message": "Corps de formulaire non valide",
		"code": 50035,
		"errors": {
			"_errors": [
				{
					"code": "DICT_TYPE_CONVERT",
					"message": "Seuls les dictionnaires peuvent \u00eatre utilis\u00e9s dans un DictType"
				}
			]
		}
	}
*/
func (C *Client) SendInteraction(config *Config) (resp *Response, err error) {
	if C.Config.Token == "" {
		return nil, fmt.Errorf("token is missing. ")
	}

	if config.Data == "" || config.Type == 0 || config.ChannelD == "" || config.GuildID == "" || config.ApplicationID == "" || config.AnalyticsLocation == "" {
		return nil, fmt.Errorf("invalid params")
	}

	data := json.RawMessage(config.Data)

	fullPayload, err := json.Marshal(&InteractionPayload{
		Nonce:             fmt.Sprint(Snowflake()),
		Type:              config.Type,
		ApplicationID:     config.ApplicationID,
		GuildID:           config.GuildID,
		ChannelID:         config.ChannelD,
		SessionID:         C.Ws.ReadyData.SessionID,
		AnalyticsLocation: config.AnalyticsLocation,
		Data:              data,
	})
	if err != nil {
		return nil, err
	}

	var bodyBuffer bytes.Buffer
	writer := multipart.NewWriter(&bodyBuffer)
	//writer.SetBoundary("WebKitFormBoundaryH5OPcFuDI6Q2nXzF")

	jsonField, err := writer.CreateFormField("payload_json")
	if err != nil {
		return nil, err
	}

	io.WriteString(jsonField, string(fullPayload))
	writer.Close()

	//fmt.Println(bodyBuffer.String())

	resp, err = C.Do(Request{
		Endpoint: fmt.Sprintf(`%s/interactions`, ENDPOINT),
		Method:   "POST",
		Body:     bodyBuffer.Bytes(),
		Header: C.GetHeader(&HeaderConfig{
			Referer: fmt.Sprintf(`%s/channels/%s/%s`, ENDPOINT, config.GuildID, config.ChannelD),
			Info: &PropInfo{
				Type: PROP_SUPER,
			},
			ContentType: writer.FormDataContentType(),
		}),
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
