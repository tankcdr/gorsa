package gorsa_test

import (
	"reflect"
	"testing"

	"github.com/tankcdr/gorsa"
)

func equal(a, b []bool) bool {
	return reflect.DeepEqual(a, b)
}

func TestMath_gcd(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"gcd(0, 0)", 0, 0, 0},
		{"gcd(0, 1)", 0, 1, 1},
		{"gcd(1, 0)", 1, 0, 1},
		{"gcd(1, 1)", 1, 1, 1},
		{"gcd(2, 1)", 2, 1, 1},
		{"gcd(1, 2)", 1, 2, 1},
		{"gcd(2, 2)", 2, 2, 2},
		{"gcd(270, 192)", 270, 192, 6},
		{"gcd(1071, 462)", 1071, 462, 21},
		{"gcd(462, 1071)", 462, 1071, 21},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gorsa.GCD(tt.a, tt.b); got != tt.want {
				t.Errorf("gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMath_lcm(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"lcm(270, 192)", 270, 192, 8640},
		{"lcm(7469, 2464)", 7469, 2464, 239008},
		{"lcm(55290, 115430)", 55290, 115430, 6579510},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gorsa.LCM(tt.a, tt.b); got != tt.want {
				t.Errorf("lcm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMath_FastExp(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"FastExp(8, 6)", 8, 6, 262144},
		{"FastExp(7, 10)", 7, 10, 282475249},
		{"FastExp(9, 13)", 9, 13, 2541865828329},
		{"FastExp(213, 5)", 213, 5, 438427732293},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gorsa.FastExp(tt.a, tt.b); got != tt.want {
				t.Errorf("gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMath_FastExpMod(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		a    int
		b    int
		c    int
		want int
	}{
		{"FastExpMod(8, 6, 10)", 8, 6, 10, 4},
		{"FastExpMod(7, 10, 101)", 7, 10, 101, 65},
		{"FastExpMod(9, 13, 283)", 9, 13, 283, 179},
		{"FastExpMod(213, 5, 1000)", 213, 5, 1000, 293},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gorsa.FastExpMod(tt.a, tt.b, tt.c); got != tt.want {
				t.Errorf("gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMath_SieveOfEratosthenes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		max  int
		want []bool
	}{
		{"SieveOfEratosthenes(10)", 10, []bool{false, false, true, true, false, true, false, true, false, false, false}},
		{"SieveOfEratosthenes(20)", 20, []bool{false, false, true, true, false, true, false, true, false, false, false, true, false, true, false, false, false, true, false, true, false}},
		{"SieveOfEratosthenes(30)", 30, []bool{false, false, true, true, false, true, false, true, false, false, false, true, false, true, false, false, false, true, false, true, false, false, false, true, false, false, false, false, false, true, false}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gorsa.SieveOfEratosthenes(tt.max); !equal(got, tt.want) {
				t.Errorf("SieveOfEratosthenes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMath_SieveToPrimes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		sieve []bool
		want  []int
	}{
		{"PrintSieve(10)", []bool{false, false, true, true, false, true, false, true, false, false, false}, []int{2, 3, 5, 7}},
		{"PrintSieve(20)", []bool{false, false, true, true, false, true, false, true, false, false, false, true, false, true, false, false, false, true, false, true, false}, []int{2, 3, 5, 7, 11, 13, 17, 19}},
		{"PrintSieve(30)", []bool{false, false, true, true, false, true, false, true, false, false, false, true, false, true, false, false, false, true, false, true, false, false, false, true, false, false, false, false, false, true, false}, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gorsa.SieveToPrimes(tt.sieve)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrintSieve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMath_SieveOfEuler(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		max  int
		want []bool
	}{
		{"SieveOfEuler(10)", 10, []bool{false, false, true, true, false, true, false, true, false, false, false}},
		{"SieveOfEuler(20)", 20, []bool{false, false, true, true, false, true, false, true, false, false, false, true, false, true, false, false, false, true, false, true, false}},
		{"SieveOfEuler(30)", 30, []bool{false, false, true, true, false, true, false, true, false, false, false, true, false, true, false, false, false, true, false, true, false, false, false, true, false, false, false, false, false, true, false}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gorsa.SieveOfEuler(tt.max); !equal(got, tt.want) {
				t.Errorf("SieveOfEuler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMath_IsProbablyPrime(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		n    int
		want bool
	}{
		{"IsProbablyPrime(-1)", -1, false},
		{"IsProbablyPrime(0)", 0, false},
		{"IsProbablyPrime(1)", 1, false},
		{"IsProbablyPrime(2)", 2, true},
		{"IsProbablyPrime(3)", 3, true},
		{"IsProbablyPrime(5)", 5, true},
		{"IsProbablyPrime(7)", 7, true},
		{"IsProbablyPrime(11)", 11, true},
		{"IsProbablyPrime(13)", 13, true},
		{"IsProbablyPrime(4)", 4, false},
		{"IsProbablyPrime(6)", 6, false},
		{"IsProbablyPrime(8)", 8, false},
		{"IsProbablyPrime(9)", 9, false},
		{"IsProbablyPrime(10)", 10, false},
		{"IsProbablyPrime(12)", 12, false},
		{"IsProbablyPrime(47)", 47, true},
		{"IsProbablyPrime(101)", 101, true},
		{"IsProbablyPrime(7919)", 7919, true},
		{"IsProbablyPrime(104729)", 104729, true},
		{"IsProbablyPrime(999331)", 999331, true},
		{"IsProbablyPrime(10000)", 10000, false},
		{"IsProbablyPrime(123456)", 123456, false},
		{"IsProbablyPrime(999998)", 999998, false},
		//Carmichael numbers
		{"IsProbablyPrime(561)", 561, false},
		{"IsProbablyPrime(1105)", 1105, false},
		//large Merseene prime
		{"IsProbablyPrime(2_147_483_647)", 2_147_483_647, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gorsa.IsProbablyPrime(tt.n, 20); got != tt.want {
				t.Errorf("IsProbablyPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMath_TotientCarmichael(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		p    int
		q    int
		want int
	}{
		{"TotientCarmichael(2,5)", 2, 5, 4},
		{"TotientCarmichael(3,5)", 3, 5, 4},
		{"TotientCarmichael(5,7)", 5, 7, 12},
		{"TotientCarmichael(11,13)", 11, 13, 60},
		{"TotientCarmichael(13,17)", 13, 17, 48},
		{"TotientCarmichael(17,23)", 17, 23, 176},
		{"TotientCarmichael(19,23)", 19, 23, 198},
		{"TotientCarmichael(29,31)", 29, 31, 420},
		{"TotientCarmichael(61,53)", 61, 53, 780},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gorsa.TotientCarmichael(tt.p, tt.q); got != tt.want {
				t.Errorf("TotientCarmichael() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMath_InverseModulo(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		a    int
		n    int
		want int
	}{
		{"InverseModulo(2, 5)", 2, 5, 3},            // 2 * 3 mod 5 = 1
		{"InverseModulo(3, 11)", 3, 11, 4},          // 3 * 4 mod 11 = 1
		{"InverseModulo(10, 17)", 10, 17, 12},       // 10 * 12 mod 17 = 1
		{"InverseModulo(7, 26)", 7, 26, 15},         // 7 * 15 mod 26 = 1
		{"InverseModulo(9, 28)", 9, 28, 25},         // 9 * 25 mod 28 = 1
		{"InverseModulo(17, 3120)", 17, 3120, 2753}, // 17 * 2753 mod 3120 = 1
		{"InverseModulo(3, 7)", 3, 7, 5},            // 3 * 5 mod 7 = 1
		{"InverseModulo(2, 9)", 2, 9, 5},            // 2 * 5 mod 9 = 1
		{"InverseModulo(4, 9)", 4, 9, 7},            // 4 * 7 mod 9 = 1
		{"InverseModulo(6, 13)", 6, 13, 11},         // 6 * 11 mod 13 = 1
		{"InverseModulo(5, 12)", 5, 12, 5},          // 5 * 5 mod 12 = 1
		{"InverseModulo(8, 15)", 8, 15, 2},          // 8 * 2 mod 15 = 1
		{"InverseModulo(11, 21)", 11, 21, 2},        // 11 * 2 mod 21 = 1
		{"InverseModulo(14, 33)", 14, 33, 26},       // 14 * 26 mod 33 = 1
		{"InverseModulo(19, 81)", 19, 81, 64},       // 19 * 64 mod 81 = 1
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gorsa.InverseModulo(tt.a, tt.n); got != tt.want {
				t.Errorf("InverseModulo() = %v, want %v", got, tt.want)
			}
		})
	}
}
