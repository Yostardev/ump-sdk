package ump_sdk

import (
	"fmt"
)

type checkAuthRequest struct {
	ApplicationId int    `json:"application_id"`
	Obj           string `json:"obj"`
	Act           string `json:"act"`
}

type checkAuthResponse struct {
	Data struct {
		Auth bool `json:"auth"`
	} `json:"data"`
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

func (c *Client) CheckAuth(obj, act string) (bool, error) {
	var resp checkAuthResponse
	res, err := c.restyClient.R().SetResult(&resp).SetBody(&checkAuthRequest{
		ApplicationId: c.applicationID,
		Obj:           obj,
		Act:           act,
	}).Post("/ump/api/v1/check")
	if err != nil {
		return false, err
	}
	if res.StatusCode() != 200 {
		return false, fmt.Errorf("check auth failed, status code: %d, response data: %s", res.StatusCode(), res.String())
	}

	return resp.Data.Auth, nil
}
