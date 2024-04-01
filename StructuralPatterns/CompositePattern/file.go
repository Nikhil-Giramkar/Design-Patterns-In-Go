package main

import "fmt"

type file struct {
	name string
}

func (f *file) search(key string) {
	fmt.Printf("Searching %s in %s  \n", key, f.name)
}
