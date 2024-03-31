package main

import (
	"fmt"
)

type Amazon struct{}

func (a *Amazon) payAmazon() {

	fmt.Printf("Payment processed via Amazon Pay \n")
}
