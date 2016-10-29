package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	defer timeTrack(time.Now())
// 	at := 1
// 	for i := 2; ; i++ {
// 		if divisors(at) > 500 {
// 			fmt.Println(at)
// 			return
// 		}
// 		at = at + i
// 	}
// }

// func divisors(n int) int {
// 	amount := 0
// 	for i := 1; i*i <= n; i++ {
// 		if n%i == 0 {
// 			if n%i != i {
// 				amount += 2
// 			} else {
// 				amount++
// 			}
// 		}
// 	}
// 	return amount
// }
