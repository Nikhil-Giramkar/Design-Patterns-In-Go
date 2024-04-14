package main

import "fmt"

func main() {
	prod1 := &Product{
		name:  "Water Bottle",
		price: 350,
	}

	prod2 := &Product{
		name:  "Keyboard",
		price: 500,
	}

	productsCollection := &ProductsCollection{
		products: []*Product{prod1, prod2},
	}

	productsIterator := productsCollection.createIterator()

	for productsIterator.hasNext() {
		product := productsIterator.getNext()
		fmt.Printf("Product Is: %+v\n", product)
	}

}
