package main

import "fmt"

// Add returns a+b without using + or -: XOR is addition-without-carry, AND
// shifted left is the carry; loop until no carry remains. Works for
// negatives via two's complement on fixed-width ints.
func Add(a, b int) int {
	for b != 0 {
		carry := (a & b) << 1
		a = a ^ b
		b = carry
	}
	return a
}

func main() {
	fmt.Println(Add(1, 2), Add(-7, 3), Add(-2, -3))
}
