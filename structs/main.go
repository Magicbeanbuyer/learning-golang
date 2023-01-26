package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func structOperations() {
	ming := person{firstName: "ming", lastName: "li"}
	fmt.Println(ming)

	var noOne person
	fmt.Printf("noOne's fname and lname are polulated with string zero value: %+v\n", noOne)

	var xiaoHong person
	xiaoHong.firstName = "xiaohong"
	xiaoHong.lastName = "wang"
	fmt.Println(xiaoHong)
}

// https://www.youtube.com/watch?v=2ybLD6_2gKM
func learnPointer() {
	var x int = 42   // variable named x of type integer is set to 4
	var pX *int = &x // variable named pX of type integer POINTER is set to the address of x
	// dereferencing, indirecting
	var y int = *pX // variable named y of type integer is set to the thing pointed to by pX, hence 42
	// value 42 is COPIED to the address of y
	y = -1 // change the value of y, not the value of x
	fmt.Printf("x: %v, y: %v, pX: %v\n", x, y, pX)
	*pX = 1000 // change the value of the thing pointed to by pX to 1000
	fmt.Printf("x: %v, y: %v, pX: %v\n", x, y, pX)
	pX = &y // change the value pX to the address of y
	fmt.Printf("x: %v, y: %v, pX: %v\n", x, y, pX)
	*pX = -50 // change the value of the thing pointed to by pX to -50
	fmt.Printf("x: %v, y: %v, pX: %v\n", x, y, pX)
}

func main() {
	structOperations()
	learnPointer()
}
