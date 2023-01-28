package main

import (
	"fmt"
	"net/http"
)

func main() {
	urls := []string{
		"https://github.com/",
		"https://www.linkedin.com/",
		"https://www.reddit.com/",
	}

	statusChannel := make(chan string)

	for _, url := range urls {
		go healthCheck(url, statusChannel)
	}

	for i := 0; i < 3; i++ {
		fmt.Println(<-statusChannel)
	}
}

func healthCheck(url string, channel chan string) {
	_, err := http.Get(url)
	if err != nil {
		channel <- url + " is down."
	}
	channel <- url + " is healthy! 🤟"
}

/*
messages send to a channel must all have the same data type
*/
