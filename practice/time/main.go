package main

import (
	"fmt"
	"time"
)

func main() {
	// This will store the complete date and time
	displayTime := time.Now()
	// This displays the complete date and time
	fmt.Println("This will print the pressent date and current time", displayTime)

	//This will print only the current date
	fmt.Println("Current/Today's date", displayTime.Format("01-02-2006"))

	//This will print the current time
	fmt.Println("The time is:- ", displayTime.Format("15:04:05"))
}
