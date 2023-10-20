package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 0b101     // binary integer 5
	b := 1_000_000 // _ for readability
	c := 9.1e+3
	d := 3.1e-2
	e := '\u90d1' //rune
	f := "hi \n hey \\n"
	g := `"la" \n da`
	var x Decimal = 1
	fmt.Println(a, b, c, d, e, f, g, x)
	fmt.Println(reflect.TypeOf(x))
}
