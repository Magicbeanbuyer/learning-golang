package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"log"
	"os"
	pb "protobufDemo/build/github.com/magicbeanbuyer/protobuf/tutorialpb"
)

func main() {
	//p := pb.Person{
	//	Id:    1234,
	//	Name:  "John Doe",
	//	Email: "jdoe@example.com",
	//	Phones: []*pb.Person_PhoneNumber{
	//		{Number: "555-4321", Type: pb.Person_HOME},
	//	},
	//}
	//
	//addressBook := pb.AddressBook{People: []*pb.Person{&p}}

	//out, err := proto.Marshal(&addressBook)
	//if err != nil {
	//	log.Fatalln("Failed to encode address book:", err)
	//}
	//if err := os.WriteFile("addressbook", out, 0644); err != nil {
	//	log.Fatalln("Failed to write address book:", err)
	//}

	in, err := os.ReadFile("addressbook")
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	book := &pb.AddressBook{}
	if err := proto.Unmarshal(in, book); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}
	fmt.Println(book) // people:{name:"John Doe" id:1234 email:"jdoe@example.com" phones:{number:"555-4321" type:HOME}}
}
