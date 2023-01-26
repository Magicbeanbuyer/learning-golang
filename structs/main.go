package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	ming := person{firstName: "ming", lastName: "li"}
	fmt.Println(ming)

	var noOne person
	fmt.Printf("noOne's fname and lname are polulated with string zero value: %+v\n", noOne)

	var xiaoHong person
	xiaoHong.firstName = "xiaohong"
	xiaoHong.lastName = "wang"
	fmt.Println(xiaoHong)
}
