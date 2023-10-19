package main

import (
	"fmt"

	"github.com/learning-go-book/package_example/formatter"
	"github.com/learning-go-book/package_example/lala"
)

func main() {
	num := math.Double(2)
	output := print.Format(num)
	fmt.Println(output)
}
