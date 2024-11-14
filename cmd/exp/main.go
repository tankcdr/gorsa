package main

import (
	"fmt"

	"github.com/tankcdr/gorsa"
)

func main() {
	for ok := true; ok; {
		println("Please enter numbers to calculate the FastExponentiation and FastExponentiationModulus.")
		// Get the target value.
		var num int
		fmt.Printf("Enter number: ")
		ascan, _ := fmt.Scanln(&num)

		if ascan == 0 {
			fmt.Print("No input provided. Exiting...\n")
			break
		}

		var power int
		fmt.Printf("Enter power: ")
		bscan, _ := fmt.Scanln(&power)

		if bscan == 0 {
			fmt.Print("No input provided. Exiting...\n")
			break
		}

		var mod int
		fmt.Printf("Enter modulos: ")
		cscan, _ := fmt.Scanln(&mod)

		if cscan == 0 {
			fmt.Print("No input provided. Exiting...\n")
			break
		}

		fmt.Printf("FastExp(%d,%d) = %d\n", num, power, gorsa.FastExp(num, power))
		fmt.Printf("FastExpMod(%d,%d,%d) = %d\n", num, power, mod, gorsa.FastExpMod(num, power, mod))
		println("------------------------------------------")
	}
}
