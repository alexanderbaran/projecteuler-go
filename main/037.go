package main

// import (
// 	"fmt"
// 	"strconv"
// 	"time"
// )

// func main() {
// 	defer timeTrack(time.Now())
// 	count := 0
// 	sum := 0
// 	for i := 11; ; i += 2 {
// 		if isPrime(i) {
// 			if isTruncatablePrime(i) {
// 				// fmt.Println(i)
// 				count++
// 				sum += i
// 			}
// 		}
// 		if count == 11 {
// 			break
// 		}
// 	}
// 	fmt.Println(sum)
// }

// func isTruncatablePrime(n int) bool {
// 	s := strconv.Itoa(n)
// 	for i := 1; i < len(s); i++ {
// 		d, _ := strconv.Atoi(s[i:])
// 		if !isPrime(d) {
// 			return false
// 		}
// 	}
// 	for i := 1; i < len(s); i++ {
// 		d, _ := strconv.Atoi(s[:i])
// 		if !isPrime(d) {
// 			return false
// 		}
// 	}
// 	return true
// }
