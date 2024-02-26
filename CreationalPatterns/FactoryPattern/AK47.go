package main

type AK47 struct {
	Gun
}

func createAK47() IGun {
	return &AK47{
		Gun: Gun{
			name:  "AK-47",
			power: 87,
		},
	}
}
