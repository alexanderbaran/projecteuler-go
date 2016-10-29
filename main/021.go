package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	defer timeTrack(time.Now())
// 	m := map[int]bool{}
// 	for i := 1; i < 10000; i++ {
// 		tmp := divisorsSum(i)
// 		if divisorsSum(tmp) == i {
// 			if tmp != i {
// 				m[i] = true
// 				m[tmp] = true
// 			}
// 		}
// 	}
// 	sum := 0
// 	for a, _ := range m {
// 		sum += a
// 	}
// 	fmt.Println(sum)
// }
