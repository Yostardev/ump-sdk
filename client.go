package ump_sdk

type Client struct {
	serverURL     string
	applicationID int
	token         string
}

func NewClient(applicationID int, token string) *Client {
	return &Client{
		serverURL:     "https://opsump.youstar.net",
		applicationID: applicationID,
		token:         token,
	}
}

func (c *Client) SetServerURL(url string) *Client {
	c.serverURL = url
	return c
}
