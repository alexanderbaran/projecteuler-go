package main

// import (
// 	"fmt"
// 	"math"
// 	"strconv"
// 	"time"
// )

// func main() {
// 	defer timeTrack(time.Now())
// 	s, d, sum, totalsum := "", 0, 0, 0
// 	for i := 10; i <= 200000; i++ {
// 		s = strconv.Itoa(i)
// 		sum = 0
// 		for _, v := range s {
// 			d, _ = strconv.Atoi(string(v))
// 			sum += int(math.Pow(float64(d), 5))
// 		}
// 		if sum == i {
// 			totalsum += i
// 		}
// 	}
// 	fmt.Println(totalsum)
// }
