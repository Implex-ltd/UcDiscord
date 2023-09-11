package discord

import (
	"fmt"

	"github.com/Implex-ltd/cleanhttp/cleanhttp"
	"github.com/Implex-ltd/fingerprint-client/fpclient"
	discord "github.com/Implex-ltd/ucdiscord/ucdiscord"

	"testing"
)

func TestClient_Register(t *testing.T) {
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
		Log:       true,
	})
	if err != nil {
		return
	}

	// Create discord session
	client, err := discord.NewClient(&discord.ClientConfig{
		GetCookies:  true,
		BuildNumber: 226220,
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
					Username:   "hcaptcha",
					InviteCode: "lol",
					CaptchaKey: "test",
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

func TestClient_Join(t *testing.T) {
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
		Log:       true,
	})
	if err != nil {
		return
	}

	// Create discord session
	client, err := discord.NewClient(&discord.ClientConfig{
		Token:       "MTE1MDYwMTA1MTc2NzA1ODU0Mg.GYT9Ze.VTRQqDmw2oXFoMje3b6cLlIPy91Hi_sCreDtIk",
		GetCookies:  true,
		BuildNumber: 226220,
		Client:      http,
	})

	if err != nil {
		panic(err)
	}

	client.WsConnect()

	type args struct {
		config *discord.JoinConfig
	}

	tests := []struct {
		name string
		c    *discord.Client
		args args
	}{
		{
			name: "add friend",
			args: args{
				config: &discord.JoinConfig{
					InviteCode: "zaSphzfm",
					GuildID:    "1149530064925499402",
					ChannelID:  "1149530065529491571",
				},
			},
			c: client,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.JoinGuild(tt.args.config)
			fmt.Println(got, err)
		})
	}
}
