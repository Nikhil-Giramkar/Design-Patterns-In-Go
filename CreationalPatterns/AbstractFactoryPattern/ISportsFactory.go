package main

type ISportsFactory interface {
	makeShoes() IShoeFactory
	makeShirts() IShirtFactory
}


