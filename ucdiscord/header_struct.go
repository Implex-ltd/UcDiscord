package ucdiscord

var (
	PROP_TRACK   = 0
	PROP_SUPER   = 1
	PROP_CONTEXT = 2

	LOCATION_JOIN_GUILD = "Join Guild"
	LOCATION_ADD_FRIEND = "Add Friend"
)

type HeaderConfig struct {
	Info                           *PropInfo
	Join, Friend                   bool
	GuildID, ChannelID             string
	CaptchaKey, CaptchaRqtoken     string
	Referer                        string
	Fingerprint, ContextProperties string
}

type PropInfo struct {
	Type int
}

type XProp struct {
	Os                     string `json:"os"`
	Browser                string `json:"browser"`
	Device                 string `json:"device"`
	SystemLocale           string `json:"system_locale"`
	BrowserUserAgent       string `json:"browser_user_agent"`
	BrowserVersion         string `json:"browser_version"`
	OsVersion              string `json:"os_version"`
	Referrer               string `json:"referrer"`
	ReferringDomain        string `json:"referring_domain"`
	ReferrerCurrent        string `json:"referrer_current"`
	ReferringDomainCurrent string `json:"referring_domain_current"`
	ReleaseChannel         string `json:"release_channel"`
	ClientBuildNumber      int    `json:"client_build_number"`
	ClientEventSource      any    `json:"client_event_source"`
}

type CtxProp struct {
	Location            string `json:"location"`
	LocationGuildID     string `json:"location_guild_id"`
	LocationChannelID   string `json:"location_channel_id"`
	LocationChannelType int    `json:"location_channel_type"`
}
