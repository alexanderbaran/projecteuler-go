package main

// import (
// 	"fmt"
// 	"sort"
// 	"time"
// )

// func main() {
// 	defer timeTrack(time.Now())
// 	for i := 1; i <= 1000000; i++ {
// 		if containsSameDigits(i, 2*i) {
// 			if containsSameDigits(i, 3*i) {
// 				if containsSameDigits(i, 4*i) {
// 					if containsSameDigits(i, 5*i) {
// 						if containsSameDigits(i, 6*i) {
// 							// fmt.Println(i, 2*i, 3*i, 4*i, 5*i, 6*i)
// 							fmt.Println(i)
// 							return
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// }

// func containsSameDigits(x, y int) bool {
// 	if digitLength(x) != digitLength(y) {
// 		return false
// 	}
// 	xs, ys := []int{}, []int{}
// 	for x > 0 && y > 0 {
// 		xs = append(xs, x%10)
// 		ys = append(ys, y%10)
// 		x = x / 10
// 		y = y / 10
// 	}
// 	sort.Ints(xs)
// 	sort.Ints(ys)
// 	for i := 0; i < len(xs); i++ {
// 		if xs[i] != ys[i] {
// 			return false
// 		}
// 	}
// 	return true
// }
