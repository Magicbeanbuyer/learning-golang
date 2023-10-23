package main

import (
	"fmt"
	"reflect"
)

func wrapper_reflect() {
	var tom waiter = human{nickname: "doudou"}
	jimmy := tom.(comedian)

	fmt.Println("p")
	fmt.Println("type: ", reflect.TypeOf(tom)) // main.human
	fmt.Println("value: ", reflect.ValueOf(tom))

	jimmy_value := reflect.ValueOf(jimmy)
	fmt.Println("jimmy")
	fmt.Println("type: ", reflect.TypeOf(jimmy)) // main.human !!!
	fmt.Println("value: ", jimmy_value)
	// Kind of a reflection object describes the underlying type, not the static type. 
	fmt.Println("kind: ", jimmy_value.Kind()) // struct, because human is a struct
}
