package main

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func SendChoice() {
	url := BaseURL + "/playground"
	contentType := "application/json"
	messageSend := Message{
		User:   LoginedUser,
		Choice: makeChoice(),
	}

	data := convertToJson(messageSend)
	resp, err := http.Post(url, contentType, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("sendChoice Post error: ", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("sendChoice ReadAll error: ", err)
		os.Exit(1)
	}

	messageReturn := convertToMsg(b)
	if messageReturn.Code != 0 {
		fmt.Println("error: ", messageReturn.Msg)
		os.Exit(1)
	}

	fmt.Println("Your choice: ", messageReturn.Round.YourChoice)
	fmt.Println("Op choice: ", messageReturn.Round.OpChoice)
	fmt.Println("Your result: ", messageReturn.Round.YourResult)
	fmt.Println("Op result: ", messageReturn.Round.OpResult)
	fmt.Println("----------------------------------------")
}

func makeChoice() string {
	rand.Seed(time.Now().UnixNano())
	if rand.Int()%2 == 0 {
		return "confess"
	} else {
		return "deny"
	}

}
