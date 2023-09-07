package discord

import (
	"fmt"
	"net/url"

	"github.com/Implex-ltd/cleanhttp/cleanhttp"
	"github.com/Implex-ltd/fingerprint-client/fpclient"
	discord "github.com/Implex-ltd/ucdiscord/ucdiscord"

	"testing"
)

func TestClient_Cookies(t *testing.T) {
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
		GetCookies:  true,
		BuildNumber: 225873,
		Client:      http,
	})

	if err != nil {
		panic(err)
	}

	header := client.GetHeader(&discord.HeaderConfig{
		IsXtrack: false,
		IsSuper:  true,
	})

	u, _ := url.Parse("https://discord.com")
	fmt.Println(client.HttpClient.FormatCookies(u))
	fmt.Println(header)
}