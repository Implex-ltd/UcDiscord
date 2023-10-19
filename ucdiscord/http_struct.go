package ucdiscord

import (
	"fmt"

	http "github.com/bogdanfinn/fhttp"
)

var (
	VERSION  = 9
	ENDPOINT = fmt.Sprintf("https://discord.com/api/v%d", VERSION)
)

type Request struct {
	Body     interface{}
	Response interface{}
	Endpoint string
	Header   http.Header
	Method   string
}

type Response struct {
	Status int
}
