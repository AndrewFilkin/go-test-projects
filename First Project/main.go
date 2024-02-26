package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	go responseSize("https://example.com/")
	go responseSize("https://golang.org/")
	go responseSize("https://golang.org/doc")

	time.Sleep(time.Second)
}
func responseSize(url string) {
	fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(body))
	fmt.Println(response.StatusCode)
	time.Sleep(time.Second)
}
