package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	//检查网络的连通性
	mux.HandleFunc("/ping", ping)

	//玩家登陆
	mux.HandleFunc("/join", join)

	//接受客户端的博弈选择
	mux.HandleFunc("/playground", playground)

	server := &http.Server{
		Addr:    ":50000",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
