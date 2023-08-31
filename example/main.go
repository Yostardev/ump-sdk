package main

import (
	"fmt"
	"github.com/Yostardev/ump-sdk"
)

var (
	url   = "https://dev-opsump.yostar.net"
	token = "QcddoPbmrnTLMgxVHFwysb4yd6Se4edkhcTyippQMo3uS4nOaryXmH8DgcSEDekowmXGSADlIZ17YdVTkjMx82iuExwqyYHIADNa"
)

func main() {
	client := ump_sdk.NewClient(url, 1, token)

	auth, err := client.CheckAuth("user", "get")
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}
	fmt.Println(auth)

	userInfo, err := client.GetUserInfo()
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}
	fmt.Printf("%+v\n", userInfo)

	authInfo, err := client.CreateAuthority("test1", "test1", "user1", "get")
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}

	fmt.Printf("%+v\n", authInfo)
}
