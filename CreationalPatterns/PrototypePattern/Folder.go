package main

import "fmt"

type Folder struct {
	name     string
	children []INode
}

func (f *Folder) print(indent string) {
	fmt.Println(indent + f.name)
	for _, item := range f.children {
		item.print(indent + indent)
	}
}

func (f *Folder) clone() INode {
	cloneFolder := &Folder{name: f.name + "_clone"}

	var clonedItems []INode
	for _, item := range f.children {
		clonedItem := item.clone()
		clonedItems = append(clonedItems, clonedItem)
	}

	cloneFolder.children = clonedItems
	return cloneFolder
}
