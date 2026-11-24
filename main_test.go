package main

import (
	"math/rand"
	"testing"
)

func TestAdd(t *testing.T) {
	cases := []struct{ a, b, want int }{
		{1, 2, 3},
		{2, 3, 5},
		{0, 0, 0},
		{-1, 1, 0},
		{-7, 3, -4},
		{-2, -3, -5},
		{1000000, 2345678, 3345678},
	}
	for _, c := range cases {
		if got := Add(c.a, c.b); got != c.want {
			t.Errorf("Add(%d,%d) = %d, want %d", c.a, c.b, got, c.want)
		}
	}
}

func TestAgainstPlus(t *testing.T) {
	rng := rand.New(rand.NewSource(11))
	for i := 0; i < 2000; i++ {
		a := rng.Intn(2000001) - 1000000
		b := rng.Intn(2000001) - 1000000
		if Add(a, b) != a+b {
			t.Fatalf("Add(%d,%d) != %d", a, b, a+b)
		}
	}
}
