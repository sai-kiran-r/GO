package main

import(
	"fmt"
)

func main(){
	fmt.Println("This is sample")
	numbers := []int{10, 11, 22}
	fmt.Println("The numbers are",numbers)
	fmt.Println("length is", len(numbers))
	fmt.Println("The second value in slice is ", numbers[1])
}