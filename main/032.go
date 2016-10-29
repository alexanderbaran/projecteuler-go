package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	defer timeTrack(time.Now())
// 	s := ""
// 	n, length, sum := 2000, 0, 0
// 	m := map[int]bool{}
// 	for i := 1; i <= n; i++ {
// 	Inner:
// 		for j := 1; j <= n; j++ {
// 			s = fmt.Sprintf("%d%d%d", i, j, i*j)
// 			length = len(s)
// 			if length > 9 {
// 				continue Inner
// 			}
// 			if length == 9 {
// 				if isPandigitalFromString(s) {
// 					m[i*j] = true
// 				}
// 			}
// 		}
// 	}
// 	for k, _ := range m {
// 		sum += k
// 	}
// 	fmt.Println(sum)
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
