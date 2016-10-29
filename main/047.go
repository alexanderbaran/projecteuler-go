package main

// import (
// 	"time"

// 	"github.com/davecgh/go-spew/spew"
// )

// func main() {
// 	defer timeTrack(time.Now())
// 	count := 0
// 	end := 0
// 	result := map[int]map[int]bool{}
// 	amount, n := 4, 150000
// 	for i := 1; i <= n; i++ {
// 		end = i
// 		m := map[int]bool{}
// 		for j := 1; j < end; j++ {
// 			mod := i % j
// 			if mod == 0 {
// 				if isPrime(j) {
// 					m[j] = true
// 				}
// 				end = i / j
// 				if isPrime(end) {
// 					m[end] = true
// 				}
// 			}
// 		}
// 		if len(m) == amount {
// 			count++
// 			result[i] = m
// 		} else {
// 			count = 0
// 			result = map[int]map[int]bool{}
// 		}
// 		if count == amount {
// 			break
// 		}
// 	}
// 	spew.Dump(result)
// }
