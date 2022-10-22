package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var BaseURL string
var LoginedUser []string

type Message struct {
	Code        int    `json:"code"`
	Msg         string `json:"msg"`
	Data        string `json:"data"`
	User        User
	Choice      string `json:"choice,omitempty""`
	Round       Round
	RoundNumber int `json:"round_number,omitempty"`
}

type User struct {
	Id    string `json:"id,omitempty"`
	Token string `json:"user_token,omitempty"`
}

type Round struct {
	YourChoice string `json:"your_choice,omitempty"`
	OpChoice   string `json:"op_choice,omitempty"`
	YourResult int    `json:"your_result,omitempty"`
	OpResult   int    `json:"op_result,omitempty"`
}

func convertToJson(message Message) []byte {
	data, err := json.MarshalIndent(message, "", "	")
	if err != nil {
		fmt.Println("JSON marshal failed: ", err)
		os.Exit(1)
	}
	return data
}

func convertToMsg(data []byte) Message {
	message := Message{}
	if err := json.Unmarshal(data, &message); err != nil {
		fmt.Println("JSON unmarshal failed: ", err)
		os.Exit(1)
	}
	return message
}
