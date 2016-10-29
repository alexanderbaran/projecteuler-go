package main

// import (
// 	"fmt"
// 	"math/big"
// 	"time"
// )

// func main() {
// 	defer timeTrack(time.Now())
// 	n := 1000
// 	sum := big.NewInt(0)
// 	p := big.NewInt(1)
// 	for i := 1; i <= n; i++ {
// 		p = big.NewInt(1)
// 		for j := 1; j <= i; j++ {
// 			bi := big.NewInt(int64(i))
// 			p = p.Mul(p, bi)
// 		}
// 		sum = sum.Add(sum, p)
// 	}
// 	t := sum.Text(10)
// 	fmt.Println(t[len(t)-10 : len(t)])
// }
