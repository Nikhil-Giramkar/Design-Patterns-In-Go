package main

import "fmt"

func main() {
	nikeFactory, _ := SetupSportsFactory("Nike")

	nikeShoes := nikeFactory.makeShoes()
	nikeShirts := nikeFactory.makeShirts()

	adidasFactory, _ := SetupSportsFactory("Adidas")
	adidasShirts := adidasFactory.makeShirts()
	adidasShoes := adidasFactory.makeShoes()

	printShirtDetails(nikeShirts)
	printShoeDetails(nikeShoes)

	fmt.Println()

	printShirtDetails(adidasShirts)
	printShoeDetails(adidasShoes)
}

func SetupSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "Nike" {
		return &Nike{}, nil
	}

	if brand == "Adidas" {
		return &Adidas{}, nil
	}

	return nil, fmt.Errorf("factory is not created")
}

func printShoeDetails(s IShoeFactory) {
	fmt.Println(s.getLace())

	fmt.Println(s.getSole())
	fmt.Println()
}

func printShirtDetails(s IShirtFactory) {
	fmt.Println(s.getLogo())

	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}
