package main

import (
	"fmt"
)

func main() {
	BaseURL = "http://localhost:50000"

	//输入服务器端url
	//fmt.Println("Please input server url :")
	//fmt.Scanf("%s", &BaseURL)

	//检查网络连通性
	ping()

	//设置游戏模式
	var roundNumber int
	fmt.Println("Please input game round: ")
	fmt.Scanf("%d", &roundNumber)

	//发送加入游戏请求
	var userid string
	fmt.Println("Please input user id :")
	fmt.Scanf("%s", &userid)
	sendJoin(userid, roundNumber)
	fmt.Println("Login successfully")
	fmt.Println("Login ID: ", LoginedUser.Id)

	//发送选择
	for i := 0; i < roundNumber; i++ {
		SendChoice()
	}
}
