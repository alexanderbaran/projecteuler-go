package main

import (
	"fmt"
	"time"
)

var primeTable []bool

func main() {
	defer timeTrack(time.Now())
	target := 8
	n := int(1e6)
	primeTable = make([]bool, n)
	for i := 3; i <= n; i += 2 {
		if isPrime(i) {
			primeTable[i] = true
		}
	}
	for i := 1; i < n; i++ {
		o := splitInt(i)
		s := make([]int, len(o))
		for j := 0; j < len(o); j++ {
			if replaceAndTest(target, o, s, j) {
				return
			}
		}
		for j := 0; j < len(o); j++ {
			for k := j + 1; k < len(o); k++ {
				if replaceAndTest(target, o, s, j, k) {
					return
				}
			}
		}
		for j := 0; j < len(o); j++ {
			for k := j + 1; k < len(o); k++ {
				for l := k + 1; l < len(o); l++ {
					if replaceAndTest(target, o, s, j, k, l) {
						return
					}
				}
			}
		}
	}
}

func replaceAndTest(target int, o, s []int, idx ...int) bool {
	count := 0
	fails := 0
	primes := make([]int, 8)
	for r := 0; r <= 9; r++ {
		copy(s, o)
		for _, i := range idx {
			s[i] = r
		}
		d := joinInts(s)
		if primeTable[d] && digitLength(d) == len(s) {
			primes[count] = d
			count++
		} else {
			fails++
		}
		if fails > 10-target {
			return false
		}
		if count >= target {
			fmt.Println(primes[0])
			return true
		}
	}
	return false
}

func splitInt(n int) []int {
	a := make([]int, digitLength(n))
	for i := len(a) - 1; n > 0; i-- {
		a[i] = n % 10
		n = n / 10
	}
	return a
}

func joinInts(a []int) int {
	n := 0
	for i := 0; i < len(a); i++ {
		n = 10*n + a[i]
	}
	return n
}
