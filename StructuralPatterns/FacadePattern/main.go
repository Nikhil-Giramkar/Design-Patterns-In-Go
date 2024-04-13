package main

import "fmt"

func main() {
	fmt.Println("-------------------------------------------")

	fmt.Println("New Wallet Initialized")
	myWallet := newWallet("ABC", 1234)

	fmt.Println("-------------------------------------------")

	fmt.Println("Adding 100 Rs. Balance")
	//err := myWallet.addMoneyToWallet("ABD", 1234, 100) 
	err := myWallet.addMoneyToWallet("ABC", 1234, 100)
	if err != nil{
		panic(err)
	}
	fmt.Println("-------------------------------------------")

	fmt.Println("Debit 20 Rs. Balance")
	//err= myWallet.debitMoneyFromWallet("ABC", 4567, 20)
	err= myWallet.debitMoneyFromWallet("ABC", 1234, 20)
	if err != nil{
		panic(err)
	}
	fmt.Println("-------------------------------------------")
}
