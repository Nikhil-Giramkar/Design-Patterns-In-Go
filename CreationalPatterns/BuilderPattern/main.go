package main

import "fmt"

func main() {
	fmt.Println("New director assigned and a builder provided, who can build Normal House")
	normalBuilder := getBuilder("Normal")

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floors)
	fmt.Println("-----------------------------")

	fmt.Println("New builder provided, who can build Igloo House")
	iglooBuilder := getBuilder("Igloo")
	director.setNewBuilder(iglooBuilder)

	iglooHouse := director.buildHouse()

	fmt.Printf("\nIgloo House Door Type: %s\n", iglooHouse.doorType)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.windowType)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.floors)

}
