package main

// import (
// 	"fmt"
// 	"math"
// 	"time"
// )

// func main() {
// 	defer timeTrack(time.Now())
// 	for i := 1; ; i++ {
// 		if isPentagonal(H(i)) {
// 			if H(i) > 40755 {
// 				fmt.Println(H(i))
// 				break
// 			}
// 		}
// 	}
// 	// n := 60000
// 	// tm := map[int]bool{}
// 	// pm := map[int]bool{}
// 	// hm := map[int]bool{}
// 	// for i := 1; i <= n; i++ {
// 	//     tm[T(i)] = true
// 	//     pm[P(i)] = true
// 	//     hm[H(i)] = true
// 	// }
// 	// for i := 1; i <= T(n); i++ {
// 	//     if intInMap(i, tm) && intInMap(i, pm) && intInMap(i, hm) {
// 	//         fmt.Println(i)
// 	//     }
// 	// }
// 	// for i := 285; i <= n; i++ {
// 	//     for j := 165; j <= i; j++ {
// 	//         for k := 143; k <= j; k++ {
// 	//             if T(i) == P(j) && T(i) == H(k) {
// 	//                 fmt.Printf("T(%d) = P(%d) = H(%d) = %d\n", i, j, k, T(i))
// 	//             }
// 	//         }
// 	//     }
// 	// }
// }

// func isPentagonal(x int) bool {
// 	n := (1 + math.Sqrt(float64(1+24*x))) / 6
// 	return math.Ceil(n) == n
// }

// func intInMap(x int, m map[int]bool) bool {
// 	if _, ok := m[x]; ok {
// 		return true
// 	}
// 	return false
// }

// func T(n int) int {
// 	return n * (n + 1) / 2
// }

// func P(n int) int {
// 	return n * (3*n - 1) / 2
// }

// func H(n int) int {
// 	return n * (2*n - 1)
// }
