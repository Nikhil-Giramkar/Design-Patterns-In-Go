package main

import "fmt"

type notification struct {
}

func (n *notification) notifyCredit(accountID string, amount int) {
	fmt.Printf("Sending %d credit notification to %s\n", amount, accountID)
}

func (n *notification) notifyDebit(accountID string, amount int) {
	fmt.Printf("Sending %d debit notification to %s\n", amount, accountID)
}
