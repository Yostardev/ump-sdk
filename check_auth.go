package ump_sdk

import (
	"fmt"
	"github.com/Yostardev/gf"
	"github.com/Yostardev/requests"
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
	res, err := requests.New().SetUrl(gf.StringJoin(c.serverURL, "/ump/api/v1/check")).SetJsonBody(&checkAuthRequest{
		ApplicationId: c.applicationID,
		Obj:           obj,
		Act:           act,
	}).AddHeader("Authorization", c.token).Post()
	if err != nil {
		return false, err
	}
	if res.StatusCode != 200 {
		return false, fmt.Errorf("check auth failed, status code: %d, response data: %s", res.StatusCode, res.Body.String())
	}

	var resp checkAuthResponse
	err = res.Body.JsonBind(&resp)
	if err != nil {
		return false, err
	}

	return resp.Data.Auth, nil
}
