package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	defer timeTrack(time.Now())
// 	n, d, sum, result := 0, 0, 0, 0
// 	for i := 10; i <= 500000; i++ {
// 		n = i
// 		sum = 0
// 		for {
// 			if n <= 0 {
// 				break
// 			}
// 			d = n % 10
// 			sum += factorial(d)
// 			n = n / 10
// 		}
// 		if sum == i {
// 			result += i
// 		}
// 	}
// 	fmt.Println(result)
// }

// func factorial(n int) int {
// 	prod := 1
// 	for i := n; i >= 1; i-- {
// 		prod *= i
// 	}
// 	return prod
// }
