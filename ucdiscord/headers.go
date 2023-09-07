package discord

import (
	"encoding/base64"
	"encoding/json"
	"net/url"
	_ "strings"

	http "github.com/bogdanfinn/fhttp"
)

var (
	u, _ = url.Parse("https://discord.com")
)

func (c *Client) getProperties(config *HeaderConfig) (string, string) {
	var build int
	var headerName string

	if config.IsXtrack {
		build = 9999
		headerName = "x-track"
	} else if config.IsSuper {
		build = c.BuildNumber
		headerName = "x-super-properties"
	} else {
		build = c.BuildNumber
		headerName = "x-context-properties"
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

	return headerName, addBase64Padding(base64.RawStdEncoding.EncodeToString(payload))
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

func (c *Client) GetHeader(config *HeaderConfig) http.Header {
	headerName, properties := c.getProperties(config)

	h := http.Header{
		`accept`:             {`*/*`},
		`accept-encoding`:    {`gzip, deflate, br`},
		`accept-language`:    {c.HttpClient.BaseHeader.AcceptLanguage},
		`authorization`:      {c.Config.Token},
		`content-type`:       {`application/json`},
		`cookie`:             {c.HttpClient.FormatCookies(u)},
		`origin`:             {"https://discord.com"},
		`referer`:            {`https://discord.com`},
		`sec-ch-ua`:          {c.HttpClient.BaseHeader.SecChUa},
		`sec-ch-ua-mobile`:   {c.HttpClient.BaseHeader.SecChUaMobile},
		`sec-ch-ua-platform`: {c.HttpClient.BaseHeader.SecChUaPlatform},
		`sec-fetch-dest`:     {`empty`},
		`sec-fetch-mode`:     {`cors`},
		`sec-fetch-site`:     {`same-origin`},
		`user-agent`:         {c.HttpClient.Config.BrowserFp.Navigator.UserAgent},
		`x-debug-options`:    {`bugReporterEnabled`},
		`x-discord-locale`:   {"fr"},           // {strings.Split(c.HttpClient.Config.BrowserFp.Navigator.Language, "-")[0]},
		`x-discord-timezone`: {`Europe/Paris`}, // todo: add country by ip or header language
		headerName:           {properties},

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

	if config.Join != nil {
		h.Set("x-context-properties", c.getContextProperties(config.Join))
	}

	if config.IsAddFriend {
		h.Set("x-context-properties", "eyJsb2NhdGlvbiI6IkFkZCBGcmllbmQifQ==")
	}

	return h
}
