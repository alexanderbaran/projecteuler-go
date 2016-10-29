package main

// import (
// 	"fmt"
// 	"strconv"
// 	"time"
// )

// func main() {
// 	defer timeTrack(time.Now())
// 	concat := ""
// 	prod, d, max := 1, 0, 0
// 	for i := 1; i <= 10000; i++ {
// 		concat = ""
// 	Inner:
// 		for j := 1; j <= 10; j++ {
// 			prod = i * j
// 			concat = fmt.Sprintf("%s%d", concat, prod)
// 			if len(concat) > 9 {
// 				break Inner
// 			} else if len(concat) == 9 {
// 				if isPandigitalFromString(concat) {
// 					d, _ = strconv.Atoi(concat)
// 					if d > max {
// 						max = d
// 						// fmt.Println(concat, i, j)
// 					}
// 				}
// 			}
// 		}
// 	}
// 	fmt.Println(max)
// }

// func isPandigitalFromString(n string) bool {
// 	if len(n) != 9 {
// 		return false
// 	}
// 	m := map[rune]bool{}
// 	for _, v := range n {
// 		if string(v) == "0" {
// 			return false
// 		}
// 		m[v] = true
// 	}
// 	if len(m) != 9 {
// 		return false
// 	}
// 	return true
// }
