package main

import "fmt"

type Reception struct {
	next Department
}

func (r *Reception) execute(p *Patient) {
	if p.registerationDone {
		fmt.Println("Registeration already done ")
		r.next.execute(p)
		return
	}

	fmt.Println("Reception registering patient")
	p.registerationDone = true
	r.next.execute(p)
}

func (r *Reception) setNext(nextDept Department) {
	r.next = nextDept
}
