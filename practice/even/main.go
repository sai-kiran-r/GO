package main

import (
	"fmt"
	// "time"
)

func main(){
	oddEven()
}

func oddEven(){
	fmt.Println("This is second function")
	// value := 1000
	count := 0
	 
	for value:= 1000; value <= 9999; value++{
		for num:= 10; num <= 99; num++{
			product := num*num
			if product % 2 == 0{
				if product >= 1000 && product <= 9999{
					count++
				}
				
				//fmt.Println("The number is even ended", num)
			} 

		}
	}
	fmt.Println("The count is", count)
}