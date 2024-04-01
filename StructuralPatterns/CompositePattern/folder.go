package main

import "fmt"

type folder struct {
	name       string
	components []component
}

func (f *folder) search(key string) {
	fmt.Printf("Searching recursively  %s in %s \n", key, f.name)
	for _, item := range f.components {
		item.search(key)
	}
}

func (f *folder) add(item component) {
	f.components = append(f.components, item)
}
