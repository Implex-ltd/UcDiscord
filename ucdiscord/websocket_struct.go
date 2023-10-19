package ucdiscord

import (
	"github.com/dgrr/fastws"
)

type ClientWebsocket struct {
	Open, Ready, LogNonImplemented, Debug bool
	Token                                 string
	Conn                                  *fastws.Conn
	Prop                                  *XProp
	ReadyData                             *D
}

type WsData struct {
	Op int64 `json:"op"`
	D  any   `json:"d"`
}

type Game struct {
	Name string `json:"name"`
	Type int64  `json:"type"`
}

type Activity struct {
	Name       string     `json:"name"`
	Type       int64      `json:"type"`
	State      string     `json:"state"`
	Timestamps Timestamps `json:"timestamps"`
	Emoji      Emoji      `json:"emoji"`
}

type Timestamps struct {
	End int64 `json:"end"`
}

type ClientState struct {
	GuildVersions            GuildVersions `json:"guild_versions"`
	HighestLastMessageID     string        `json:"highest_last_message_id"`
	ReadStateVersion         int64         `json:"read_state_version"`
	UserGuildSettingsVersion int64         `json:"user_guild_settings_version"`
	UserSettingsVersion      int64         `json:"user_settings_version"`
	PrivateChannelsVersion   string        `json:"private_channels_version"`
	APICodeVersion           int64         `json:"api_code_version"`
}

type GuildVersions struct {
}

type Presence struct {
	Status     string     `json:"status"`
	Since      int64      `json:"since"`
	Game       Game       `json:"game"`
	Activities []Activity `json:"activities"`
	Afk        bool       `json:"afk"`
}

type OpLoginResponse struct {
	T  any `json:"t"`
	S  any `json:"s"`
	Op int `json:"op"`
	D  struct {
		HeartbeatInterval int      `json:"heartbeat_interval"`
		Trace             []string `json:"_trace"`
	} `json:"d"`
}

// reply op_ready

type TData struct {
	T  string `json:"t"`
	S  int64  `json:"s"`
	Op int64  `json:"op"`
}

type Reply struct {
	T  string `json:"t"`
	S  int64  `json:"s"`
	Op int64  `json:"op"`
	D  D      `json:"d"`
}

type D struct {
	// login reply
	V                 int64            `json:"v"`
	Users             []UserElement    `json:"users"`
	UserSettingsProto string           `json:"user_settings_proto"`
	UserGuildSettings ReadState        `json:"user_guild_settings"`
	User              PurpleUser       `json:"user"`
	Tutorial          Tutorial         `json:"tutorial"`
	Sessions          []Session        `json:"sessions"`
	SessionType       string           `json:"session_type"`
	SessionID         string           `json:"session_id"`
	ResumeGatewayURL  string           `json:"resume_gateway_url"`
	Relationships     []interface{}    `json:"relationships"`
	ReadState         ReadState        `json:"read_state"`
	PrivateChannels   []interface{}    `json:"private_channels"`
	MergedMembers     [][]MergedMember `json:"merged_members"`
	Guilds            []Guild          `json:"guilds"`
	GuildJoinRequests []interface{}    `json:"guild_join_requests"`
	//GuildExperiments      [][]DGuildExperiment `json:"guild_experiments"`
	GeoOrderedRTCRegions  []string      `json:"geo_ordered_rtc_regions"`
	FriendSuggestionCount int64         `json:"friend_suggestion_count"`
	Experiments           [][]int64     `json:"experiments"`
	CountryCode           string        `json:"country_code"`
	Consents              Consents      `json:"consents"`
	ConnectedAccounts     []interface{} `json:"connected_accounts"`
	AuthSessionIDHash     string        `json:"auth_session_id_hash"`
	Auth                  Auth          `json:"auth"`
	APICodeVersion        int64         `json:"api_code_version"`
	AnalyticsToken        string        `json:"analytics_token"`
	Trace                 []string      `json:"_trace"`

	// login
	Token        *string      `json:"token,omitempty"`
	Capabilities *int         `json:"capabilities,omitempty"`
	Properties   *XProp       `json:"properties,omitempty"`
	Presence     *Presence    `json:"presence,omitempty"`
	Compress     *bool        `json:"compress,omitempty"`
	ClientState  *ClientState `json:"client_state,omitempty"`
	Status       *string      `json:"status,omitempty"`
	Since        *int64       `json:"since,omitempty"`
	Activities   []Activity   `json:"activities,omitempty"`
	Afk          *bool        `json:"afk,omitempty"`
}

type Auth struct {
	AuthenticatorTypes []interface{} `json:"authenticator_types"`
}

type ClientInfo struct {
	Version int64  `json:"version"`
	OS      string `json:"os"`
	Client  string `json:"client"`
}

type Consents struct {
	Personalization Personalization `json:"personalization"`
}

type Personalization struct {
	Consented bool `json:"consented"`
}

type PurpleGuildExperiment struct {
	S int64 `json:"s"`
	E int64 `json:"e"`
}

type FluffyGuildExperiment struct {
	K []string `json:"k"`
	B int64    `json:"b"`
}

type Guild struct {
	Version                  int64            `json:"version"`
	Threads                  []interface{}    `json:"threads"`
	Stickers                 []interface{}    `json:"stickers"`
	StageInstances           []interface{}    `json:"stage_instances"`
	Roles                    []Role           `json:"roles"`
	Properties               Properties       `json:"properties"`
	PremiumSubscriptionCount int64            `json:"premium_subscription_count"`
	MemberCount              int64            `json:"member_count"`
	Lazy                     bool             `json:"lazy"`
	Large                    bool             `json:"large"`
	JoinedAt                 string           `json:"joined_at"`
	ID                       string           `json:"id"`
	GuildScheduledEvents     []interface{}    `json:"guild_scheduled_events"`
	Emojis                   []Emoji          `json:"emojis"`
	DataMode                 string           `json:"data_mode"`
	Channels                 []Channel        `json:"channels"`
	ApplicationCommandCounts map[string]int64 `json:"application_command_counts"`
}

type Channel struct {
	Type                 int64                 `json:"type"`
	Topic                *string               `json:"topic"`
	ThemeColor           interface{}           `json:"theme_color"`
	Status               interface{}           `json:"status"`
	RateLimitPerUser     *int64                `json:"rate_limit_per_user,omitempty"`
	Position             int64                 `json:"position"`
	PermissionOverwrites []PermissionOverwrite `json:"permission_overwrites"`
	ParentID             *string               `json:"parent_id"`
	Nsfw                 *bool                 `json:"nsfw,omitempty"`
	Name                 string                `json:"name"`
	LastMessageID        *string               `json:"last_message_id"`
	ID                   string                `json:"id"`
	IconEmoji            *IconEmoji            `json:"icon_emoji,omitempty"`
	Flags                int64                 `json:"flags"`
	UserLimit            *int64                `json:"user_limit,omitempty"`
	RTCRegion            interface{}           `json:"rtc_region"`
	Bitrate              *int64                `json:"bitrate,omitempty"`
}

type IconEmoji struct {
	Name string      `json:"name"`
	ID   interface{} `json:"id"`
}

type PermissionOverwrite struct {
	Type  int64  `json:"type"`
	ID    string `json:"id"`
	Deny  string `json:"deny"`
	Allow string `json:"allow"`
}

type Emoji struct {
	Roles         []interface{} `json:"roles"`
	RequireColons bool          `json:"require_colons"`
	Name          string        `json:"name"`
	Managed       bool          `json:"managed"`
	ID            string        `json:"id"`
	Available     bool          `json:"available"`
	Animated      bool          `json:"animated"`
}

type Properties struct {
	MaxMembers                  int64       `json:"max_members"`
	SystemChannelFlags          int64       `json:"system_channel_flags"`
	AfkTimeout                  int64       `json:"afk_timeout"`
	DiscoverySplash             interface{} `json:"discovery_splash"`
	Icon                        *string     `json:"icon"`
	Features                    []string    `json:"features"`
	LatestOnboardingQuestionID  interface{} `json:"latest_onboarding_question_id"`
	ApplicationID               interface{} `json:"application_id"`
	NsfwLevel                   int64       `json:"nsfw_level"`
	PremiumTier                 int64       `json:"premium_tier"`
	ID                          string      `json:"id"`
	Nsfw                        bool        `json:"nsfw"`
	MaxStageVideoChannelUsers   int64       `json:"max_stage_video_channel_users"`
	Description                 interface{} `json:"description"`
	ExplicitContentFilter       int64       `json:"explicit_content_filter"`
	IncidentsData               interface{} `json:"incidents_data"`
	OwnerID                     string      `json:"owner_id"`
	PublicUpdatesChannelID      interface{} `json:"public_updates_channel_id"`
	InventorySettings           interface{} `json:"inventory_settings"`
	HomeHeader                  interface{} `json:"home_header"`
	SystemChannelID             *string     `json:"system_channel_id"`
	Splash                      *string     `json:"splash"`
	VanityURLCode               *string     `json:"vanity_url_code"`
	HubType                     interface{} `json:"hub_type"`
	MfaLevel                    int64       `json:"mfa_level"`
	RulesChannelID              interface{} `json:"rules_channel_id"`
	SafetyAlertsChannelID       interface{} `json:"safety_alerts_channel_id"`
	PreferredLocale             string      `json:"preferred_locale"`
	DefaultMessageNotifications int64       `json:"default_message_notifications"`
	Name                        string      `json:"name"`
	PremiumProgressBarEnabled   bool        `json:"premium_progress_bar_enabled"`
	MaxVideoChannelUsers        int64       `json:"max_video_channel_users"`
	VerificationLevel           int64       `json:"verification_level"`
	Banner                      *string     `json:"banner"`
	AfkChannelID                interface{} `json:"afk_channel_id"`
}

type Role struct {
	UnicodeEmoji interface{} `json:"unicode_emoji"`
	Tags         Tags        `json:"tags"`
	Position     int64       `json:"position"`
	Permissions  string      `json:"permissions"`
	Name         string      `json:"name"`
	Mentionable  bool        `json:"mentionable"`
	Managed      bool        `json:"managed"`
	ID           string      `json:"id"`
	Icon         interface{} `json:"icon"`
	Hoist        bool        `json:"hoist"`
	Flags        int64       `json:"flags"`
	Color        int64       `json:"color"`
}

type Tags struct {
	BotID             *string     `json:"bot_id,omitempty"`
	PremiumSubscriber interface{} `json:"premium_subscriber"`
}

type MergedMember struct {
	UserID                     string        `json:"user_id"`
	Roles                      []interface{} `json:"roles"`
	PremiumSince               interface{}   `json:"premium_since"`
	Pending                    bool          `json:"pending"`
	Nick                       interface{}   `json:"nick"`
	Mute                       bool          `json:"mute"`
	JoinedAt                   string        `json:"joined_at"`
	Flags                      int64         `json:"flags"`
	Deaf                       bool          `json:"deaf"`
	CommunicationDisabledUntil interface{}   `json:"communication_disabled_until"`
	Avatar                     interface{}   `json:"avatar"`
}

type ReadState struct {
	Version int64   `json:"version"`
	Partial bool    `json:"partial"`
	Entries []Entry `json:"entries"`
}

type Entry struct {
	MentionCount     *int64  `json:"mention_count,omitempty"`
	LastPinTimestamp *string `json:"last_pin_timestamp,omitempty"`
	LastMessageID    *string `json:"last_message_id,omitempty"`
	ID               string  `json:"id"`
	Flags            *int64  `json:"flags,omitempty"`
	ReadStateType    *int64  `json:"read_state_type,omitempty"`
	LastAckedID      *int64  `json:"last_acked_id,omitempty"`
	BadgeCount       *int64  `json:"badge_count,omitempty"`
}

type Session struct {
	Status     string        `json:"status"`
	SessionID  string        `json:"session_id"`
	ClientInfo ClientInfo    `json:"client_info"`
	Activities []interface{} `json:"activities"`
}

type Tutorial struct {
	IndicatorsSuppressed bool     `json:"indicators_suppressed"`
	IndicatorsConfirmed  []string `json:"indicators_confirmed"`
}

type PurpleUser struct {
	Verified             bool        `json:"verified"`
	Username             string      `json:"username"`
	PurchasedFlags       int64       `json:"purchased_flags"`
	Pronouns             string      `json:"pronouns"`
	PremiumType          int64       `json:"premium_type"`
	Premium              bool        `json:"premium"`
	Phone                interface{} `json:"phone"`
	NsfwAllowed          bool        `json:"nsfw_allowed"`
	Mobile               bool        `json:"mobile"`
	MfaEnabled           bool        `json:"mfa_enabled"`
	ID                   string      `json:"id"`
	HasBouncedEmail      bool        `json:"has_bounced_email"`
	GlobalName           string      `json:"global_name"`
	Flags                int64       `json:"flags"`
	Email                interface{} `json:"email"`
	Discriminator        string      `json:"discriminator"`
	Desktop              bool        `json:"desktop"`
	Bio                  string      `json:"bio"`
	BannerColor          interface{} `json:"banner_color"`
	Banner               interface{} `json:"banner"`
	AvatarDecorationData interface{} `json:"avatar_decoration_data"`
	Avatar               interface{} `json:"avatar"`
	AccentColor          interface{} `json:"accent_color"`
}

type UserElement struct {
	Username             string      `json:"username"`
	PublicFlags          int64       `json:"public_flags"`
	ID                   string      `json:"id"`
	GlobalName           string      `json:"global_name"`
	DisplayName          string      `json:"display_name"`
	Discriminator        string      `json:"discriminator"`
	Bot                  bool        `json:"bot"`
	AvatarDecorationData interface{} `json:"avatar_decoration_data"`
	Avatar               interface{} `json:"avatar"`
}

type DGuildExperiment struct {
	Integer    *int64
	String     *string
	UnionArray []TentacledGuildExperiment
}

type TentacledGuildExperiment struct {
	FluffyGuildExperiment *FluffyGuildExperiment
	UnionArrayArrayArray  [][][]StickyGuildExperiment
}

type StickyGuildExperiment struct {
	Integer    *int64
	UnionArray []IndigoGuildExperiment
}

type IndigoGuildExperiment struct {
	Integer               *int64
	PurpleGuildExperiment *PurpleGuildExperiment
	UnionArray            []IndecentGuildExperiment
}

type IndecentGuildExperiment struct {
	Bool                  *bool
	Integer               *int64
	PurpleGuildExperiment *PurpleGuildExperiment
	UnionArray            []HilariousGuildExperiment
}

type HilariousGuildExperiment struct {
	Integer     *int64
	String      *string
	StringArray []string
}
