package main

import (
	"log"

	"github.com/Implex-ltd/cleanhttp/cleanhttp"
	"github.com/Implex-ltd/fingerprint-client/fpclient"
	u "github.com/Implex-ltd/ucdiscord/ucdiscord"
)

func main() {
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
		Proxy:     "http://0.0.0.0:8080",
	})
	if err != nil {
		return
	}

	// Create discord session
	client, err := u.NewClient(&u.ClientConfig{
		Token:       "FOOBAR123",
		GetCookies:  true,
		BuildNumber: 00000,
		Client:      http,
	})

	if err != nil {
		panic(err)
	}

	// join a server
	resp, err := client.JoinGuild(&u.JoinConfig{
		InviteCode: "supercode",
		GuildID:    "00000000000000",
		ChannelID:  "00000000000000",
	})

	if err != nil {
		panic(err)
	}

	// access captcha: resp.CaptchaKey, resp.CaptchaRqdata, resp.CaptchaRqtoken

	log.Printf("Joined %s (%s)\n", resp.Guild.Name, resp.Guild.ID)

	// create a account
	response, err := client.Register(&u.RegisterConfig{
		Username: "foo",
		InviteCode: "bar",
		CaptchaKey: "P0_..",
	})

	if err != nil {
		panic(err)
	}

	log.Printf("Created account %s\n", response.Token)
}
