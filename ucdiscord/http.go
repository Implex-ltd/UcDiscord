package ucdiscord

import (
	"bytes"
	"encoding/json"

	"github.com/Implex-ltd/cleanhttp/cleanhttp"
)

func (c *Client) Do(config Request) (*Response, error) {
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
	
	req, err := c.Config.Http.Do(opt)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	
	if req.Header.Get("Content-Type") == "application/json" {
		err := json.NewDecoder(req.Body).Decode(&config.Response)
		if err != nil {
			return nil, err
		}
	}

	return &Response{
		Status: req.StatusCode,
		Body: req.Body,
	}, nil
}
