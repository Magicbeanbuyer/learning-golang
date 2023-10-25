package main

import "fmt"

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

	call_by_value_pointer_copies()
	change_pointed_value_by_dereferencing()
}

func zero(pI *int) {
	*pI = 0
}

type myInt int

func change_pointed_value_by_dereferencing() {
	ii := myInt(3) // var ii myInt = 3
	fmt.Println("ii value: ", ii)
	ii.zeroMyInt() // a receiver function can take the value variable directly and infer the address
	fmt.Println("ii value: ", ii)
}

func (m *myInt) zeroMyInt() {
	fmt.Println("Dereference the pointer and change the value of the pointed variable.")
	*m = 0
}

func (m *myInt) print_pointer_receiver_address_and_value() {
	// Learning Go p.113
	fmt.Printf("address of m: %v value of m: %v\n", &m, m)
	fmt.Println("The address of the pointer from within the method differs from the address of the outer scope " +
		"pointers. Because a copy of the pointer is passed to the method.")
	if m == nil {
		/*
			m := &myInt(5) doesn't work
			cannot take address of myInt(5) (constant 5 of type myInt)
			see Learning Go p.110
		*/
		x := myInt(5)
		m = &x
	}
	fmt.Printf("address of m: %v value of m: %v\n", &m, m)
}


func call_by_value_pointer_copies() {
	var num_nil_ptr *myInt
	fmt.Printf("address of num_nil_ptr: %v value of num_nil_ptr: %v\n", &num_nil_ptr, num_nil_ptr)
	num_nil_ptr.print_pointer_receiver_address_and_value()
	fmt.Println()
}
