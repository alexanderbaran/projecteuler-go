package main

// import (
// 	"fmt"
// 	"math/big"
// 	"time"
// )

// func main() {
// 	defer timeTrack(time.Now())
// 	count := 0
// 	for i := 0; i < 10000; i++ {
// 		if isLychrel(int64(i)) {
// 			count++
// 		}
// 	}
// 	fmt.Println(count)
// }

// func isLychrel(n int64) bool {
// 	bi := big.NewInt(n)
// 	for i := 1; i < 50; i++ {
// 		bi = bi.Add(bi, reverseBigInt(bi))
// 		if isPalindromeFromString(bi.String()) {
// 			return false
// 		}
// 	}
// 	return true
// }

// func reverseBigInt(bi *big.Int) *big.Int {
// 	revstr := reverseString(bi.String())
// 	newInt := &big.Int{}
// 	newInt.SetString(revstr, 10)
// 	return newInt
// }
