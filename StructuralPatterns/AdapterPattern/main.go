package main

import (
	"fmt"
)

func main() {
	fmt.Println("Payment via PayPal")
	paymentGateway := getPaymentOption("PayPal")

	fmt.Println("Proceeding with payment")
	paymentGateway.processPayment()
	fmt.Println("Payment Done...")

	fmt.Println(".....................")

	paymentGateway = getPaymentOption("Amazon")
	fmt.Println("Proceeding with payment")
	paymentGateway.processPayment()
	fmt.Println("Payment Done...")

}
