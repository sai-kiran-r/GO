// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	// v := 10
// 	// fmt.Print("This is working ", v)

// 	//lets get the name
// 	readName := bufio.NewReader(os.Stdin)

// 	fmt.Print("Type your name: ")
// 	input, _ := readName.ReadString('\n')
// 	fmt.Println("Hi", input, "enter a value")
// 	// fmt.Println("Hi %V enter a value:", input)

// }
// package main

// import "fmt"

// func main() {
// 	var ary = []int16{12, 7, 4, 67, 82}
// 	ary = append(ary[2:len(ary)])
// 	fmt.Println(ary[2])
// }

// package main

// import "fmt"

// func main() {
// 	var out int
// 	for j := 0; j < 20; j++ {
// 		out = j*j + out
// 		if out > 12 {
// 			goto theEnd
// 		}
// 	}
// theEnd:
// 	fmt.Println(out)
// }
// package main

// import "fmt"

// func main() {
// 	var output string
// 	j := 12.47
// 	if j < 12.0 {
// 		output = "less"
// 	} else if j > 13 {
// 		output = "more"
// 	} else {
// 		output = "near"
// 	}
// 	fmt.Println(output)
// }

package main

import (
	"fmt"
)

func main() {
	// fmt.Print("hi\n")
	trail()

}

func trail() {
	fmt.Print("This is second function\n")
}
