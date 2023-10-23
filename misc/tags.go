// package misc

package main

import (
	"fmt"
	"reflect"
)

func wrapper_tag() {
	type Me struct {
		Firstname string `mytag:"FirstName"`
		Lastname  string `mytag:"LastName"`
	}

	m := Me{"Anjana", "Shankar"}
	t := reflect.TypeOf(m)

	for _, fieldName := range []string{"Firstname", "Lastname"} {
		field, found := t.FieldByName(fieldName)
		if !found {
			continue
		}
		fmt.Printf("\nField: Me.%s\n", fieldName)
		fmt.Printf("\tWhole tag value : %q\n", field.Tag)
		fmt.Printf("\tValue of 'mytag': %q\n", field.Tag.Get("mytag"))
	}
}
