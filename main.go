package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	for {
		makeGetRequest("http://10.103.170.46:9090/")
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

	fmt.Println(string(body), time.Now())

	defer resp.Body.Close()
}
