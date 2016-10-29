package main

// import (
// 	"fmt"
// 	"strconv"
// 	"time"
// )

// func main() {
// 	defer timeTrack(time.Now())
// 	primes := []int{2, 3, 5, 7, 11, 13, 17}
// 	m := map[int]map[string]bool{}
// 	for _, p := range primes {
// 		mp := map[string]bool{}
// 		for i := p; i <= 999; i += p {
// 			s := fmt.Sprintf("%03d", i)
// 			if hasDistinctDigitsFromString(s) {
// 				mp[s] = true
// 				m[p] = mp
// 			}
// 		}
// 	}
// 	sum := 0
// 	for s17, _ := range m[17] {
// 		for s13, _ := range m[13] {
// 			if s17[0] == s13[1] && s17[1] == s13[2] {
// 				for s11, _ := range m[11] {
// 					if s13[0] == s11[1] && s13[1] == s11[2] {
// 						for s7, _ := range m[7] {
// 							if s11[0] == s7[1] && s11[1] == s7[2] {
// 								for s5, _ := range m[5] {
// 									if s7[0] == s5[1] && s7[1] == s5[2] {
// 										for s3, _ := range m[3] {
// 											if s5[0] == s3[1] && s5[1] == s3[2] {
// 												for s2, _ := range m[2] {
// 													if s3[0] == s2[1] && s3[1] == s2[2] {
// 														for i := 1; i <= 9; i++ {
// 															s := fmt.Sprintf("%d%s%s%s%s%s%s%s%s%s", i, string(s2[0]), string(s3[0]), string(s5[0]), string(s7[0]), string(s11[0]), string(s13[0]), string(s17[0]), string(s17[1]), string(s17[2]))
// 															if hasDistinctDigitsFromString(s) {
// 																d, _ := strconv.Atoi(s)
// 																sum += d
// 															}
// 														}
// 													}
// 												}
// 											}
// 										}
// 									}
// 								}
// 							}
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// 	fmt.Println(sum)
// }
