package main

import "fmt"

type PayPal struct{}

func (p *PayPal) makePayment() {
	fmt.Printf("Payment processed via Paypal\n")
}
