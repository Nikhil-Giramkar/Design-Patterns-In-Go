package main

import "fmt"

func main() {
	book := newBook("The Jungle Book")

	reader1 := &Reader{
		emailId: "reader1@yahoo.in",
	}

	reader2 := &Reader{
		emailId: "reader2@gmail.com",
	}

	book.register(reader1)
	book.register(reader2)

	book.updateAvailability()

	fmt.Println("\nReader2 unsubscibed for email notifications")
	book.deregister(reader2)

	book.updateAvailability()
}
