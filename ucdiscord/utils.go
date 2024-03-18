package ucdiscord

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"github.com/andybalholm/brotli"
	http "github.com/bogdanfinn/fhttp"
)

// Thanks to https://github.com/V4NSH4J/discord-mass-DM-GO ✨
func Snowflake() int {
	snowflake := strconv.FormatInt((time.Now().UTC().UnixNano()/1000000)-1420070400000, 2) + "0000000000000000000000"
	nonce, _ := strconv.ParseInt(snowflake, 2, 64)
	return int(nonce)
}

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

	if P.Username != "" {
		data["username"] = P.Username
	}

	if P.Avatar != "" {
		data["avatar"] = P.Avatar
	}

	if P.DisplayName != "" {
		data["global_name"] = P.DisplayName
	}

	return data
}

// https://github.com/V4NSH4J/discord-mass-DM-GO/blob/main/utilities/encryption.go
func ReadBody(resp *http.Response) ([]byte, error) {

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.Header.Get("Content-Encoding") == "gzip" {
		gzipreader, err := zlib.NewReader(bytes.NewReader(body))
		if err != nil {
			return nil, err
		}
		gzipbody, err := ioutil.ReadAll(gzipreader)
		if err != nil {
			return nil, err
		}
		return gzipbody, nil
	}

	if resp.Header.Get("Content-Encoding") == "br" {
		brreader := brotli.NewReader(bytes.NewReader(body))
		brbody, err := ioutil.ReadAll(brreader)
		if err != nil {
			fmt.Println(string(brbody))
			return nil, err
		}

		return brbody, nil
	}
	return body, nil
}
