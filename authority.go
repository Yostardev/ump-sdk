package ump_sdk

import (
	"fmt"
	"github.com/Yostardev/gf"
	"github.com/Yostardev/requests"
	"strconv"
	"time"
)

type authorityRequest struct {
	Name          string `json:"name"`
	Describe      string `json:"describe"`
	ApplicationID int    `json:"application_id"`
	Obj           string `json:"obj"`
	Act           string `json:"act"`
}

type authorityResponse struct {
	Data *Authority `json:"data"`
	Msg  string     `json:"msg"`
	Code int        `json:"code"`
}

type Authority struct {
	Id            int       `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Name          string    `json:"name"`
	Describe      string    `json:"describe"`
	Obj           string    `json:"obj"`
	Act           string    `json:"act"`
	ApplicationId int       `json:"application_id"`
}

type authorityListResponse struct {
	Data []*Authority `json:"data"`
	Msg  string       `json:"msg"`
	Code int          `json:"code"`
}

func (c *Client) CreateAuthority(name, describe, obj, act string) (*Authority, error) {
	res, err := requests.New().SetUrl(gf.StringJoin(c.serverURL, "/ump/api/v1/authority")).SetJsonBody(&authorityRequest{
		Name:          name,
		Describe:      describe,
		ApplicationID: c.applicationID,
		Obj:           obj,
		Act:           act,
	}).AddHeader("Authorization", c.token).Post()
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("create authority failed, status code: %d, response data: %s", res.StatusCode, res.Body.String())
	}

	var resp authorityResponse
	err = res.Body.JsonBind(&resp)
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (c *Client) UpdateAuthority(id int, name, describe, obj, act string) (*Authority, error) {
	res, err := requests.New().SetUrl(gf.StringJoin(c.serverURL, "/ump/api/v1/authority/", strconv.Itoa(id))).SetJsonBody(&authorityRequest{
		Name:          name,
		Describe:      describe,
		ApplicationID: c.applicationID,
		Obj:           obj,
		Act:           act,
	}).AddHeader("Authorization", c.token).Put()
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("update authority failed, status code: %d, response data: %s", res.StatusCode, res.Body.String())
	}

	var resp authorityResponse
	err = res.Body.JsonBind(&resp)
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (c *Client) DeleteAuthority(id int) error {
	res, err := requests.New().SetUrl(gf.StringJoin(c.serverURL, "/ump/api/v1/authority/", strconv.Itoa(id))).
		AddHeader("Authorization", c.token).
		Delete()
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("delete authority failed, status code: %d, response data: %s", res.StatusCode, res.Body.String())
	}

	return nil
}

func (c *Client) GetAuthorityByObjAndAct(obj, act string) ([]*Authority, error) {
	res, err := requests.New().SetUrl(gf.StringJoin(c.serverURL, "/ump/api/v1/authority/all")).
		AddQuery("application_id", strconv.Itoa(c.applicationID)).
		AddQuery("obj", obj).
		AddQuery("act", act).
		AddHeader("Authorization", c.token).Get()
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("get authority failed, status code: %d, response data: %s", res.StatusCode, res.Body.String())
	}

	var resp authorityListResponse
	err = res.Body.JsonBind(&resp)
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}
