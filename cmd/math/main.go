package main

import (
	"fmt"

	"github.com/tankcdr/gorsa"
)

func main() {

	for ok := true; ok; {
		// Get the target value.
		var a int
		fmt.Printf("Enter A: ")
		ascan, _ := fmt.Scanln(&a)

		if ascan == 0 {
			fmt.Print("No input provided. Exiting...\n")
			break
		}

		var b int
		fmt.Printf("Enter B: ")
		bscan, _ := fmt.Scanln(&b)

		if bscan == 0 {
			fmt.Print("No input provided. Exiting...\n")
			break
		}

		fmt.Printf("GCD(%d,%d) = %d\n", a, b, gorsa.GCD(a, b))
		fmt.Printf("LCM(%d,%d) = %d\n", a, b, gorsa.LCM(a, b))
	}
}
