package main

import "fmt"

type securityCode struct {
	pin int
}

func newPin(newPin int) *securityCode {
	fmt.Println("Security Code created")
	return &securityCode{
		pin: newPin,
	}
}
func (s *securityCode) checkPin(enteredPin int) (ok bool, err error) {
	if s.pin != enteredPin {
		return false, fmt.Errorf("pin does not match")
	}
	fmt.Println("Security Code Verified")
	return true, nil
}
