package gorsa

import (
	"fmt"
	"math/rand"
)

// Encrypt returns the ciphertext of a message using the public key (e, n).
func Encrypt(message, e, n int) int {
	return FastExpMod(message, e, n)
}

// Decrypt returns the plaintext of a ciphertext using the private key (d, n).
func Decrypt(ciphertext, d, n int) int {
	return FastExpMod(ciphertext, d, n)
}

// Totient returns the number of positive integers less than n that are coprime to n.
// This function uses Carmichael's totient function.
func TotientCarmichael(p, q int) int {
	return LCM(p-1, q-1)
}

// RandomExponent returns a random exponent e such that 1 < e < totient and GCD(e, totient) = 1.
func RandomExponent(totient int) int {
	for {
		e := randRange(2, totient)
		if GCD(e, totient) == 1 {
			return e
		}
	}
}

// Extended Euclidean Algorithm
// Given a and n, this function returns the multiplicative inverse of a modulo n.
func InverseModulo(a, n int) int {
	t := 0
	newt := 1
	r := n
	newr := a

	for newr != 0 {
		quotient := r / newr
		t, newt = newt, t-quotient*newt
		r, newr = newr, r-quotient*newr
	}

	if r > 1 {
		return -1
	}
	if t < 0 {
		t += n
	}
	return t
}

func randRange(min, max int) int {
	return rand.Intn(max-min) + min
}

// IsProbablyPrime returns true if n is probably prime, and false if it is definitely composite.
// The function uses Fermat's little theorem to test for primality.
func IsProbablyPrime(n, numTests int) bool {
	//no need to check for 0 or 1 or negatives
	if n < 2 {
		return false
	}
	//2 and 3 are prime
	if n == 2 || n == 3 {
		return true
	}
	//even numbers are not prime
	if n%2 == 0 {
		return false
	}

	//Fermat's little theorem
	//a^(n-1) ≡ 1 (mod n)
	//numTests increases the probability of correctness
	for i := 0; i < numTests; i++ {
		randomNum := randRange(2, n)
		result := FastExpMod(randomNum, n-1, n)

		if result != 1 {
			return false
		}
	}
	return true
}

// FindPrime returns a prime number between min and max.
// The function uses IsProbablyPrime to test for primality.
func FindPrime(min, max, numTests int) int {
	for {
		n := randRange(min, max)
		if IsProbablyPrime(n, numTests) {
			return n
		}
	}
}

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

// FastExp returns num raised to the power of pow.
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

// FastExpMod returns num raised to the power of pow modulo mod.
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

// Build Euler's sieve.
func SieveOfEuler(max int) []bool {
	sieve := make([]bool, max+1)

	sieve[2] = true

	for p := 3; p < len(sieve); p += 2 {
		sieve[p] = true
	}

	for p := 3; p <= max; p += 2 {
		if sieve[p] {
			maxQ := max / p
			if maxQ%2 == 0 {
				maxQ--
			}

			for q := maxQ; q >= p; q -= 2 {
				if sieve[q] == true {
					sieve[p*q] = false
				}
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

// Convert a sieve into a slice holding prime numbers.
func SieveToPrimes(sieve []bool) []int {
	primes := []int{}
	for i, v := range sieve {
		if v {
			primes = append(primes, i)
		}
	}
	return primes
}
