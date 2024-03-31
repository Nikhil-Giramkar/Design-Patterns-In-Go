package main

type PaymentGateway interface {
	processPayment()
}

func getPaymentOption(option string) PaymentGateway {
	if option == "Amazon" {
		return &AmazonAdapter{
			amazon: &Amazon{},
		}
	} else if option == "PayPal" {
		return &PayPalAdapter{
			paypal: &PayPal{},
		}
	}

	return nil
}
