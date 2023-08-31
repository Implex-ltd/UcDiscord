package discord

import (
	"encoding/base64"
	"encoding/json"
	"strings"
	_ "strings"

	http "github.com/bogdanfinn/fhttp"
)

const (
	PROPERTYPE_XTRACK  = 0
	PROPERTYPE_SUPER   = 1
	PROPERTYPE_CONTEXT = 2
)

func (c *Client) getProperties(ProperType int) (string, string) {
	build := 0
	headerName := ""

	switch ProperType {
	case PROPERTYPE_CONTEXT:
		build = c.BuildNumber
		headerName = "x-context-properties"
		break
	case PROPERTYPE_SUPER:
		build = c.BuildNumber
		headerName = "x-super-properties"
		break
	case PROPERTYPE_XTRACK:
		build = 9999
		headerName = "x-track"
		break
	}

	payload, _ := json.Marshal(&XProperties{
		Os:                     c.UaInfo.OSName,
		Browser:                c.UaInfo.BrowserName,
		Device:                 "",
		SystemLocale:           c.HttpClient.Config.BrowserFp.Navigator.Language,
		BrowserUserAgent:       c.HttpClient.Config.BrowserFp.Navigator.UserAgent,
		BrowserVersion:         c.UaInfo.BrowserVersion,
		OsVersion:              c.UaInfo.OSVersion,
		Referrer:               "",
		ReferringDomain:        "",
		ReferrerCurrent:        "",
		ReferringDomainCurrent: "",
		ReleaseChannel:         "stable",
		ClientBuildNumber:      build,
		ClientEventSource:      nil,
	})

	return headerName, strings.ReplaceAll(addBase64Padding(base64.RawStdEncoding.EncodeToString(payload)), "=", "")
}

func (c *Client) getContextProperties(config *JoinConfig) string {
	payload, _ := json.Marshal(&ContextProperties{
		LocationGuildID:     config.GuildID,
		LocationChannelID:   config.ChannelID,
		LocationChannelType: 0,
		Location:            "Join Guild",
	})

	return addBase64Padding(base64.RawStdEncoding.EncodeToString(payload))
}

func (c *Client) getHeader(config *HeaderConfig) http.Header {
	ctx := ""

	if config.Join != nil {
		ctx = c.getContextProperties(config.Join)
	}

	if config.IsAddFriend {
		ctx = "eyJsb2NhdGlvbiI6IkFkZCBGcmllbmQifQ=="

	}

	headerName, properties := c.getProperties(config.ProperType)

	return http.Header{
		`accept`:               {`*/*`},
		`accept-encoding`:      {`gzip, deflate, br`},
		`accept-language`:      {c.HttpClient.BaseHeader.AcceptLanguage},
		`authorization`:        {c.Config.Token},
		`content-type`:         {`application/json`},
		`cookie`:               {c.HttpClient.BaseHeader.Cookies},
		`origin`:               {"https://discord.com"},
		`referer`:              {`https://discord.com`},
		`sec-ch-ua`:            {c.HttpClient.BaseHeader.SecChUa},
		`sec-ch-ua-mobile`:     {c.HttpClient.BaseHeader.SecChUaMobile},
		`sec-ch-ua-platform`:   {c.HttpClient.BaseHeader.SecChUaPlatform},
		`sec-fetch-dest`:       {`empty`},
		`sec-fetch-mode`:       {`cors`},
		`sec-fetch-site`:       {`same-origin`},
		`user-agent`:           {c.HttpClient.Config.BrowserFp.Navigator.UserAgent},
		`x-context-properties`: {ctx},
		`x-debug-options`:      {`bugReporterEnabled`},
		`x-discord-locale`:     {"fr"},           // {strings.Split(c.HttpClient.Config.BrowserFp.Navigator.Language, "-")[0]},
		`x-discord-timezone`:   {`Europe/Paris`}, // todo: add country by ip or header language
		headerName:             {properties},

		http.HeaderOrderKey: {
			`authority`,
			`accept`,
			`accept-encoding`,
			`accept-language`,
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
			`x-context-properties`,
			`x-track`,
			`x-debug-options`,
			`x-discord-locale`,
			`x-discord-timezone`,
			headerName,
		},
	}
}
