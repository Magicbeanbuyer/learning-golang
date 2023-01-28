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
	print("😃\n")
}

type chBot struct {
}

func (cb chBot) greet(name string) string {
	return "嗨，" + name
}

// chBot has function not defined in interface bot, but chBot still qualifies as an implementation of interface bot
func (cb chBot) bye() {
	print("拜拜\n")
}

func (cb chBot) smile() {
	print("☺️\n")
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
