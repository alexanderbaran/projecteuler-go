package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	defer timeTrack(time.Now())
// 	count := 0
// 	max := 0
// 	maxA := 0
// 	maxB := 0
// 	for a := -999; a <= 999; a++ {
// 		for b := -1000; b <= 1000; b++ {
// 			count = 0
// 			for n := 0; ; n++ {
// 				if isPrime(n*n + a*n + b) {
// 					count++
// 				} else {
// 					break
// 				}
// 			}
// 			if count > max {
// 				max = count
// 				maxA = a
// 				maxB = b
// 			}
// 		}
// 	}
// 	fmt.Println(maxA * maxB)
// }
