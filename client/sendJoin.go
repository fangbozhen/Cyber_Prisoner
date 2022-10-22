package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func sendJoin(userid string, roundNumber int) {
	url := BaseURL + "/join"
	contentType := "application/json"
	messageSend := Message{
		User: User{
			Id: userid,
		},
		RoundNumber: roundNumber,
	}

	data := convertToJson(messageSend)
	resp, err := http.Post(url, contentType, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("Post error: ", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ReadAll error: ", err)
		os.Exit(1)
	}

	messageReturn := convertToMsg(b)
	if messageReturn.Code != 0 {
		fmt.Println("Login error: ", messageReturn.Msg)
		os.Exit(1)
	}
	LoginedUser = messageReturn.User
}
