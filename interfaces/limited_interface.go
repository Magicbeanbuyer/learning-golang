package main

// The static type of the interface determines what methods may be invoked with an interface variable, even though the 
// concrete value inside may have a larger set of methods.
// interface variable always has the form (value, concrete type) 
func li_wrapper() {
	var p person = kid{}
	p.greet("lucy")
	// ./limited_interface.go:23:4: p.smile undefined (type person has no field or method smile)
	// interface person only provide access to method greet, not smile
	// p.smile()
	
	// type assertion
	// asserts that the interface value p holds the concrete type comedian
	// and assigns the underlying comedian value to the variable jimmy.
	jimmy := p.(comedian)
	jimmy.smile()
}

type person interface {
	greet(name string) string
}

type comedian interface {
	smile()
}

// type enBot has to implement every function of interface bot in order to qualify
type kid struct {
}

func (eb kid) greet(name string) string {
	var s string = "Hi, " + name
	print(s + "\n")
	return s
}

func (kid) smile() {
	print("😃\n")
}
