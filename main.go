package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	for {
		makeGetRequest("http://localhost:9090")
		time.Sleep(3 * time.Second)
	}
}

func makeGetRequest(url string) {
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body), time.Now().Format("2006-01-02 15:04:05"))

	defer resp.Body.Close()
}
