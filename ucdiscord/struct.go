package discord

import (
	"github.com/Implex-ltd/cleanhttp/cleanhttp"
)

type ClientConfig struct {
	Token               string
	GetCookies          bool
	GetCloudflareCookes bool
	BuildNumber         int
	Client              *cleanhttp.CleanHttp
}

type JoinConfig struct {
	InviteCode string
	GuildID    string
	ChannelID  string
}

type JoinPayload struct {
	SessionID string `json:"session_id"`
}

type RegisterConfig struct {
	Username         string
	InviteCode       string
	CaptchaKey       string
	CustomProperties string
}

type RegisterSimplePayload struct {
	Consent                    bool   `json:"consent"`
	Fingerprint                string `json:"fingerprint"`
	CaptchaKey                 string `json:"captcha_key"`
	GlobalName                 string `json:"global_name"`
	UniqueUsernameRegistration bool   `json:"unique_username_registration"`
}

type RegisterInvitePayload struct {
	Fingerprint                string `json:"fingerprint"`
	GlobalName                 string `json:"global_name"`
	Invite                     string `json:"invite"`
	Consent                    bool   `json:"consent"`
	GiftCodeSkuID              any    `json:"gift_code_sku_id"`
	UniqueUsernameRegistration bool   `json:"unique_username_registration"`
}

type RegisterResponse struct {
	// no captcha key provided
	CaptchaKey     []string `json:"captcha_key"`
	CaptchaSitekey string   `json:"captcha_sitekey"`
	CaptchaService string   `json:"captcha_service"`

	// success
	Token string `json:"token"`
}

type AvatarConfig struct {
	FilePath     string
	Base64String string
	IsFromBase64 bool
}

type EditBirthConfig struct {
	Date string
}

type EditBirthPayload struct {
	DateOfBirth string `json:"date_of_birth"`
}

type EditProfilConfig struct {
	Bio         string
	AccentColor int
}

type EditProfilPayload struct {
	Bio         string `json:"bio"`
	AccentColor int    `json:"accent_color"`
}

type Client struct {
	Config       *ClientConfig
	HttpClient   *cleanhttp.CleanHttp
	WsProperties WsLoginResponse
	SessionID    string
	BuildNumber  int
	UaInfo       *cleanhttp.UserAgentInfo
	xfingerprint string
}

type XProperties struct {
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

type ContextProperties struct {
	Location            string `json:"location"`
	LocationGuildID     string `json:"location_guild_id"`
	LocationChannelID   string `json:"location_channel_id"`
	LocationChannelType int    `json:"location_channel_type"`
}

type HeaderConfig struct {
	Join        *JoinConfig
	IsAddFriend bool
	IsXtrack    bool
	IsSuper     bool
}

type FingerprintResponse struct {
	Fingerprint string  `json:"fingerprint"`
	Assignments [][]any `json:"assignments"`
}

type JoinServerResponse struct {
	// in case of error
	Message        string   `json:"message"`
	CaptchaKey     []string `json:"captcha_key"`
	CaptchaSitekey string   `json:"captcha_sitekey"`
	CaptchaService string   `json:"captcha_service"`
	CaptchaRqdata  string   `json:"captcha_rqdata"`
	CaptchaRqtoken string   `json:"captcha_rqtoken"`

	Code      int `json:"code"`
	Type      int `json:"type"`
	ExpiresAt any `json:"expires_at"`
	Guild     struct {
		ID                       string   `json:"id"`
		Name                     string   `json:"name"`
		Splash                   string   `json:"splash"`
		Banner                   string   `json:"banner"`
		Description              any      `json:"description"`
		Icon                     string   `json:"icon"`
		Features                 []string `json:"features"`
		VerificationLevel        int      `json:"verification_level"`
		VanityURLCode            string   `json:"vanity_url_code"`
		PremiumSubscriptionCount int      `json:"premium_subscription_count"`
		Nsfw                     bool     `json:"nsfw"`
		NsfwLevel                int      `json:"nsfw_level"`
		WelcomeScreen            struct {
			Description     string `json:"description"`
			WelcomeChannels []struct {
				ChannelID   string `json:"channel_id"`
				Description string `json:"description"`
				EmojiID     string `json:"emoji_id"`
				EmojiName   string `json:"emoji_name"`
			} `json:"welcome_channels"`
		} `json:"welcome_screen"`
	} `json:"guild"`
	Channel struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Type int    `json:"type"`
	} `json:"channel"`
	NewMember bool `json:"new_member"`
}

type SendMessageConfig struct {
	Content   string
	Nonce     string
	Tts       bool
	Flags     int
	ChannelID string
}

type MessagePayload struct {
	Content string `json:"content"`
	Nonce   string `json:"nonce"`
	Tts     bool   `json:"tts"`
	Flags   int    `json:"flags"`
}

type ScProperties struct {
	CaptchaEventName               string `json:"captcha_event_name"`
	CaptchaFlowKey                 string `json:"captcha_flow_key"`
	CaptchaService                 string `json:"captcha_service"`
	SiteKey                        string `json:"sitekey"`
	ClientTrackTimestamp           int64  `json:"client_track_timestamp"`
	ClientHeartbeatSessionID       string `json:"client_heartbeat_session_id"`
	TotalCompressedByteSize        int64  `json:"total_compressed_byte_size,omitempty"`
	TotalUncompressedByteSize      int64  `json:"total_uncompressed_byte_size,omitempty"`
	TotalTransferByteSize          int64  `json:"total_transfer_byte_size,omitempty"`
	JSCompressedByteSize           int64  `json:"js_compressed_byte_size,omitempty"`
	JSUncompressedByteSize         int64  `json:"js_uncompressed_byte_size,omitempty"`
	JSTransferByteSize             int64  `json:"js_transfer_byte_size,omitempty"`
	CSSCompressedByteSize          int64  `json:"css_compressed_byte_size,omitempty"`
	CSSUncompressedByteSize        int64  `json:"css_uncompressed_byte_size,omitempty"`
	CSSTransferByteSize            int64  `json:"css_transfer_byte_size,omitempty"`
	LoadID                         string `json:"load_id,omitempty"`
	ScreenName                     string `json:"screen_name,omitempty"`
	DurationMSSinceAppOpened       int64  `json:"duration_ms_since_app_opened,omitempty"`
	ClientPerformanceMemory        int64  `json:"client_performance_memory"`
	AccessibilityFeatures          int64  `json:"accessibility_features"`
	RenderedLocale                 string `json:"rendered_locale"`
	AccessibilitySupportEnabled    bool   `json:"accessibility_support_enabled"`
	ClientUUID                     string `json:"client_uuid"`
	ClientSendTimestamp            int64  `json:"client_send_timestamp"`
	CompressedByteSize             int64  `json:"compressed_byte_size,omitempty"`
	UncompressedByteSize           int64  `json:"uncompressed_byte_size,omitempty"`
	CompressionAlgorithm           string `json:"compression_algorithm,omitempty"`
	PackingAlgorithm               string `json:"packing_algorithm,omitempty"`
	UnpackDurationMS               int64  `json:"unpack_duration_ms,omitempty"`
	IdentifyTotalServerDurationMS  int64  `json:"identify_total_server_duration_ms,omitempty"`
	IdentifyAPIDurationMS          int64  `json:"identify_api_duration_ms,omitempty"`
	IdentifyGuildsDurationMS       int64  `json:"identify_guilds_duration_ms,omitempty"`
	NumGuilds                      int64  `json:"num_guilds,omitempty"`
	NumGuildChannels               int64  `json:"num_guild_channels,omitempty"`
	NumGuildCategoryChannels       int64  `json:"num_guild_category_channels,omitempty"`
	PresencesSize                  int64  `json:"presences_size,omitempty"`
	UsersSize                      int64  `json:"users_size,omitempty"`
	ReadStatesSize                 int64  `json:"read_states_size,omitempty"`
	PrivateChannelsSize            int64  `json:"private_channels_size,omitempty"`
	UserSettingsSize               int64  `json:"user_settings_size,omitempty"`
	ExperimentsSize                int64  `json:"experiments_size,omitempty"`
	UserGuildSettingsSize          int64  `json:"user_guild_settings_size,omitempty"`
	RelationshipsSize              int64  `json:"relationships_size,omitempty"`
	RemainingDataSize              int64  `json:"remaining_data_size,omitempty"`
	GuildChannelsSize              int64  `json:"guild_channels_size,omitempty"`
	GuildMembersSize               int64  `json:"guild_members_size,omitempty"`
	GuildPresencesSize             int64  `json:"guild_presences_size,omitempty"`
	GuildRolesSize                 int64  `json:"guild_roles_size,omitempty"`
	GuildEmojisSize                int64  `json:"guild_emojis_size,omitempty"`
	GuildThreadsSize               int64  `json:"guild_threads_size,omitempty"`
	GuildStickersSize              int64  `json:"guild_stickers_size,omitempty"`
	GuildEventsSize                int64  `json:"guild_events_size,omitempty"`
	GuildFeaturesSize              int64  `json:"guild_features_size,omitempty"`
	GuildRemainingDataSize         int64  `json:"guild_remaining_data_size,omitempty"`
	SizeMetricsDurationMS          int64  `json:"size_metrics_duration_ms,omitempty"`
	DurationMSSinceIdentifyStart   int64  `json:"duration_ms_since_identify_start,omitempty"`
	DurationMSSinceConnectionStart int64  `json:"duration_ms_since_connection_start,omitempty"`
	DurationMSSinceEmitStart       int64  `json:"duration_ms_since_emit_start,omitempty"`
	IsReconnect                    bool   `json:"is_reconnect,omitempty"`
	IsFastConnect                  bool   `json:"is_fast_connect,omitempty"`
	DidForceClearGuildHashes       bool   `json:"did_force_clear_guild_hashes,omitempty"`
	IdentifyUncompressedByteSize   int64  `json:"identify_uncompressed_byte_size,omitempty"`
	IdentifyCompressedByteSize     int64  `json:"identify_compressed_byte_size,omitempty"`
	HadCacheAtStartup              bool   `json:"had_cache_at_startup,omitempty"`
	UsedCacheAtStartup             bool   `json:"used_cache_at_startup,omitempty"`
	TabOpened                      string `json:"tab_opened,omitempty"`
	NoticeType                     string `json:"notice_type,omitempty"`
}

type ScEvent struct {
	Type       string       `json:"type"`
	Properties ScProperties `json:"properties"`
}

type FriendScience struct {
	Token  string    `json:"token"`
	Events []ScEvent `json:"events"`
}

type FriendConfig struct {
	Username string
	Captcha  string
	RqToken  string
}

type FriendPayload struct {
	Username      string `json:"username"`
	Discriminator any    `json:"discriminator"`
}

type CaptchaResponse struct {
	CaptchaKey     []string `json:"captcha_key"`
	CaptchaSitekey string   `json:"captcha_sitekey"`
	CaptchaService string   `json:"captcha_service"`
	CaptchaRqdata  string   `json:"captcha_rqdata"`
	CaptchaRqtoken string   `json:"captcha_rqtoken"`
}
