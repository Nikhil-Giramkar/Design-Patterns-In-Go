package main

import "fmt"

type IShoeFactory interface {
	getLace() string
	getSole() string
}

type Shoe struct {
	lace string
	sole string
}

func (s *Shoe) getLace() string {
	return fmt.Sprintf("Lace Color: %s", s.lace)
}

func (s *Shoe) getSole() string {
	return fmt.Sprintf("Sole Color: %s", s.sole)
}
