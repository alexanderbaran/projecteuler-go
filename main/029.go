package main

// import (
// 	"fmt"
// 	"math/big"
// 	"time"
// )

// func main() {
// 	defer timeTrack(time.Now())
// 	m := map[string]bool{}
// 	for a := 2; a <= 100; a++ {
// 		for b := 2; b <= 100; b++ {
// 			term := powCustom(a, b)
// 			m[term] = true
// 		}
// 	}
// 	fmt.Println(len(m))
// }

// func powCustom(x, y int) string {
// 	xi := big.NewInt(int64(x))
// 	result := big.NewInt(1)
// 	for i := 1; i <= y; i++ {
// 		result = result.Mul(result, xi)
// 	}
// 	return result.String()
// }
