package ump_sdk

import (
	"fmt"
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
	var resp authorityResponse
	res, err := c.restyClient.R().SetResult(&resp).SetBody(&authorityRequest{
		Name:          name,
		Describe:      describe,
		ApplicationID: c.applicationID,
		Obj:           obj,
		Act:           act,
	}).Post("/ump/api/v1/authority")
	if err != nil {
		return nil, err
	}
	if res.StatusCode() != 200 {
		return nil, fmt.Errorf("create authority failed, status code: %d, response data: %s", res.StatusCode(), res.String())
	}

	return resp.Data, nil
}

func (c *Client) UpdateAuthority(id int, name, describe, obj, act string) (*Authority, error) {
	var resp authorityResponse
	res, err := c.restyClient.R().SetResult(&resp).SetBody(&authorityRequest{
		Name:          name,
		Describe:      describe,
		ApplicationID: c.applicationID,
		Obj:           obj,
		Act:           act,
	}).SetPathParam("id", strconv.Itoa(id)).Put("/ump/api/v1/authority/{id}")
	if err != nil {
		return nil, err
	}
	if res.StatusCode() != 200 {
		return nil, fmt.Errorf("update authority failed, status code: %d, response data: %s", res.StatusCode(), res.String())
	}

	return resp.Data, nil
}

func (c *Client) DeleteAuthority(id int) error {
	res, err := c.restyClient.R().SetPathParam("id", strconv.Itoa(id)).Delete("/ump/api/v1/authority/{id}")
	if err != nil {
		return err
	}
	if res.StatusCode() != 200 {
		return fmt.Errorf("delete authority failed, status code: %d, response data: %s", res.StatusCode(), res.String())
	}

	return nil
}

func (c *Client) GetAuthorityByObjAndAct(obj, act string) ([]*Authority, error) {
	var resp authorityListResponse
	res, err := c.restyClient.R().SetResult(&resp).SetQueryParams(map[string]string{
		"application_id": strconv.Itoa(c.applicationID),
		"obj":            obj,
		"act":            act,
	}).Get("/ump/api/v1/authority/all")
	if err != nil {
		return nil, err
	}
	if res.StatusCode() != 200 {
		return nil, fmt.Errorf("get authority failed, status code: %d, response data: %s", res.StatusCode(), res.String())
	}

	return resp.Data, nil
}
