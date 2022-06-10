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
	fmt.Println("Choose one among the following \n 1.Addition \n 2.Subraction \n 3.Multiplication \n 4.Division")
	fmt.Scan(&option)
	fmt.Println("Type two numbers")
	fmt.Scan(&numberOne, &numberTwo)

	if option == 1 {
		addition()
	} else if option == 2 {
		subraction()
	} else if option == 3 {
		multiplication()
	} else if option == 4 {
		division()
	} else {
		fmt.Println("Please choose right option")
	}

}

func addition() {
	total := numberOne + numberTwo
	fmt.Println("Added value is:", total)
}

func subraction() {
	println("Subraction")
}

func multiplication() {
	println("Multiplication")
}

func division() {
	println("Division")
}
