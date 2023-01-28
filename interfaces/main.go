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
	fmt.Println(string(bs))
	fmt.Printf("HTML size: %v\n", len(bs))
	return len(bs), nil
}

func main() {
	wrapper()
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

/*
var body []int -> WRONG, initialize a zero length slice, and the read function below does not resize the slice.
bBody := make([]byte, 99999)
n, _ := resp.Body.Read(bBody)
fmt.Printf("%+v, %+v\n\n", resp, n)
fmt.Println(string(body))

type ReadCloser interface {
	Reader // interface
	Closer // interface
}

to qualify as an implementation of ReadCloser interface, the implemented type must satisfy both the Reader interface and
the Closer interface
*/
