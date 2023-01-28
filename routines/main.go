package main

import (
	"fmt"
	"net/http"
)

func main() {
	urls := []string{
		"https://github.com/",
		"https://www.bing.com/",
		"https://www.google.com/",
	}

	statusChannel := make(chan string)

	for _, url := range urls {
		go healthCheck(url, statusChannel)
	}

	for {
		go healthCheck(<-statusChannel, statusChannel) // mind blown 🤯
	}
}

func healthCheck(url string, channel chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Printf("%v is down.\n", url)
		channel <- url
	}
	fmt.Printf("%v is healthy! 🤟\n", url)
	channel <- url
}

/*
messages send to a channel must all have the same data type
*/
