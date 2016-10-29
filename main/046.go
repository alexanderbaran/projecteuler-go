package main

// import (
// 	"fmt"
// 	"math"
// 	"time"
// )

// func main() {
// 	defer timeTrack(time.Now())
// 	found := true
// 	last := 0
// I:
// 	for i := 3; ; i += 2 {
// 		if !isPrime(i) {
// 			if !found {
// 				fmt.Println(last)
// 				return
// 			} else {
// 				found = false
// 			}
// 			last = i
// 		J:
// 			for j := 2; j < i; j += 2 {
// 				if j == 4 {
// 					j--
// 				}
// 				if isPrime(j) {
// 					for k := 1; ; k++ {
// 						check := (j + 2*int(math.Pow(float64(k), 2)))
// 						if check > i {
// 							continue J
// 						}
// 						if i == check {
// 							found = true
// 							continue I
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// }
