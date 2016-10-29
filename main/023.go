package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	defer timeTrack(time.Now())
// 	n := 28123
// 	abundant := []int{}
// 	for i := 12; i <= n; i++ {
// 		if divisorsSum(i) > i {
// 			abundant = append(abundant, i)
// 		}
// 	}
// 	sums := make([]bool, n)
// 	for i := 0; i < len(abundant); i++ {
// 		for j := i; j < len(abundant); j++ {
// 			if abundant[i]+abundant[j] < n {
// 				sums[abundant[i]+abundant[j]] = true
// 			} else {
// 				break
// 			}
// 		}
// 	}
// 	sum := 0
// 	for i := 0; i < len(sums); i++ {
// 		if !sums[i] {
// 			sum += i
// 		}
// 	}
// 	fmt.Println(sum)
// }
