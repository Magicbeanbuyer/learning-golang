package main

import "fmt"

type bot interface {
	greet(name string) string
	smile()
}

// type enBot has to implement every function of interface bot in order to qualify
type enBot struct {
}

func (eb enBot) greet(name string) string {
	return "Hi, " + name
}

func (enBot) smile() {
	print("π\n")
}

type chBot struct {
}

func (cb chBot) greet(name string) string {
	return "ε¨οΌ" + name
}

// chBot has function not defined in interface bot, but chBot still qualifies as an implementation of interface bot
func (cb chBot) bye() {
	print("ζζ\n")
}

func (cb chBot) smile() {
	print("βΊοΈ\n")
}

func helloWorld(b bot) {
	name := "World"
	fmt.Println(b.greet(name))
}

func wrapper() {
	enBot := enBot{}
	chBot := chBot{}
	helloWorld(enBot)
	helloWorld(chBot)
	enBot.smile()
	chBot.smile()
	chBot.bye()
}
