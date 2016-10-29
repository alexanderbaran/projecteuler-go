package main

// import (
// 	"fmt"
// 	"strconv"
// 	"time"
// )

// func main() {
// 	defer timeTrack(time.Now())
// 	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	s := ""
// 	count := 0
// 	for i, _ := range arr {
// 		if lexicographicPermutation(i, &count, s, arr) {
// 			break
// 		}
// 	}
// }

// func lexicographicPermutation(i int, count *int, s string, arr []int) bool {
// 	if len(arr) == 1 {
// 		*count++
// 		if *count >= 1e6 {
// 			fmt.Printf("%s%d\n", s, arr[i])
// 			return true
// 		} else {
// 			return false
// 		}
// 	} else {
// 		tmp := make([]int, len(arr))
// 		copy(tmp, arr)
// 		tmp = append(tmp[:i], tmp[i+1:]...)
// 		s += strconv.Itoa(arr[i])
// 		for i, _ := range tmp {
// 			if lexicographicPermutation(i, count, s, tmp) {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }
