package main

import (
	"fmt"
	"net/url"
)

func main() {
	url, err := url.ParseRequestURI("./hey")
	if err != nil {
		panic(err)
	}
	fmt.Println(url.Hostname())
}
