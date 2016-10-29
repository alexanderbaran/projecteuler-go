package main

// import (
// 	"fmt"
// 	"math"
// 	"time"
// )

// func main() {
// 	defer timeTrack(time.Now())
// 	n := 2500
// 	m := map[int]bool{}
// 	for i := 1; i <= n; i++ {
// 		m[P(i)] = true
// 	}
// 	// sumOk := false
// 	// diffOk := false
// 	for i := 1; i <= n; i++ {
// 		for j := i; j > 1; j-- {
// 			// sumOk = false
// 			// diffOk = false
// 			diff := int(math.Abs(float64(P(i) - P(j))))
// 			sum := P(i) + P(j)
// 			if isPentagonal(diff, m) && isPentagonal(sum, m) {
// 				fmt.Println(diff)
// 				return
// 			}
// 			// for l := int(math.Max(float64(i), float64(j))); l > 1; l-- {
// 			//     if diff == P(l) {
// 			//         diffOk = true
// 			//         sum := P(i) + P(j)
// 			//         for k := int(math.Max(float64(i), float64(j))); k <= n; k++ {
// 			//             if sum == P(k) {
// 			//                 sumOk = true

// 			//             }
// 			//         }
// 			//     }
// 			// }
// 			// if sumOk && diffOk {
// 			//     fmt.Println(i, j, diff)
// 			// }
// 		}
// 	}
// }

// // func isPentagonal(x int) bool {
// //     n := (1 + math.Sqrt(float64(1+24*x))) / 6
// //     return math.Ceil(n) == n
// // }

// func isPentagonal(x int, m map[int]bool) bool {
// 	if _, ok := m[x]; ok {
// 		return true
// 	}
// 	return false
// }

// func P(n int) int {
// 	return n * (3*n - 1) / 2
// }
