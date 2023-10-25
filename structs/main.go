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
	fmt.Printf("%+v\n", p) // %+v prints both the field name and the value
}

func (p *person) update_mobile(newNumber int) {
	// p is a COPY of the pointer from the outer scope, not the outer scope pointer itself
	(*p).contactInfo.mobile = newNumber
}

func main() {
	ming := person{firstName: "ming", lastName: "li"}
	fmt.Println(ming)

	var noOne person
	fmt.Printf("noOne's fname and lname are polulated with string zero value: %+v\n", noOne)

	var xiaoHong person
	xiaoHong.firstName = "xiaohong"
	xiaoHong.lastName = "wang"
	xiaoHong.contactInfo.mobile = 123
	xiaoHong.print()
	xiaoHong.update_mobile(456)
	xiaoHong.print()

	var ann person
	ann.firstName = "ann"
	ann.print()

	person_empty_ptr := &person{}
	fmt.Println(person_empty_ptr == nil) // false
	person_empty_ptr.print()

	var person_nil_ptr *person
	fmt.Println(person_nil_ptr == nil) // true
	// person_nil_ptr.print() panic: runtime error: invalid memory address or nil pointer dereference
}

/*
cannot iterate through a struct in a for loop, because the keys are not indexed
value type
*/
