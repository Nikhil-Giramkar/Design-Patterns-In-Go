package main

type AmazonAdapter struct {
	amazon *Amazon
}

func (a *AmazonAdapter) processPayment() {
	a.amazon.payAmazon()
}
