package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	pb "tutorial/examples/tutorialpb"

	"github.com/golang/protobuf/proto"
)

func main() {
	fname := "address_book.txt"
	// Marshal
	/*
		p := &pb.Person{
			Id:    1024,
			Name:  "Alex",
			Email: "alex@example.com",
			Phones: []*pb.Person_PhoneNumber{
				{Number: "555-4321", Type: pb.Person_HOME},
			},
		}

		book := &pb.AddressBook{}
		book.People = append(book.People, p)

		out, err := proto.Marshal(book)
		if err != nil {
			log.Fatalln("Failed to encode address bool:", err)
		}

		fname := "address_book.txt"
		if err := ioutil.WriteFile(fname, out, 0644); err != nil {
			log.Fatalln("Failed to write address book:", err)
		}
	*/

	// Unmarshal
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("%s: File not found")
		} else {
			log.Fatalln("Failed to read file:", err)
		}
	}
	book2 := &pb.AddressBook{}
	if err := proto.Unmarshal(in, book2); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}

	fmt.Printf("Address book2: %v", book2)
}
