package gorsa_test

import (
	"testing"

	"github.com/tankcdr/gorsa"
)

func TestMath_gcd(t *testing.T) {
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
