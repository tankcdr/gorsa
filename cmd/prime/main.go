package main

import (
	"fmt"
	"math"

	"github.com/tankcdr/gorsa"
)

const NUM_TESTS = 20

func printType(prime bool) string {
	if prime {
		return "Prime"
	}
	return "Composite"
}

func testKnownValues() {
	primes := []int{
		10009, 11113, 11699, 12809, 14149,
		15643, 17107, 17881, 19301, 19793,
	}
	composites := []int{
		10323, 11397, 12212, 13503, 14599,
		16113, 17547, 17549, 18893, 19999,
	}

	fmt.Printf("Probabilty : %f\n", 1-math.Pow(0.5, float64(NUM_TESTS)))

	fmt.Println("Primes:")
	for _, p := range primes {
		fmt.Printf("%d: %s\n", p, printType(gorsa.IsProbablyPrime(p, NUM_TESTS)))
	}

	fmt.Println("Composites:")
	for _, c := range composites {
		fmt.Printf("%d: %s\n", c, printType(gorsa.IsProbablyPrime(c, NUM_TESTS)))
	}

}

func main() {
	testKnownValues()

	for true {
		println("Enter a range of numbers to randomly find a prime number.")
		var low int
		fmt.Print("Enter the low end of the range: ")
		fmt.Scanf("%d", &low)

		if low < 0 {
			break
		}

		var high int
		fmt.Print("Enter the high end of the range: ")
		fmt.Scanf("%d", &high)

		if high <= low {
			break
		}

		fmt.Printf("%d:%d %d\n", low, high, gorsa.FindPrime(low, high, NUM_TESTS))

	}
}
