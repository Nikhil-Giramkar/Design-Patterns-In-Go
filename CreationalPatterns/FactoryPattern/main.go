package main

import "fmt"

func main() {

	ak47, _ := getGun("AK-47")
	m16, _ := getGun("M-16")

	printDetails(ak47)
	fmt.Println("-------------------")
	printDetails(m16)

}

func printDetails(gun IGun) {
	fmt.Println("Name: ", gun.getName())
	fmt.Println("Power: ", gun.getPower())
}
