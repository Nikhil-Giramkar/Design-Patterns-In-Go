package main

import "fmt"

type Reader struct {
	emailId string
}

func (r *Reader) update(bookName string) {
	fmt.Printf("Sending mail to %s for \"%s\"\n", r.emailId, bookName)
}

func (r *Reader) getId() string {
	return r.emailId
}
