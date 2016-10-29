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
// 	n := 1000
// 	rat := map[string]bool{}
// 	for a := 1; a <= n; a++ {
// 		for b := 1; b <= n-a; b++ {
// 			for c := 1; c <= n-a-b; c++ {
// 				if a*a+b*b == c*c {
// 					s := []int{a, b, c}
// 					sort.Ints(s)
// 					rat[fmt.Sprintf("%d %d %d", s[0], s[1], s[2])] = true
// 				}
// 			}
// 		}
// 	}
// 	count := 0
// 	max := 0
// 	maxp := 0
// 	for p := 1; p <= 1000; p++ {
// 		count = 0
// 		for s, _ := range rat {
// 			sl := strings.Split(s, " ")
// 			a, _ := strconv.Atoi(sl[0])
// 			b, _ := strconv.Atoi(sl[1])
// 			c, _ := strconv.Atoi(sl[2])
// 			if a+b+c == p {
// 				count++
// 			}
// 			if count > max {
// 				max = count
// 				maxp = p
// 			}
// 		}
// 	}
// 	fmt.Println(maxp)
// }
