package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type htmlWriter struct {
}

func (htmlWriter) Write(bs []byte) (n int, err error) {
	fmt.Printf("HTML size: %v\n", len(bs))
	return len(bs), nil
}

func html_writer_wrapper() {
	resp, err := http.Get("https://www.google.de/")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	hw := htmlWriter{}
	n, err := io.Copy(hw, resp.Body)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(n)
}
