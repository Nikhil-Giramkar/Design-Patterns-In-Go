package main

import "fmt"

type Doctor struct {
	next Department
}

func (d *Doctor) execute(p *Patient) {
	if p.doctorCheckupDone {
		fmt.Println("Doctor Checkup already done ")
		d.next.execute(p)
		return
	}

	fmt.Println("Doctor checking patient")
	p.doctorCheckupDone = true
	d.next.execute(p)
}

func (d *Doctor) setNext(nextDept Department) {
	d.next = nextDept
}