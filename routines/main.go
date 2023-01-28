package main

import (
	"fmt"
	"net/http"
	"time"
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

	for link := range statusChannel {
		go func(copiedLink string) {
			/*
				putting the pause within the function literal, a.k.a. lambda function, makes sure that a child routine
				is spawned immediately after a message is sent to the channel, and the child routine pauses before it
				sends another request to the server.
			*/
			time.Sleep(1 * time.Second)
			healthCheck(copiedLink, statusChannel) // mind blown 🤯
		}(link)
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
