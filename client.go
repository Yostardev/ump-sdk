package ump_sdk

type Client struct {
	serverURL     string
	applicationID int
	token         string
}

func NewClient(url string, applicationID int, token string) *Client {
	return &Client{
		serverURL:     url,
		applicationID: applicationID,
		token:         token,
	}
}
