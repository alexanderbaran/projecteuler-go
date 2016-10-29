package main

// import (
// 	"bytes"
// 	"fmt"
// 	"strconv"
// 	"time"
// )

// var buffer bytes.Buffer

// func main() {
// 	defer timeTrack(time.Now())
// 	n := 1000
// 	max := 0
// 	d := 0
// 	precision := 2000
// 	numerator := 1
// 	for i := 2; i <= n; i++ {
// 		denominator := i
// 		remainder := numerator
// 		divisor := denominator
// 		buffer.Reset()
// 		for i := 1; i <= precision && remainder != 0; i++ {
// 			dividend := remainder * 10
// 			quotient := dividend / divisor
// 			remainder = dividend % divisor
// 			buffer.WriteString(strconv.Itoa(quotient))
// 		}
// 		t := buffer.String()
// 		count := 0
// 		for j := 0; j <= len(t)-1; j++ {
// 		K:
// 			for k := j + 1; k <= len(t)-1; k++ {
// 				if t[j] == t[k] {
// 					count = 0
// 					for z := 1; z <= k-j-1 && k+z <= len(t)-1; z++ {
// 						if t[j+z] == t[k+z] {
// 							count++
// 						} else {
// 							continue K
// 						}
// 					}
// 					if count == k-j-1 {
// 						count++
// 					} else {
// 						count = 0
// 					}
// 					if count > max {
// 						max = count
// 						d = i
// 					}
// 					break
// 				}
// 			}
// 		}
// 	}
// 	fmt.Println(d)
// }
