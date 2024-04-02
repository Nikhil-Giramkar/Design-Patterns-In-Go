package main

import "fmt"

func main() {

	//Lets create a pizza
	vegPizza := &VeggieMania{}

	//Add tomato toppping
	tomatoToppedPizza := &TomatoTopping{
		pizza: vegPizza,
	}

	//Add tomato + Cheese topping
	cheeseAndTomatoToppedPizza := &CheeseTopping{
		pizza: tomatoToppedPizza,
	}

	//Cheese topping only
	cheeseToppedPizza := &CheeseTopping{
		pizza: vegPizza,
	}

	fmt.Println("Veg Pizza Plain = ", vegPizza.getPrice())
	fmt.Println("Veg Pizza + Cheese Topping = ", cheeseToppedPizza.getPrice())
	fmt.Println("Veg Pizza + Tomato Topping = ", tomatoToppedPizza.getPrice())
	fmt.Println("Veg Pizza + Tomato + Cheese Topping = ", cheeseAndTomatoToppedPizza.getPrice())

}
