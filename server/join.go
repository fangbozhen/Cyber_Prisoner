package main

import (
	"fmt"
	"io"
	"net/http"
)

var ConnectedUserNumber int

func join(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("ReadAll error: ", err)
		return
	}

	message := convertToMsg(b)
	messageReturn := Message{}

	//请求中未提供玩家id
	if ConnectedUserNumber == 2 {
		messageReturn.Code = 100
		messageReturn.Msg = "game has already started"
	} else if message.User.Id == "" {
		messageReturn.Code = 101
		messageReturn.Msg = "should provide gamer ID"
	} else {
		userid := message.User.Id
		messageReturn.User.Id = userid
		messageReturn.User.Token = generateToken(userid)
		ConnectedUserNumber++
		RoundNumber = message.RoundNumber
		LoginedUser = append(LoginedUser, userid)
		fmt.Printf("Successfully Login User%d : %s\n", ConnectedUserNumber, userid)
	}
	data := convertToJson(messageReturn)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
