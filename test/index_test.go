package discord

import (
	"fmt"

	"github.com/Implex-ltd/cleanhttp/cleanhttp"
	"github.com/Implex-ltd/fingerprint-client/fpclient"
	discord "github.com/Implex-ltd/ucdiscord/ucdiscord"

	"testing"
)

func TestClient_SendFriend(t *testing.T) {
	// Load fingerprint
	fp, err := fpclient.LoadFingerprint(&fpclient.LoadingConfig{
		FilePath: "./chrome.json",
	})
	if err != nil {
		panic(err)


	}

	// Load HTTP client
	http, err := cleanhttp.NewCleanHttpClient(&cleanhttp.Config{
		BrowserFp: fp,
	})
	if err != nil {
		return
	}

	// Create discord session
	client, err := discord.NewClient(&discord.ClientConfig{
		Token:       "MTE0NjgzNDY5NTc3NTA2NDA3NQ.G9edxF.1fuXWK-QGGqb6tz0pMNoVttOPCs7ncaSAFTmFA",
		GetCookies:  true,
		BuildNumber: 224244,
		Client:      http,
	})

	if err != nil {
		panic(err)
	}

	client.WsConnect()

	type args struct {
		config *discord.FriendConfig
	}

	tests := []struct {
		name string
		c    *discord.Client
		args args
	}{
		{
			name: "add friend",
			args: args{
				config: &discord.FriendConfig{
					Username: "hcaptcha",
				},
			},
			c: client,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.c.SendFriend(tt.args.config)
			fmt.Println(got, got1, err)
		})
	}
}
