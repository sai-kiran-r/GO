package main

import (
	"fmt"
)

func main() {
	var option int
	var numberOne int
	var numberTwo int
	c := "Welcome to Calculator"

	fmt.Println(c)
	
	fmt.Println("Enter two numbers")
	fmt.Scan(&numberOne, &numberTwo)

	fmt.Println("Choose one among the following \n 1.Addition \n 2.Subraction \n 3.Multiplication \n 4.Division")
	fmt.Scan(&option)

	if option == 1 {
		fmt.Println("Sum of two numbers is: ", numberOne + numberTwo) 
	} else if option == 2 {
		fmt.Println("Difference of two numbers is: ", numberOne - numberTwo)
	} else if option == 3 {
		fmt.Println("Multiplication of two numbers is: ", numberOne * numberTwo)
	} else if option == 4 {
		fmt.Println("Division of two numbers is: ", numberOne / numberTwo)
	} else {
		fmt.Println("Please choose right option")
	}

}



