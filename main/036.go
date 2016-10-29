package main

// import (
// 	"bytes"
// 	"fmt"
// 	"strconv"
// 	"time"
// )

// func main() {
// 	defer timeTrack(time.Now())
// 	sum := 0
// 	for i := 1; i < 1e6; i++ {
// 		s := strconv.Itoa(i)
// 		if isPalindromeFromString(s) {
// 			if isPalindromeFromString(base10Tobase2(i)) {
// 				sum += i
// 			}
// 		}
// 	}
// 	fmt.Println(sum)
// }

// var buffer bytes.Buffer

// func base10Tobase2(n int) string {
// 	buffer.Reset()
// 	for n >= 1 {
// 		buffer.WriteString(strconv.Itoa(n % 2))
// 		n = n / 2
// 	}
// 	return reverseString(buffer.String())
// }
