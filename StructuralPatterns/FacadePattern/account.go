package main

import "fmt"

type account struct {
	accountID string
	pin       *securityCode
	balance   int
}

func newAccount(accountID string, pin int) *account {
	fmt.Println("Account Created")
	return &account{
		accountID: accountID,
		pin:       newPin(pin),
		balance:   0,
	}
}

func (a *account) checkAccountID(enteredID string) error {
	if a.accountID != enteredID {
		return fmt.Errorf("AccountID mismatch")
	}
	fmt.Println("Account ID verified")
	return nil
}

func (a *account) creditAmount(amount int) {
	a.balance += amount
	fmt.Println("Amount added to balance successfully")
}

func (a *account) debitAmount(amount int) error {
	if amount > a.balance {
		return fmt.Errorf("insufficient Balance")
	}
	a.balance -= amount
	fmt.Println("Amount debited from balance successfully")
	return nil
}
