package main

import "fmt"

type contactInfo struct {
	email  string
	mobile int
}

type person struct {
	firstName   string
	lastName    string
	contactInfo // declare a field with name contactInfo and type contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p) // %+v prints both the field name and the value
}

func structOperations() {
	ming := person{firstName: "ming", lastName: "li"}
	fmt.Println(ming)

	var noOne person
	fmt.Printf("noOne's fname and lname are polulated with string zero value: %+v\n", noOne)

	var xiaoHong person
	xiaoHong.firstName = "xiaohong"
	xiaoHong.lastName = "wang"
	xiaoHong.contactInfo.mobile = 123
	xiaoHong.print()
}

func main() {
	structOperations()
}
