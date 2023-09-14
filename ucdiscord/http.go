package ucdiscord

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/Implex-ltd/cleanhttp/cleanhttp"
)

func (C *Client) Do(config Request) (*Response, error) {
	opt := cleanhttp.RequestOption{
		Url:    config.Endpoint,
		Method: config.Method,
		Header: config.Header,
	}

	if config.Method != "GET" && config.Body != nil {
		body, err := json.Marshal(config.Body)
		if err != nil {
			return nil, err
		}

		opt.Body = bytes.NewReader(body)
	}

	req, err := C.Config.Http.Do(opt)
	if err != nil {
		return nil, err
	}

	defer req.Body.Close()

	if req.Header.Get("Content-Type") == "application/json" {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(body, &config.Response); err != nil {
			return nil, err
		}
	}

	return &Response{
		Status: req.StatusCode,
	}, nil
}
