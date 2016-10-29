package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	defer timeTrack(time.Now())
// 	for i := 1009; i <= 9973; i += 2 {
// 		if isPrime(i) {
// 		Inner:
// 			for j := 1; i+(2*j) < 10000; j++ {
// 				second := i + j
// 				third := i + 2*j
// 				if !isPrime(second) {
// 					continue Inner
// 				}
// 				if !isPrime(third) {
// 					continue Inner
// 				}
// 				if !arePermutations(i, second) {
// 					continue Inner
// 				}
// 				if !arePermutations(i, third) {
// 					continue Inner
// 				}
// 				fmt.Printf("i: %d, j: %d, i+j: %d, i+2*j: %d, concat: %d%d%d\n", i, j, i+j, i+2*j, i, i+j, i+2*j)
// 			}
// 		}
// 	}
// }

// func arePermutations(x, y int) bool {
// 	if x == y {
// 		return false
// 	}
// 	if digitLength(x) != digitLength(y) {
// 		return false
// 	}
// 	xm := map[int]bool{}
// 	for ; x > 0; x = x / 10 {
// 		xm[x%10] = true
// 	}
// 	ym := map[int]bool{}
// 	for ; y > 0; y = y / 10 {
// 		ym[y%10] = true
// 	}
// 	found := false
// 	for i, _ := range xm {
// 		found = false
// 		for j, _ := range ym {
// 			if i == j {
// 				found = true
// 			}
// 		}
// 		if !found {
// 			return false
// 		}
// 	}
// 	if len(xm) != len(ym) {
// 		return false
// 	}
// 	return true
// }
