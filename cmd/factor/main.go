package main

import (
	"fmt"
	"time"

	"github.com/tankcdr/gorsa"
)

func findFactors(num int) []int {
	factors := []int{}

	factor := 2
	for num%factor == 0 {
		factors = append(factors, factor)
		num /= factor
	}

	for factor = 3; factor*factor <= num; factor += 2 {
		if num%factor == 0 {
			factors = append(factors, factor)
			num /= factor
		}
	}

	if num > 1 {
		factors = append(factors, num)
	}

	return factors
}

func findFactorsSieve(num int) []int {
	factors := []int{}

	factor := 2
	for num%factor == 0 {
		factors = append(factors, factor)
		num /= factor
	}

	for index := 1; primes[index]*primes[index] <= num; index++ {
		if num%primes[index] == 0 {
			factors = append(factors, primes[index])
			num /= primes[index]
		}
	}

	if num > 1 {
		factors = append(factors, num)
	}

	return factors
}

func multiplySlice(arr []int) int {
	multiple := 1
	for _, val := range arr {
		multiple *= val
	}
	return multiple
}

var primes []int

func main() {

	primes = gorsa.SieveToPrimes(gorsa.SieveOfEuler(1700000000))

	var numToFactor int
	ok := true

	for ok {
		fmt.Printf("\nNumber to Factor: ")
		fmt.Scan(&numToFactor)

		if numToFactor <= 2 {
			ok = false
		} else {
			// Find the factors the slow way.
			start := time.Now()
			factors := findFactors(numToFactor)
			elapsed := time.Since(start)
			fmt.Printf("findFactors:       %f seconds\n", elapsed.Seconds())
			fmt.Println(multiplySlice(factors))
			fmt.Println(factors)
			fmt.Println()

			// Use the Euler's sieve to find the factors.
			start = time.Now()
			factors = findFactorsSieve(numToFactor)
			elapsed = time.Since(start)
			fmt.Printf("findFactorsSieve: %f seconds\n", elapsed.Seconds())
			fmt.Println(multiplySlice(factors))
			fmt.Println(factors)
			fmt.Println()
		}
	}
}
