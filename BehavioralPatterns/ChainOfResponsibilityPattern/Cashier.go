package main

import "fmt"

type Cashier struct {
	next Department
}

func (d *Cashier) execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Payment already done ")
		return
	}

	fmt.Println("Cashier getting money from patient")
	p.paymentDone = true
}

func (d *Cashier) setNext(nextDept Department) {
	d.next = nextDept
}
