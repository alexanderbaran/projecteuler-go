package main

// import (
// 	"fmt"
// 	"strconv"
// 	"time"
// )

// func main() {
// 	defer timeTrack(time.Now())
// 	max := 0
// 	for i := 11; i <= 10000000; i += 2 {
// 		if isPandigital(i) {
// 			if isPrime(i) {
// 				if i > max {
// 					max = i
// 				}
// 			}
// 		}
// 	}
// 	fmt.Println(max)
// }

// func isPandigital(n int) bool {
// 	m := map[int]bool{}
// 	s := strconv.Itoa(n)
// 	length, d := len(s), 0
// 	for {
// 		if n <= 0 {
// 			break
// 		}
// 		d = n % 10
// 		if d == 0 || d > length {
// 			return false
// 		}
// 		m[d] = true
// 		n = n / 10
// 	}
// 	if len(m) != length {
// 		return false
// 	}
// 	return true
// }
