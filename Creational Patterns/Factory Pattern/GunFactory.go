package main

import "fmt"

func getGun(GunType string) (IGun, error) {

	if GunType == "AK-47" {
		return createAK47(), nil
	}

	if GunType == "M-16" {
		return createM16(), nil
	}

	return nil, fmt.Errorf("Wrong gun type passed")
}
