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
		//Token:       "MTE0NjU4NTk4MDA5MDk5NDczOA.G9mS8b.rfnoKAdQsBNZJVc4hGRzYokf2lFbYQAZX5vAeg",
		GetCookies:  true,
		BuildNumber: 224244,
		Client:      http,
	})

	if err != nil {
		panic(err)
	}

	client.WsConnect()

	type args struct {
		config *discord.RegisterConfig
	}

	tests := []struct {
		name string
		c    *discord.Client
		args args
	}{
		{
			name: "add friend",
			args: args{
				config: &discord.RegisterConfig{
					Username: "hcaptcha",
					InviteCode: "uwu",
					CaptchaKey: "bop",
				},
			},
			c: client,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Register(tt.args.config)
			fmt.Println(got, err)
		})
	}
}
