package main

type M16 struct {
	Gun
}

func createM16() IGun {
	return &M16{
		Gun: Gun{
			name:  "M-16",
			power: 98,
		},
	}
}
