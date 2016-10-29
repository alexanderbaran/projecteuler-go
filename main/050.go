package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	defer timeTrack(time.Now())
// 	sum := 0
// 	maxSum := 0
// 	count := 0
// 	maxCount := 0
// 	for i := 2; i < 10; i += 1 {
// 		count = 0
// 		sum = 0
// 		for j := i; ; j += 1 {
// 			if isPrime(j) {
// 				sum += j
// 				if sum >= 1000000 {
// 					break
// 				}
// 				count++
// 				if isPrime(sum) {
// 					if count > maxCount {
// 						maxCount = count
// 						maxSum = sum
// 					}
// 				}
// 			}
// 		}
// 	}
// 	fmt.Println(maxSum)
// }
