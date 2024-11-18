package main

import (
	"fmt"
)

func findFactors(num int) []int {
	factors := []int{}

	factor := 2
	for num%factor == 0 {
		factors = append(factors, factor)
		num /= factor
	}

	println(num)
	factor = 3
	for factor*factor <= num {
		if num%factor == 0 {
			factors = append(factors, factor)
			num /= factor
		}
		factor += 2
	}

	if num > 1 {
		factors = append(factors, num)
	}

	return factors
}

func main() {
	var numToFactor int
	ok := true

	for ok {
		fmt.Printf("\nNumber to Factor: ")
		fmt.Scan(&numToFactor)

		if numToFactor <= 2 {
			ok = false
		} else {
			factors := findFactors(numToFactor)
			fmt.Printf("%v\n", factors)
		}
	}
}
