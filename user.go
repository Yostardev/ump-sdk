package ump_sdk

import (
	"fmt"
	"github.com/Yostardev/gf"
	"github.com/Yostardev/requests"
	"strconv"
	"time"
)

type userInfoRequest struct {
	Data *UserInfo `json:"data"`
	Msg  string    `json:"msg"`
	Code int       `json:"code"`
}

type userInfoAllRequest struct {
	Data []*UserInfo `json:"data"`
	Msg  string      `json:"msg"`
	Code int         `json:"code"`
}

type UserInfo struct {
	Id            int       `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Name          string    `json:"name"`
	Username      string    `json:"username"`
	Email         []string  `json:"email"`
	Phone         []string  `json:"phone"`
	IsApiUser     bool      `json:"is_api_user"`
	IdaasId       string    `json:"idaas_id"`
	Describe      string    `json:"describe"`
	Avatar        string    `json:"avatar"`
	FeishuUnionId string    `json:"feishu_union_id"`
	FeishuUserId  string    `json:"feishu_user_id"`
	FeishuOpenId  string    `json:"feishu_open_id"`
	OuDirectory   string    `json:"ou_directory"`
	RoleIds       []int     `json:"role_ids"`
	Roles         []struct {
		Id            int    `json:"id"`
		Name          string `json:"name"`
		Description   string `json:"description"`
		ApplicationId int    `json:"application_id"`
	} `json:"roles"`
	AuthorityIds []int `json:"authority_ids"`
	Authorities  []struct {
		Id            int    `json:"id"`
		Name          string `json:"name"`
		Description   string `json:"description"`
		Obj           string `json:"obj"`
		Act           string `json:"act"`
		ApplicationId int    `json:"application_id"`
	} `json:"authorities"`
}

func (c *Client) GetUserInfo() (*UserInfo, error) {
	res, err := requests.New().SetUrl(gf.StringJoin(c.serverURL, "/ump/api/v1/user/self")).AddQuery("application_id", strconv.Itoa(c.applicationID)).AddHeader("Authorization", c.token).Get()
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("get user info failed, status code: %d, response data: %s", res.StatusCode, res.Body.String())
	}

	var resp userInfoRequest
	err = res.Body.JsonBind(&resp)
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (c *Client) GetAllUser() ([]*UserInfo, error) {
	res, err := requests.New().SetUrl(gf.StringJoin(c.serverURL, "/ump/api/v1/user/all")).AddHeader("Authorization", c.token).Get()
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("get user info failed, status code: %d, response data: %s", res.StatusCode, res.Body.String())
	}

	var resp userInfoAllRequest
	err = res.Body.JsonBind(&resp)
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (c *Client) GetUserByUsername(username string) (*UserInfo, error) {
	res, err := requests.New().SetUrl(gf.StringJoin(c.serverURL, "/ump/api/v1/user/all")).AddQuery("username", username).AddHeader("Authorization", c.token).Get()
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("get user info failed, status code: %d, response data: %s", res.StatusCode, res.Body.String())
	}

	var resp userInfoAllRequest
	err = res.Body.JsonBind(&resp)
	if err != nil {
		return nil, err
	}

	if len(resp.Data) != 1 {
		return nil, fmt.Errorf("get user info failed, found %d item", len(resp.Data))
	}

	return resp.Data[0], nil
}

func (c *Client) UpdateUserAvatar(avatar string) error {
	res, err := requests.New().SetUrl(gf.StringJoin(c.serverURL, "/ump/api/v1/user/avatar")).SetJsonBody(map[string]string{
		"avatar": avatar,
	}).AddHeader("Authorization", c.token).Post()
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return fmt.Errorf("update user avatar failed, status code: %d, response data: %s", res.StatusCode, res.Body.String())
	}
	return nil
}
