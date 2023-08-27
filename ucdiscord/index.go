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
	//"github.com/Implex-ltd/cloudflare-reverse/cloudflarereverse"
	http "github.com/bogdanfinn/fhttp"
)

var (
	dUrl, _ = url.Parse("https://discord.com")
)

// Create new discord client. Return *Client.
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

// Get cookies and x-fingerprint. this function is called by defaut if you set "GetCookies" params ClientConfig.
func (c *Client) GetCookies() error {
	c.HttpClient.Client.SetCookies(dUrl, []*http.Cookie{{
		Name:  "locale",
		Value: strings.Split(c.HttpClient.Config.BrowserFp.Navigator.Language, "-")[0],
	}})

	response, err := c.HttpClient.Do(cleanhttp.RequestOption{
		Method: "GET",
		Url:    "https://discord.com/api/v9/experiments",
		Header: c.getHeader(&HeaderConfig{}),
	})
	if err != nil {
		return fmt.Errorf("error making HTTP request: %v", err.Error())
	}

	defer response.Body.Close()

	resp, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	/*cookies := []*cyclepls.Cookie{}

	if c.Config.GetCloudflareCookes {
		cfbm, err := cloudflarereverse.GetCfbm(c.HttpClient.Config.BrowserFp, c.HttpClient.Config.Proxy) // make it proxyless because they are detecting proxies...
		if err != nil {
			return nil, fmt.Errorf("error getting Cloudflare cookies: %v", err.Error())
		}

		cfCookie := &http.Cookie{
			Name:  "cf_clearance",
			Value: cfbm,
		}
		cookies = append(cookies, cfCookie)
	}*/

	/*	c.HttpClient.Cookies = append(c.HttpClient.Cookies, &cyclepls.Cookie{
		Name:  "locale",
		Value: strings.Split(c.HttpClient.Config.BrowserFp.Navigator.Language, "-")[0],
	})*/

	//c.HttpClient.Cookies = append(c.HttpClient.Cookies, cookies...)

	var fp FingerprintResponse
	if err := json.Unmarshal([]byte(resp), &fp); err != nil {
		return fmt.Errorf("cant unmarshal fingerpint: %v", err.Error())
	}

	c.xfingerprint = fp.Fingerprint

	return nil
}

// Join server and return *JoinServerResponse, take *JoinConfig as params. WARN: need to connect trougth websocket first.
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

	header := c.getHeader(&HeaderConfig{
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

// Create discord accoutn and return *RegisterResponse, take *ResgisterConfig as param.
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

		header = c.getHeader(&HeaderConfig{
			IsXtrack: false,
		})

		header.Set("x-fingerprint", c.xfingerprint)
		header.Set("x-super-properties", "eyJvcyI6IldpbmRvd3MiLCJicm93c2VyIjoiQ2hyb21lIiwiZGV2aWNlIjoiIiwic3lzdGVtX2xvY2FsZSI6ImZyLUZSIiwiYnJvd3Nlcl91c2VyX2FnZW50IjoiTW96aWxsYS81LjAgKFdpbmRvd3MgTlQgMTAuMDsgV2luNjQ7IHg2NCkgQXBwbGVXZWJLaXQvNTM3LjM2IChLSFRNTCwgbGlrZSBHZWNrbykgQ2hyb21lLzExMy4wLjU2NzIuOTQgU2FmYXJpLzUzNy4zNiIsImJyb3dzZXJfdmVyc2lvbiI6IjExMy4wLjU2NzIuOTQiLCJvc192ZXJzaW9uIjoiMTAiLCJyZWZlcnJlciI6IiIsInJlZmVycmluZ19kb21haW4iOiIiLCJyZWZlcnJlcl9jdXJyZW50IjoiIiwicmVmZXJyaW5nX2RvbWFpbl9jdXJyZW50IjoiIiwicmVsZWFzZV9jaGFubmVsIjoic3RhYmxlIiwiY2xpZW50X2J1aWxkX251bWJlciI6MjIyMzUyLCJjbGllbnRfZXZlbnRfc291cmNlIjpudWxsfQ==")

		//header.Del("x-context-properties")
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

		header = c.getHeader(&HeaderConfig{
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

// Add avatar to the discord account, return error, take *AvatarConfig as param.
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

	header := c.getHeader(&HeaderConfig{})
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

// Edit user profil. take *EditProfilConfig as param, return error
func (c *Client) SetBirth(config *EditBirthConfig) error {
	payload, err := json.Marshal(&EditBirthPayload{
		DateOfBirth: config.Date,
	})
	if err != nil {
		return fmt.Errorf("error marshaling payload: %v", err.Error())
	}

	header := c.getHeader(&HeaderConfig{})
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

// Edit user profil. take *EditProfilConfig as param, return error
func (c *Client) SetProfil(config *EditProfilConfig) error {
	payload, err := json.Marshal(&EditProfilPayload{
		Bio:         config.Bio,
		AccentColor: config.AccentColor,
	})
	if err != nil {
		return fmt.Errorf("error marshaling payload: %v", err.Error())
	}

	header := c.getHeader(&HeaderConfig{})
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

// Send message to a server.
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
		Header: c.getHeader(&HeaderConfig{}),
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

// Check if token is locked
func (c *Client) IsLocked() (bool, error) {
	response, err := c.HttpClient.Do(cleanhttp.RequestOption{
		Method: "GET",
		Url:    "https://discord.com/api/v9/users/@me/affinities/users",
		Body:   nil,
		Header: c.getHeader(&HeaderConfig{}),
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

// Send friend request to a user.
func (c *Client) SendFriend(config *FriendConfig) (bool, *CaptchaResponse, error) {
	if c.WsProperties.D.SessionID == "" {
		return false, nil, fmt.Errorf("please connect to the websocket first")
	}

	payload, err := json.Marshal(&FriendScience{
		Token: c.WsProperties.D.AnalyticsToken,
		Events: []ScEvent{
			{
				Type: "friends_list_viewed",
				Properties: ScProperties{
					ClientTrackTimestamp:        time.Now().UnixNano(),
					ClientHeartbeatSessionID:    c.WsProperties.D.SessionID,
					TabOpened:                   "ADD_FRIEND",
					ClientPerformanceMemory:     0,
					AccessibilityFeatures:       524544,
					RenderedLocale:              "fr",
					AccessibilitySupportEnabled: false,
					ClientUUID:                  c.WsProperties.D.AuthSessionIDHash,
					ClientSendTimestamp:         time.Now().Unix(),
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
		Header: c.getHeader(&HeaderConfig{}),
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

	response, err := c.HttpClient.Do(cleanhttp.RequestOption{
		Method: "POST",
		Url:    "https://discord.com/api/v9/users/@me/relationships",
		Body:   bytes.NewReader(payload),
		Header: c.getHeader(&HeaderConfig{
			IsAddFriend: true,
		}),
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

		fmt.Println(string(body))

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
