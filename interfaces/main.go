package main

func main() {
	wrapper()
	html_writer_wrapper()
	li_wrapper()
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
