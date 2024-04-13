package main

import "fmt"

const (
	CREDIT = "credit"
	DEBIT  = "debit"
)

type wallet struct {
	notifier *notification
	ledger   *ledger
	account  *account
}

func newWallet(accountID string, pin int) *wallet {
	fmt.Println("Wallet Created successfully")
	return &wallet{
		account:  newAccount(accountID, pin),
		ledger:   &ledger{},
		notifier: &notification{},
	}
}

func (w *wallet) addMoneyToWallet(accountID string, enteredPin int, amount int) error {
	fmt.Println("Beginning to add money")

	err := w.account.checkAccountID(accountID)
	if err != nil {
		return err
	}

	_, err = w.account.pin.checkPin(enteredPin)
	if err != nil {
		return err
	}
	w.account.creditAmount(amount)
	w.notifier.notifyCredit(accountID, amount)
	w.ledger.makeEntry(accountID, CREDIT, amount)
	return nil
}

func (w *wallet) debitMoneyFromWallet(accountID string, enteredPin int, amount int) error {
	fmt.Println("Beginning to debit money")

	err := w.account.checkAccountID(accountID)
	if err != nil {
		return err
	}

	_, err = w.account.pin.checkPin(enteredPin)
	if err != nil {
		return err
	}
	w.account.debitAmount(amount)
	w.notifier.notifyDebit(accountID, amount)
	w.ledger.makeEntry(accountID, DEBIT, amount)
	return nil
}
