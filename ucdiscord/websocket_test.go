package ucdiscord

import (
	"fmt"
	"testing"
)

func TestNewWebsocket(t *testing.T) {
	type args struct {
		Token string
	}
	tests := []struct {
		name    string
		args    args
		want    *ClientWebsocket
		wantErr bool
	}{
		{
			name: "uwu",
			args: args{
				Token: "MTE1MTQwMzIwNjc4NjY5OTMyNQ.GMb5o_.Xionc9mZVH6NxP73a0Sg4mbDD3H_EGV7F2HWcQ",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ws, err := NewWebsocket(tt.args.Token, &XProp{
				Os:                "Windows",
				Browser:           "Chrome",
				Device:            "",
				SystemLocale:      "fr-FR",
				BrowserUserAgent:  "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36",
				BrowserVersion:    "117.0.0.0",
				OsVersion:         "10",
				ReleaseChannel:    "stable",
				ClientBuildNumber: 227102,
			})
			if err != nil {
				panic(err)
			}

			ws.Debug = true
			ws.LogNonImplemented = true

			if err := ws.Login(); err != nil {
				panic(err)
			}

			if err := ws.UpdateStatus("online", "Minecraft"); err != nil {
				panic(err)
			}

			fmt.Println("Logged in !", ws.ReadyData.SessionID)
		})
	}
}
