package main

import (
	"fmt"
	"time"
)

func main() {

	n := time.Now()
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter text here: ")
	// input, _ := reader.ReadString('\n')
	// fmt.Println("You entered: ", input)

	// fmt.Print("Enter a number: ")
	// numInput, _ := reader.ReadString('\n')
	// aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
	// if err != nil {
	// 	fmt.Println(err)
	// 	fmt.Println("entered a value", numInput)
	// } else {
	// 	fmt.Println("value: ", aFloat)
	// }
	fmt.Print(n.Format(time.ANSIC))
	t := time.Now()
	fmt.Print(t.Format("m/d/yyyy"))
}
