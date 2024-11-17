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
	fmt.Printf("Sieve of Eratosthenes Elapsed: %f seconds\n", elapsed.Seconds())

	startEuler := time.Now()
	sieveEuler := gorsa.SieveOfEuler(max)
	elapsedEuler := time.Since(startEuler)
	fmt.Printf("Sieve of Euler Elapsed: %f seconds\n", elapsedEuler.Seconds())

	if max <= 1000 {
		gorsa.PrintSieve(sieve)
		gorsa.PrintSieve(sieveEuler)

		primes := gorsa.SieveToPrimes(sieve)
		fmt.Println(primes)

		primes = gorsa.SieveToPrimes(sieveEuler)
		fmt.Println(primes)
	}
}
