package main

// import (
// 	"fmt"
// 	"sort"
// 	"strconv"
// 	"strings"
// 	"time"
// )

// func main() {
// 	defer timeTrack(time.Now())
// 	max := 200
// 	m := map[string]bool{}
// 	s := make([][]int, max+1)
// 	s[200] = []int{100, 100}
// 	s[100] = []int{50, 50}
// 	s[50] = []int{20, 20, 10}
// 	s[20] = []int{10, 10}
// 	s[10] = []int{5, 5}
// 	s[5] = []int{2, 2, 1}
// 	s[2] = []int{1, 1}
// 	s[1] = []int{}
// 	m[strconv.Itoa(max)] = true
// 	combinations(s[max], s, m, max)
// 	for i, v := range s {
// 		if v != nil && i != max && i != 1 {
// 			if max%i == 0 {
// 				d := max / i
// 				arr := []int{}
// 				for k := 1; k <= d; k++ {
// 					arr = append(arr, i)
// 				}
// 				if _, ok := m[format(arr)]; !ok {
// 					combinations(arr, s, m, max)
// 				}
// 			}
// 		}
// 	}
// 	fmt.Println(len(m))
// }

// func combinations(slice []int, s [][]int, m map[string]bool, max int) {
// 	m[format(slice)] = true
// 	copied := make([]int, len(slice))
// 	copy(copied, slice)
// 	for i, v := range copied {
// 		if v != 1 {
// 			tmp := make([]int, len(copied))
// 			copy(tmp, copied)
// 			tmp = append(tmp[:i], tmp[i+1:]...)
// 			tmp = append(s[v], tmp...)
// 			if _, ok := m[format(tmp)]; !ok {
// 				combinations(tmp, s, m, max)
// 			}
// 		}
// 		if v == 10 {
// 			tmp := make([]int, len(copied))
// 			copy(tmp, copied)
// 			tmp = append(tmp[:i], tmp[i+1:]...)
// 			tmp = append([]int{2, 2, 2, 2, 2}, tmp...)
// 			if _, ok := m[format(tmp)]; !ok {
// 				combinations(tmp, s, m, max)
// 			}
// 		}
// 		if v == 100 {
// 			tmp := make([]int, len(copied))
// 			copy(tmp, copied)
// 			tmp = append(tmp[:i], tmp[i+1:]...)
// 			tmp = append([]int{20, 20, 20, 20, 20}, tmp...)
// 			if _, ok := m[format(tmp)]; !ok {
// 				combinations(tmp, s, m, max)
// 			}
// 		}
// 	}
// }

// func format(slice []int) string {
// 	tmp := make([]int, len(slice))
// 	copy(tmp, slice)
// 	sort.Ints(tmp)
// 	strArr := make([]string, len(tmp))
// 	for i := 0; i < len(strArr); i++ {
// 		strArr[i] = strconv.Itoa(tmp[i])
// 	}
// 	return strings.Join(strArr, " ")
// }

// func main() {
// 	defer timeTrack(time.Now())
// 	sum, count := 200, 0
// 	for c200 := 0; c200 <= sum; c200 += 200 {
// 		for c100 := 0; c100 <= sum; c100 += 100 {
// 			for c50 := 0; c50 <= sum; c50 += 50 {
// 				for c20 := 0; c20 <= sum; c20 += 20 {
// 					for c10 := 0; c10 <= sum; c10 += 10 {
// 						for c5 := 0; c5 <= sum; c5 += 5 {
// 							for c2 := 0; c2 <= sum; c2 += 2 {
// 								for c1 := 0; c1 <= sum; c1 += 1 {
// 									if c200+c100+c50+c20+c10+c5+c2+c1 == sum {
// 										count++
// 									}
// 								}
// 							}
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// 	fmt.Println(count)
// }
