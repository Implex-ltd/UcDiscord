package ucdiscord

import (
	"encoding/json"
	"strings"
)

func AddBase64Padding(base64String string) string {
	paddingCount := 4 - (len(base64String) % 4)
	if paddingCount != 4 {
		padding := strings.Repeat("=", paddingCount)
		return base64String + padding
	}
	return base64String
}

func CustomMarshalProp(props CtxProp) ([]byte, error) {
	data := make(map[string]interface{})

	if props.Location != "" {
		data["location"] = props.Location
	}
	if props.LocationGuildID != "" {
		data["location_guild_id"] = props.LocationGuildID
	}
	if props.LocationChannelID != "" {
		data["location_channel_id"] = props.LocationChannelID
	}
	if props.LocationChannelType != 0 {
		data["location_channel_type"] = props.LocationChannelType
	}

	return json.Marshal(data)
}

func CleanSetProfil(P *SetProfilPayload) map[string]interface{} {
	data := make(map[string]interface{})

	if P.Bio != "" {
		data["bio"] = P.Bio
	}
	if P.Pronouns != "" {
		data["pronouns"] = P.Pronouns
	}
	if P.AccentColor != 0 {
		data["accent_color"] = P.AccentColor
	}

	return data
}

func CleanPatchUser(P *PatchPayload) map[string]interface{} {
	data := make(map[string]interface{})

	if P.Date != "" {
		data["date_of_birth"] = P.Date
	}

	if P.Email != "" {
		data["email"] = P.Email
	}

	if P.Password != "" {
		data["password"] = P.Password
	}

	if P.Avatar != "" {
		data["avatar"] = P.Avatar
	}

	if P.DisplayName != "" {
		data["global_name"] = P.DisplayName
	}

	return data
}
