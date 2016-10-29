package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	defer timeTrack(time.Now())
// 	md := make([]int, 12)
// 	md[0] = 31  // Jan
// 	md[1] = 28  // Feb
// 	md[2] = 31  // Mar
// 	md[3] = 30  // Apr
// 	md[4] = 31  // May
// 	md[5] = 30  // Jun
// 	md[6] = 31  // Jul
// 	md[7] = 31  // Aug
// 	md[8] = 30  // Sep
// 	md[9] = 31  // Oct
// 	md[10] = 30 // Nov
// 	md[11] = 31 // Dec
// 	count := 0
// 	track := 0
// 	for y := 1900; y <= 2000; y++ {
// 		leap := false
// 		if y%4 == 0 && y%100 != 0 {
// 			leap = true
// 		} else if y%100 == 0 && y%400 == 0 {
// 			leap = true
// 		}
// 		for i, v := range md {
// 			end := v
// 			if leap && i == 1 {
// 				end++
// 			}
// 			for d := 1; d <= end; d++ {
// 				track++
// 				if d == 1 && track%7 == 0 && y >= 1901 {
// 					count++
// 				}
// 			}
// 		}
// 	}
// 	fmt.Println(count)
// }
