package main

func main() {
	patient := &Patient{
		name: "Nikhil",
	}

	reception := &Reception{}
	doctor := &Doctor{}
	medical := &Medical{}
	cashier := &Cashier{}

	reception.setNext(doctor)
	doctor.setNext(medical)
	medical.setNext(cashier)

	reception.execute(patient)
}
