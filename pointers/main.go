package main

import "fmt"

func zero(pI *int) {
	*pI = 0
}

type myInt int

func (m *myInt) zeroMyInt() {
	*m = 0
}

func main() {
	var x int = 42 // variable named x of type integer is set to 4

	var pX *int = &x // variable named pX of type integer POINTER is set to the memory address of x

	var pointerToPx **int = &pX // variable named pointerToPx of type "integer-pointer" pointer is set to the memory
	// address of pX

	// dereferencing, indirecting
	var y int = *pX // variable named y of type integer is set to the thing pointed to by pX, hence 42. value 42 is
	// COPIED to the address of y

	y = -1 // change the value of y, not the value of x
	fmt.Printf("x: %v, y: %v, pX_value: %v, pX_address: %v\n", x, y, pX, pointerToPx)
	*pX = 1000 // change the value of the thing pointed to by pX to 1000
	fmt.Printf("x: %v, y: %v, pX_value: %v, pX_address: %v\n", x, y, pX, pointerToPx)
	pX = &y // change the value pX to the address of y
	fmt.Printf("x: %v, y: %v, pX_value: %v, pX_address: %v\n", x, y, pX, pointerToPx)
	*pX = -50 // change the value of the thing pointed to by pX to -50
	fmt.Printf("x: %v, y: %v, pX_value: %v, pX_address: %v\n", x, y, pX, pointerToPx)

	i := 9
	zero(&i) // a normal function has to take the address variable
	fmt.Println(i)

	ii := myInt(3) // var ii myInt = 3
	ii.zeroMyInt() // a receiver function can take the value variable directly and infer the address
	fmt.Println(ii)
}
