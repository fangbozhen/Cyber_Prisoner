package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

var (
	cnt         int
	round       [2]Round
	wait        sync.WaitGroup
	start       time.Time
	RoundNumber int
	RoundCount  int
	result      [2][]Round
)

func playground(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("ReadAll error: ", err)
		return
	}

	message := convertToMsg(b)
	if ok := checkToken(message.User); !ok {
		fmt.Printf("%s token error", message.User.Id)
		return
	}

	var tmp int
	messageReturn := Message{
		User:   message.User,
		Choice: message.Choice,
	}
	if message.Choice != "confess" && message.Choice != "deny" {
		messageReturn.Code = 200
		messageReturn.Msg = "undefined choice"
	} else if false {
		messageReturn.Code = 201
		messageReturn.Msg = "opponent did not request"
	} else {
		round[cnt].YourChoice = message.Choice
		fmt.Printf("Logined user: %s, %d\n", message.User.Id, cnt)
		if cnt == 0 {
			tmp = cnt
			cnt++
			start = time.Now()
			wait.Add(1)
			wait.Wait()
		} else if cnt == 1 {
			setRound(0)
			setRound(1)
			tmp = cnt
			RoundCount++
		}
		messageReturn.Round = round[tmp]
	}
	data := convertToJson(messageReturn)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	if tmp == 1 {
		wait.Done()
		cnt = 0
	}

	var index int
	if message.User.Id == LoginedUser[0] {
		index = 0
	} else {
		index = 1
	}
	result[index] = append(result[index], round[tmp])
	if RoundCount == RoundNumber {
		printResult(0)
		printResult(1)
	}
}

func checkToken(user User) bool {
	if parseToken(user.Token)["userid"] == user.Id {
		return true
	} else {
		return false
	}
}

func setRound(index int) {
	round[index].OpChoice = round[1-index].YourChoice
	if round[index].YourChoice == "confess" && round[index].OpChoice == "confess" {
		round[index].YourResult = 8
		round[index].OpResult = 8
	} else if round[index].YourChoice == "confess" && round[index].OpChoice == "deny" {
		round[index].YourResult = 0
		round[index].OpResult = 10
	} else if round[index].YourChoice == "deny" && round[index].OpChoice == "confess" {
		round[index].YourResult = 10
		round[index].OpResult = 0
	} else {
		round[index].YourResult = 1
		round[index].OpResult = 1
	}
}

func printResult(index int) {
	fmt.Printf("user %d: %s\n", index, LoginedUser[index])
	for i := 0; i < RoundNumber; i++ {
		fmt.Println("Your choice: ", result[index][i].YourChoice)
		fmt.Println("Op choice: ", result[index][i].OpChoice)
		fmt.Println("Your result: ", result[index][i].YourResult)
		fmt.Println("Op result: ", result[index][i].OpResult)
		fmt.Println("-----------------------------------------")
	}
}
