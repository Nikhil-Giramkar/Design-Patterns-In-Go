package main

type PayPalAdapter struct {
	paypal *PayPal
}

func (p *PayPalAdapter) processPayment() {
	p.paypal.makePayment()
}
