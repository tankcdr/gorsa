package main

import (
	"fmt"
	"time"

	"github.com/tankcdr/gorsa"
)

func main() {
	var max int
	fmt.Printf("Max: ")
	fmt.Scan(&max)

	start := time.Now()
	sieve := gorsa.SieveOfEratosthenes(max)
	elapsed := time.Since(start)
	fmt.Printf("Elapsed: %f seconds\n", elapsed.Seconds())

	if max <= 1000 {
		gorsa.PrintSieve(sieve)

		primes := gorsa.SieveToPrimes(sieve)
		fmt.Println(primes)
	}
}
