package gorsa

import "fmt"

// GCD returns the greatest common divisor of a and b.
// using the Euclidean algorithm
func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

func LCM(a, b int) int {
	//(a*b)/GCD(a,b) may cause overflow
	//this works GCD divids both evenly
	return a * (b / GCD(a, b))
}

func FastExp(num, pow int) int {
	res := 1
	for pow > 0 {
		if pow%2 == 1 {
			res *= num
		}
		num *= num
		pow /= 2
	}
	return res
}

func FastExpMod(num, pow, mod int) int {
	res := 1
	for pow > 0 {
		if pow%2 == 1 {
			res = (res * num) % mod
		}
		num = (num * num) % mod
		pow /= 2
	}
	return res
}

// Build a sieve of Eratosthenes.
func SieveOfEratosthenes(max int) []bool {
	sieve := make([]bool, max+1)
	for i := 2; i <= max; i++ {
		sieve[i] = true
	}
	for i := 2; i*i <= max; i++ {
		if sieve[i] {
			for j := i * i; j <= max; j += i {
				sieve[j] = false
			}
		}
	}
	return sieve
}

func PrintSieve(sieve []bool) {
	for i, _ := range sieve {
		if sieve[i] {
			fmt.Printf("%d ", i)
		}
	}
	println()
}

// Convert the sieve into a slice holding prime numbers.
func SieveToPrimes(sieve []bool) []int {
	primes := []int{}
	for i, v := range sieve {
		if v {
			primes = append(primes, i)
		}
	}
	return primes
}
