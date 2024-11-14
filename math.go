package gorsa

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
