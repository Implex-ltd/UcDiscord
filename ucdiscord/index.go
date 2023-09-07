package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

func NewClient(config *ClientConfig) (*Client, error) {
	c := Client{
		Config:      config,
		HttpClient:  config.Client,
		BuildNumber: config.BuildNumber,
		UaInfo:      cleanhttp.ParseUserAgent(config.Client.Config.BrowserFp.Navigator.UserAgent),
	}

	if config.GetCookies {
		if err := c.GetCookies(); err != nil {
			return nil, fmt.Errorf("error getting cookies: %v", err.Error())
		}
	}

	return &c, nil
}

func (c *Client) GetCookies() error {
	c.HttpClient.Client.SetCookies(dUrl, []*http.Cookie{{
		Name:  "locale",
		Value: strings.Split(c.HttpClient.Config.BrowserFp.Navigator.Language, "-")[0],
	}})

	response, err := c.HttpClient.Do(cleanhttp.RequestOption{
		Method: "GET",
		Url:    "https://discord.com/api/v9/experiments",
		Header: c.GetHeader(&HeaderConfig{}),
	})
	if err != nil {
		return fmt.Errorf("error making HTTP request: %v", err.Error())
	}

	defer response.Body.Close()

	resp, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var fp FingerprintResponse
	if err := json.Unmarshal([]byte(resp), &fp); err != nil {
		return fmt.Errorf("cant unmarshal fingerpint: %v", err.Error())
	}

	c.xfingerprint = fp.Fingerprint

	return nil
}

func (c *Client) JoinGuild(config *JoinConfig) (*JoinServerResponse, error) {
	if c.WsProperties.D.SessionID == "" {
		return nil, fmt.Errorf("please connect to the websocket first")
	}

	payload, err := json.Marshal(&JoinPayload{
		SessionID: c.WsProperties.D.SessionID,
	})
	if err != nil {
		return nil, fmt.Errorf("error marshaling payload: %v", err.Error())
	}

	header := c.GetHeader(&HeaderConfig{
		Join: config,
	})

	response, err := c.HttpClient.Do(cleanhttp.RequestOption{
		Method: "POST",
		Url:    fmt.Sprintf("https://discord.com/api/v9/invites/%s", config.InviteCode),
		Body:   bytes.NewReader(payload),
		Header: header,
	})
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %v", err.Error())
	}

	defer response.Body.Close()

	resp, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var data JoinServerResponse
	if err := json.Unmarshal([]byte(resp), &data); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %v", err.Error())
	}

	return &data, nil
}

func (c *Client) Register(config *RegisterConfig) (*RegisterResponse, error) {
	var pl any
	var header http.Header

	if config.InviteCode != "" {
		pl = RegisterInvitePayload{
			Fingerprint:                c.xfingerprint,
			GlobalName:                 config.Username,
			Invite:                     config.InviteCode,
			Consent:                    true,
			GiftCodeSkuID:              nil,
			UniqueUsernameRegistration: true,
		}

		header = c.GetHeader(&HeaderConfig{
			IsXtrack: false,
			IsSuper:  true,
		})

		header.Set("x-fingerprint", c.xfingerprint)

		if config.CustomProperties == "" {
			header.Del("x-context-properties")
			header.Add("x-super-properties", "eyJvcyI6IldpbmRvd3MiLCJicm93c2VyIjoiQ2hyb21lIiwiZGV2aWNlIjoiIiwic3lzdGVtX2xvY2FsZSI6ImZyLUZSIiwiYnJvd3Nlcl91c2VyX2FnZW50IjoiTW96aWxsYS81LjAgKFdpbmRvd3MgTlQgMTAuMDsgV2luNjQ7IHg2NCkgQXBwbGVXZWJLaXQvNTM3LjM2IChLSFRNTCwgbGlrZSBHZWNrbykgQ2hyb21lLzExNS4wLjAuMCBTYWZhcmkvNTM3LjM2IiwiYnJvd3Nlcl92ZXJzaW9uIjoiMTE1LjAuMC4wIiwib3NfdmVyc2lvbiI6IjEwIiwicmVmZXJyZXIiOiIiLCJyZWZlcnJpbmdfZG9tYWluIjoiIiwicmVmZXJyZXJfY3VycmVudCI6IiIsInJlZmVycmluZ19kb21haW5fY3VycmVudCI6IiIsInJlbGVhc2VfY2hhbm5lbCI6InN0YWJsZSIsImNsaWVudF9idWlsZF9udW1iZXIiOjIyNDI0NCwiY2xpZW50X2V2ZW50X3NvdXJjZSI6bnVsbH0=")
		}

		header.Add("x-captcha-key", config.CaptchaKey)
		header.Set("referer", fmt.Sprintf("https://discord.com/invite/%s", config.InviteCode))
	} else {
		pl = RegisterSimplePayload{
			Consent:                    true,
			Fingerprint:                c.xfingerprint,
			CaptchaKey:                 config.CaptchaKey,
			GlobalName:                 config.Username,
			UniqueUsernameRegistration: true,
		}

		header = c.GetHeader(&HeaderConfig{
			IsXtrack: false,
		})

		header.Del("x-context-properties")
		header.Del("authorization")
		header.Del("x-discord-timezone")
		header.Del("x-discord-locale")
		header.Del("x-debug-options")
		header.Del("x-track")
	}

	payload, err := json.Marshal(&pl)
	if err != nil {
		return nil, fmt.Errorf("error marshaling payload: %v", err.Error())
	}

	header.Add("x-fingerprint", c.xfingerprint)

	response, err := c.HttpClient.Do(cleanhttp.RequestOption{
		Method: "POST",
		Url:    "https://discord.com/api/v9/auth/register",
		Body:   bytes.NewReader(payload),
		Header: header,
	})

	fmt.Println(string(payload), header)

	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %v", err.Error())
	}

	defer response.Body.Close()

	resp, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var data RegisterResponse
	if err := json.Unmarshal([]byte(resp), &data); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %v", err.Error())
	}

	if data.Token != "" {
		c.Config.Token = data.Token
	}

	return &data, nil
}

func (c *Client) SetAvatar(config *AvatarConfig) error {
	var pfp string
	var err error

	if config.IsFromBase64 {
		pfp = config.Base64String
	} else {
		pfp, err = cleanhttp.ImageToBase64(config.FilePath)
		if err != nil {
			return err
		}
	}

	payload := fmt.Sprintf(`{"avatar": "%s"}`, pfp)

	header := c.GetHeader(&HeaderConfig{})
	header.Set("referer", "https://discord.com/channels/@me")

	response, err := c.HttpClient.Do(cleanhttp.RequestOption{
		Method: "PATCH",
		Url:    "https://discord.com/api/v9/users/@me",
		Body:   bytes.NewReader([]byte(payload)),
		Header: header,
	})

	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to set pfp, status code: %d", response.StatusCode)
	}

	return nil
}

func (c *Client) SetBirth(config *EditBirthConfig) error {
	payload, err := json.Marshal(&EditBirthPayload{
		DateOfBirth: config.Date,
	})
	if err != nil {
		return fmt.Errorf("error marshaling payload: %v", err.Error())
	}

	header := c.GetHeader(&HeaderConfig{})
	header.Set("referer", "https://discord.com/channels/@me")

	response, err := c.HttpClient.Do(cleanhttp.RequestOption{
		Method: "PATCH",
		Url:    "https://discord.com/api/v9/users/@me",
		Body:   bytes.NewReader(payload),
		Header: header,
	})
	if err != nil {
		return fmt.Errorf("error making HTTP request: %v", err.Error())
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to add birth-date, status code: %d", response.StatusCode)
	}

	return nil
}

func (c *Client) SetProfil(config *EditProfilConfig) error {
	payload, err := json.Marshal(&EditProfilPayload{
		Bio:         config.Bio,
		AccentColor: config.AccentColor,
	})
	if err != nil {
		return fmt.Errorf("error marshaling payload: %v", err.Error())
	}

	header := c.GetHeader(&HeaderConfig{})
	header.Set("referer", "https://discord.com/channels/@me")

	response, err := c.HttpClient.Do(cleanhttp.RequestOption{
		Method: "PATCH",
		Url:    "https://discord.com/api/v9/users/%40me/profile",
		Body:   bytes.NewReader(payload),
		Header: header,
	})
	if err != nil {
		return fmt.Errorf("error making HTTP request: %v", err.Error())
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to edit profile, status code: %d", response.StatusCode)
	}

	return nil
}

func (c *Client) SendMessage(config *SendMessageConfig) (any, error) {
	payload, err := json.Marshal(&MessagePayload{
		Content: config.Content,
		Tts:     config.Tts,
	})
	if err != nil {
		return nil, fmt.Errorf("error marshaling payload: %v", err.Error())
	}

	response, err := c.HttpClient.Do(cleanhttp.RequestOption{
		Method: "POST",
		Url:    fmt.Sprintf("https://discord.com/api/v9/channels/%s/messages", config.ChannelID),
		Body:   bytes.NewReader(payload),
		Header: c.GetHeader(&HeaderConfig{}),
	})
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %v", err.Error())
	}

	defer response.Body.Close()

	resp, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var data JoinServerResponse // wtf ?
	if err := json.Unmarshal([]byte(resp), &data); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %v", err.Error())
	}

	return &data, nil
}

func (c *Client) IsLocked() (bool, error) {
	response, err := c.HttpClient.Do(cleanhttp.RequestOption{
		Method: "GET",
		Url:    "https://discord.com/api/v9/users/@me/affinities/users",
		Body:   nil,
		Header: c.GetHeader(&HeaderConfig{}),
	})
	if err != nil {
		return true, fmt.Errorf("error making HTTP request: %v", err.Error())
	}

	defer response.Body.Close()

	resp, err := io.ReadAll(response.Body)
	if err != nil {
		return false, err
	}

	if strings.Contains("You need to verify your account in order to perform this action.", string(resp)) {
		return true, nil
	}

	return false, nil
}

func (c *Client) SendFriend(config *FriendConfig) (bool, *CaptchaResponse, error) {
	if c.WsProperties.D.SessionID == "" {
		return false, nil, fmt.Errorf("please connect to the websocket first")
	}

	payload, err := json.Marshal(&FriendScience{
		Token: c.WsProperties.D.AnalyticsToken,
		Events: []ScEvent{
			{
				Type: "app_ui_viewed",
				Properties: ScProperties{
					ClientTrackTimestamp:        time.Now().UnixNano() / int64(time.Millisecond),
					ClientHeartbeatSessionID:    c.WsProperties.D.SessionID,
					TotalCompressedByteSize:     10196255,
					TotalUncompressedByteSize:   45546783,
					TotalTransferByteSize:       0,
					JSCompressedByteSize:        6291974,
					JSUncompressedByteSize:      28751914,
					JSTransferByteSize:          0,
					CSSCompressedByteSize:       513801,
					CSSUncompressedByteSize:     2806741,
					CSSTransferByteSize:         0,
					LoadID:                      uuid.NewString(),
					ScreenName:                  "friends",
					DurationMSSinceAppOpened:    1073,
					ClientPerformanceMemory:     0,
					AccessibilityFeatures:       524544,
					RenderedLocale:              "fr",
					AccessibilitySupportEnabled: false,
					ClientUUID:                  c.WsProperties.D.AuthSessionIDHash,
					ClientSendTimestamp:         time.Now().UnixNano() / int64(time.Millisecond),
				},
			},
			{
				Type: "ready_payload_received",
				Properties: ScProperties{
					ClientTrackTimestamp:           time.Now().UnixNano() / int64(time.Millisecond),
					ClientHeartbeatSessionID:       c.WsProperties.D.SessionID,
					CompressedByteSize:             13135,
					UncompressedByteSize:           46105,
					CompressionAlgorithm:           "zlib-stream",
					PackingAlgorithm:               "json",
					UnpackDurationMS:               0,
					IdentifyTotalServerDurationMS:  201,
					IdentifyAPIDurationMS:          100,
					IdentifyGuildsDurationMS:       1,
					NumGuilds:                      1,
					NumGuildChannels:               5,
					NumGuildCategoryChannels:       2,
					PresencesSize:                  2,
					UsersSize:                      223,
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
					AccessibilitySupportEnabled:    false,
					ClientUUID:                     c.WsProperties.D.AuthSessionIDHash,
					ClientSendTimestamp:            time.Now().UnixNano() / int64(time.Millisecond),
				},
			},
			{
				Type: "friends_list_viewed",
				Properties: ScProperties{
					ClientTrackTimestamp:        time.Now().UnixNano() / int64(time.Millisecond),
					ClientHeartbeatSessionID:    c.WsProperties.D.SessionID,
					TabOpened:                   "ADD_FRIEND",
					ClientPerformanceMemory:     0,
					AccessibilityFeatures:       524544,
					RenderedLocale:              "fr",
					AccessibilitySupportEnabled: false,
					ClientUUID:                  c.WsProperties.D.AuthSessionIDHash,
					ClientSendTimestamp:         time.Now().UnixNano() / int64(time.Millisecond),
				},
			},
			{
				Type: "app_notice_viewed",
				Properties: ScProperties{
					ClientTrackTimestamp:        time.Now().UnixNano() / int64(time.Millisecond),
					ClientHeartbeatSessionID:    c.WsProperties.D.SessionID,
					TabOpened:                   "UNCLAIMED_ACCOUNT",
					ClientPerformanceMemory:     0,
					AccessibilityFeatures:       524544,
					RenderedLocale:              "fr",
					AccessibilitySupportEnabled: false,
					ClientUUID:                  c.WsProperties.D.AuthSessionIDHash,
					ClientSendTimestamp:         time.Now().UnixNano() / int64(time.Millisecond),
				},
			},
		},
	})
	if err != nil {
		return false, nil, fmt.Errorf("error marshaling payload: %v", err.Error())
	}

	scresponse, err := c.HttpClient.Do(cleanhttp.RequestOption{
		Method: "POST",
		Url:    "https://discord.com/api/v9/science",
		Body:   bytes.NewReader(payload),
		Header: c.GetHeader(&HeaderConfig{}),
	})
	if err != nil {
		return false, nil, fmt.Errorf("error making HTTP request: %v", err.Error())
	}

	defer scresponse.Body.Close()

	payload, err = json.Marshal(&FriendPayload{
		Username: config.Username,
	})
	if err != nil {
		return false, nil, fmt.Errorf("error marshaling payload: %v", err.Error())
	}

	header := c.GetHeader(&HeaderConfig{
		IsAddFriend: true,
	})

	if config.Captcha != "" && config.RqToken != "" {
		header.Add("x-captcha-key", config.Captcha)
		header.Add("x-captcha-rqtoken", config.RqToken)
	}

	response, err := c.HttpClient.Do(cleanhttp.RequestOption{
		Method: "POST",
		Url:    "https://discord.com/api/v9/users/@me/relationships",
		Body:   bytes.NewReader(payload),
		Header: header,
	})
	if err != nil {
		return false, nil, fmt.Errorf("error making HTTP request: %v", err.Error())
	}

	defer response.Body.Close()

	switch response.StatusCode {
	case 204:
		return true, nil, nil
	case 400:
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return false, nil, err
		}

		var c CaptchaResponse
		if err := json.Unmarshal(body, &c); err != nil {
			return false, nil, err
		}

		return false, &c, fmt.Errorf("captcha spawn")
	case 403:
		return false, nil, fmt.Errorf("user not found")
	default:
		return false, nil, fmt.Errorf("unknown status: %d", response.StatusCode)
	}
}

func (c *Client) SendCaptchaEvent(SiteKey string) error {
	if c.WsProperties.D.SessionID == "" {
		return fmt.Errorf("please connect to the websocket first")
	}

	payload, err := json.Marshal(&FriendScience{
		Token: c.WsProperties.D.AnalyticsToken,
		Events: []ScEvent{
			{
				Type: "captcha_event",
				Properties: ScProperties{
					AccessibilityFeatures:       256,
					AccessibilitySupportEnabled: true,
					CaptchaEventName:            "verify",
					CaptchaFlowKey:              "",
					CaptchaService:              "hcaptcha",
					ClientPerformanceMemory:     0,
					ClientSendTimestamp:         time.Now().UnixNano() / int64(time.Millisecond),
					ClientTrackTimestamp:        time.Now().UnixNano() / int64(time.Millisecond),
					ClientUUID:                  c.WsProperties.D.AuthSessionIDHash,
					RenderedLocale:              c.HttpClient.Config.BrowserFp.Navigator.Language,
					SiteKey:                     SiteKey,
				},
			},
		},
	})
	if err != nil {
		return fmt.Errorf("error marshaling payload: %v", err.Error())
	}

	scresponse, err := c.HttpClient.Do(cleanhttp.RequestOption{
		Method: "POST",
		Url:    "https://discord.com/api/v9/science",
		Body:   bytes.NewReader(payload),
		Header: c.GetHeader(&HeaderConfig{}),
	})
	if err != nil {
		return fmt.Errorf("error making HTTP request: %v", err.Error())
	}

	defer scresponse.Body.Close()

	return nil
}
