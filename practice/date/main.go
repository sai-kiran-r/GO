package main

import (
	"fmt"
	"time"
)

func main() {
	// This will store the complete date and time
	displayDate := time.Now()

	// This displays the complete date and time
	fmt.Println("This will print the pressent date and current time", displayDate)

	//This displays current date - "Jan 2, 2006" is just a sample
	fmt.Println("This prints only the date", displayDate.Format("Jan 2, 2006"))

	//Below line prints the time stamp - month day time(HH:MM:SS)
	fmt.Println("Current day and time", displayDate.Format(time.Stamp))
}
