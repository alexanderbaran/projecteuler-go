package main

// import (
// 	"fmt"
// 	"strconv"
// 	"time"
// )

// func main() {
// 	defer timeTrack(time.Now())
// 	counter := 1
// 	for i := 3; i <= 1000000; i += 2 {
// 		if isPrime(i) {
// 			if isCircularPrime(i) {
// 				counter++
// 			}
// 		}
// 	}
// 	fmt.Println(counter)
// }

// func isCircularPrime(n int) bool {
// 	// Assuming n is prime
// 	if n > 0 && n <= 9 {
// 		return true
// 	}
// 	s, d := strconv.Itoa(n), 0
// 	for i := 1; i < len(s); i++ {
// 		s = s[1:len(s)] + s[0:1]
// 		d, _ = strconv.Atoi(s)
// 		if !isPrime(d) {
// 			return false
// 		}
// 	}
// 	return true
// }
