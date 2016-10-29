package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	defer timeTrack(time.Now())
// 	longest := 0
// 	index := 1
// 	for i := 1; i < 1e6; i++ {
// 		if chainLenth(i) > longest {
// 			index = i
// 			longest = chainLenth(i)
// 		}
// 	}
// 	fmt.Println(index)
// }

// func chainLenth(n int) int {
// 	length := 1
// 	for {
// 		if isEven(n) {
// 			n = even(n)
// 		} else {
// 			n = odd(n)
// 		}
// 		length++
// 		if n == 1 {
// 			return length
// 		}
// 	}
// }

// func isEven(n int) bool {
// 	if n%2 == 0 {
// 		return true
// 	}
// 	return false
// }

// func even(n int) int {
// 	return n / 2
// }

// func odd(n int) int {
// 	return 3*n + 1
// }
