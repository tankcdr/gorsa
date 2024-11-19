package main

import (
	"fmt"

	"github.com/tankcdr/gorsa"
)

const PRIMALITY_TESTS = 20

func main() {

	// generate two random prime numbers
	p := gorsa.FindPrime(10000, 50000, PRIMALITY_TESTS)
	q := gorsa.FindPrime(10000, 50000, PRIMALITY_TESTS)

	// public key modulus
	n := p * q

	// Carmichael's totient function
	totient := gorsa.TotientCarmichael(p, q)

	// public exponent
	e := gorsa.RandomExponent(totient)

	// private exponent
	d := gorsa.InverseModulo(e, totient)

	println("*** Public key ***")
	println("Public key modulus: ", n)
	println("Public exponent: ", e)
	println("*** Private  ***")
	println("Primes: ", p, q)
	println("Totient: ", totient)
	println("Private exponent: ", d)
	println()

	for {
		var message int
		fmt.Printf("Enter a message to encrypt (0 to quit): ")
		fmt.Scanln(&message)

		if message == 0 {
			break
		}

		ciphertext := gorsa.Encrypt(message, e, n)
		decrypted := gorsa.Decrypt(ciphertext, d, n)

		println("Ciphertext: ", ciphertext)
		println("Decrypted: ", decrypted)
		println()
	}
}
