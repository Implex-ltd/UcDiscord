package ucdiscord

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	http "github.com/bogdanfinn/fhttp"

)

func (C *Client) GetProperties(config *PropInfo) (headerName string, headerValue string, err error) {
	bn := C.Config.Build

	switch config.Type {
	case PROP_CONTEXT:
		headerName = "x-context-properties"
	case PROP_SUPER:
		headerName = "x-super-properties"
	case PROP_TRACK:
		headerName = "x-track"
		bn = 9999
	default:
		return "", "", fmt.Errorf("invalid type")
	}

	payload, err := json.Marshal(&XProp{
		Os:                     C.UserAgent.OSName,
		Browser:                C.UserAgent.BrowserName,
		Device:                 "",
		SystemLocale:           C.Config.Http.Config.BrowserFp.Navigator.Language,
		BrowserUserAgent:       C.Config.Http.Config.BrowserFp.Navigator.UserAgent,
		BrowserVersion:         C.UserAgent.BrowserVersion,
		OsVersion:              C.UserAgent.OSVersion,
		Referrer:               "",
		ReferringDomain:        "",
		ReferrerCurrent:        "",
		ReferringDomainCurrent: "",
		ReleaseChannel:         "stable",
		ClientBuildNumber:      bn,
		ClientEventSource:      nil,
	})

	if err != nil {
		return "", "", err
	}

	return headerName, AddBase64Padding(base64.RawStdEncoding.EncodeToString(payload)), nil
}

func (C *Client) GetCtxProperties(GuildID, ChannelID, Location string) (string, error) {
	props := CtxProp{
		Location: Location,
	}

	if GuildID != "" && ChannelID != "" && Location == LOCATION_JOIN_GUILD {
		props.LocationGuildID = GuildID
		props.LocationChannelID = ChannelID
		props.LocationChannelType = 0
	}

	payload, err := CustomMarshalProp(props)
	if err != nil {
		return "", err
	}

	return AddBase64Padding(base64.RawStdEncoding.EncodeToString(payload)), nil
}

func (C *Client) GetHeader(config *HeaderConfig) http.Header {
	header := http.Header{
		`authority`:          {`discord.com`},
		`accept`:             {`*/*`},
		`accept-language`:    {C.Config.Http.BaseHeader.AcceptLanguage},
		`accept-encoding`:    {`gzip, deflate, br`},
		`content-type`:       {`application/json`},
		`origin`:             {`https://discord.com`},
		`referer`:            {fmt.Sprintf("https://discord.com%s", config.Referer)},
		`sec-ch-ua`:          {C.Config.Http.BaseHeader.SecChUa},
		`sec-ch-ua-mobile`:   {C.Config.Http.BaseHeader.SecChUaMobile},
		`sec-ch-ua-platform`: {C.Config.Http.BaseHeader.SecChUaPlatform},
		`sec-fetch-dest`:     {`empty`},
		`sec-fetch-mode`:     {`cors`},
		`sec-fetch-site`:     {`same-origin`},
		`user-agent`:         {C.Config.Http.Config.BrowserFp.Navigator.UserAgent},

		http.HeaderOrderKey: {
			`authority`,
			`accept`,
			`accept-language`,
			`accept-encoding`,
			`authorization`,
			`content-type`,
			`cookie`,
			`origin`,
			`referer`,
			`sec-ch-ua`,
			`sec-ch-ua-mobile`,
			`sec-ch-ua-platform`,
			`sec-fetch-dest`,
			`sec-fetch-mode`,
			`sec-fetch-site`,
			`user-agent`,
			`x-captcha-key`,
			`x-captcha-rqtoken`,
			`x-context-properties`,
			`x-debug-options`,
			`x-discord-locale`,
			`x-discord-timezone`,
			`x-fingerprint`,
			`x-super-properties`,
			`x-track`,
		},
	}

	name, value, _ := C.GetProperties(config.Info)
	header.Set(name, value)

	if config.Info.Type != PROP_TRACK {
		for k, v := range map[string]string{
			`x-debug-options`:    `bugReporterEnabled`,
			`x-discord-locale`:   `fr`,
			`x-discord-timezone`: `Europe/Paris`,
		} {
			header.Set(k, v)
		}
	}

	if C.Config.Token != "" {
		header.Set("authorization", C.Config.Token)
	}

	if config.Join {
		value, _ := C.GetCtxProperties(config.GuildID, config.ChannelID, LOCATION_JOIN_GUILD)
		header.Set(`x-context-properties`, value)
	}

	if config.Friend {
		value, _ := C.GetCtxProperties(config.GuildID, config.ChannelID, LOCATION_ADD_FRIEND)
		header.Set(`x-context-properties`, value)
	}

	if config.ContextProperties != "" {
		header.Set(`x-context-properties`, config.ContextProperties)
	}

	if config.CaptchaKey != "" {
		header.Set(`x-captcha-key`, config.CaptchaKey)
	}

	if config.CaptchaRqtoken != "" {
		header.Set(`x-captcha-rqtoken`, config.CaptchaRqtoken)
	}

	if config.Fingerprint != "" {
		header.Set(`x-fingerprint`, config.Fingerprint)
	}

	return header
}
