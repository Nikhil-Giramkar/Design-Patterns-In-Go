package main

import "fmt"

type File struct {
	name string
}

func (f *File) print(indent string) {
	fmt.Println(indent + f.name)
}

func (f *File) clone() INode {
	return &File{
		name: f.name + "_clone",
	}
}
