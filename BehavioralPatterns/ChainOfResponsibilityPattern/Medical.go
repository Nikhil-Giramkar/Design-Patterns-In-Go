package main

import "fmt"

type Medical struct {
	next Department
}

func (d *Medical) execute(p *Patient) {
	if p.medicineCollectionDone {
		fmt.Println("Medicines already collected ")
		d.next.execute(p)
		return
	}

	fmt.Println("Medical giving medicine to patient")
	p.medicineCollectionDone = true
	d.next.execute(p)
}

func (d *Medical) setNext(nextDept Department) {
	d.next = nextDept
}