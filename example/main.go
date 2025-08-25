package main

import (
	"fmt"

	"github.com/Yostardev/ump-sdk"
)

var (
	url   = "https://dev-opsump.yostar.net"
	token = "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJ4dXRpYW5jaGVuZyIsIm5hbWUiOiLlvpDlpKnmiJAiLCJpc3MiOiJZb1N0YXIgVXNlciBNYW5hZ2UgUGxhdGZvcm0iLCJleHAiOjE3NTYzNDU4MTIsIm5iZiI6MTc1NjA4NjYxMiwiaWF0IjoxNzU2MDg2NjEyLCJqdGkiOiI0NzkwMGJjNi0zMjk2LTQ2ZTgtYWY0NC0yYTRiNTFiNTYxZTcifQ.EVrjwKxdA58pIvbYK6y7LSQjUVL_Ml5KaSMS5OlON3Vp2nulZ7NneojHssJAIi5-OtgAQXRKSV77dt4F4K377A"
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
