package main

import (
	"library/protocol"
	"log"

	"google.golang.org/protobuf/proto"
)

func main() {
	book := &protocol.Book{
		Name: "Danial",
		Isbn: 123456,
	}
	data, err := proto.Marshal(book)
	if err != nil {
		log.Fatal("Marshaling error:", err)
	}
	newBook := protocol.Book{}
	err = proto.Unmarshal(data, &newBook)
	if err != nil {
		log.Fatal("Unmarshaling error:", err)
	}
	println("Name:", newBook.Name)
	println("Isbn:", newBook.Isbn)
}
