package ump_sdk

import (
	"time"

	"github.com/bytedance/sonic"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	applicationID int
	restyClient   *resty.Client
}

func NewClient(url string, applicationID int, token string) *Client {
	return &Client{
		applicationID: applicationID,
		restyClient: resty.New().
			SetJSONMarshaler(sonic.Marshal).
			SetJSONUnmarshaler(sonic.Unmarshal).
			SetTimeout(10*time.Second).
			SetAuthToken(token).
			SetAuthScheme("").
			SetHeader("Content-Type", "application/json").
			SetBaseURL(url),
	}
}
