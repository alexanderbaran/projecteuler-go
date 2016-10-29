package main

// import (
// 	"fmt"
// 	"strconv"
// 	"time"
// )

// func main() {
// 	defer timeTrack(time.Now())
// 	numerator, denominator := 1, 1
// 	for i := 11; i <= 99; i++ {
// 		for j := 11; j < i; j++ {
// 			if i%10 == 0 || j%10 == 0 {
// 				continue
// 			}
// 			is := strconv.Itoa(i)
// 			js := strconv.Itoa(j)
// 			x, y := 0.0, 0.0
// 			if is[0] == js[0] {
// 				x, _ = strconv.ParseFloat(string(is[1]), 64)
// 				y, _ = strconv.ParseFloat(string(js[1]), 64)
// 			} else if is[0] == js[1] {
// 				x, _ = strconv.ParseFloat(string(is[1]), 64)
// 				y, _ = strconv.ParseFloat(string(js[0]), 64)
// 			} else if is[1] == js[0] {
// 				x, _ = strconv.ParseFloat(string(is[0]), 64)
// 				y, _ = strconv.ParseFloat(string(js[1]), 64)
// 			} else if is[1] == js[1] {
// 				x, _ = strconv.ParseFloat(string(is[0]), 64)
// 				y, _ = strconv.ParseFloat(string(js[0]), 64)
// 			}
// 			if float64(j)/float64(i) == y/x {
// 				// fmt.Printf("%d/%d\n", j, i)
// 				numerator *= int(y)
// 				denominator *= int(x)
// 			}
// 		}
// 	}
// 	// fmt.Printf("%d/%d\n", numerator/numerator, denominator/numerator)
// 	fmt.Printf("%d\n", denominator/numerator)
// }
