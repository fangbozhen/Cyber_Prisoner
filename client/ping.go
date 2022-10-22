package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func ping() {
	url := BaseURL + "/ping"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Reading response error: ", err)
		os.Exit(1)
	}

	message := convertToMsg(data)

	if message.Code != 0 {
		fmt.Println("connected failed")
		os.Exit(1)
	}
}
