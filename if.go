package main

import "fmt"

func main() {
	age := 20

	//Check if person is old enough to drink
	if age > 21 {
		fmt.Println("You may enter the bar because you can drink")
	}

	if age == 21 {
		fmt.Println("BlackJack! You may enter the bar")
	}

	if age >= 18 && age < 21 {
		fmt.Println("Cant drink here, but you can go to Canada")
	}

	if age < 21 {
		fmt.Println("You shall not pass!")
	}
}
