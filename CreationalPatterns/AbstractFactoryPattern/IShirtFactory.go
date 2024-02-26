package main

import "fmt"

type IShirtFactory interface {
	getLogo() string
	getSize() int
}

type Shirt struct {
	logo string
	size int
}

func (s *Shirt) getLogo() string {
	return fmt.Sprintf("Logo: %s", s.logo)
}

func (s *Shirt) getSize() int {
	return s.size
}
