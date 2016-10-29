package main

// import (
// 	"bytes"
// 	"fmt"
// 	"strconv"
// 	"time"
// )

// var str string
// var buffer bytes.Buffer

// func main() {
// 	start := time.Now()
// 	defer func() {
// 		elapsed := time.Since(start)
// 		fmt.Printf("%s\n", elapsed)
// 	}()

// 	counter := 0
// 	for i := 1; ; i++ {
// 		s := strconv.Itoa(i)
// 		buffer.WriteString(s)
// 		counter += len(s)
// 		if counter >= 1000000 {
// 			break
// 		}
// 	}
// 	str = buffer.String()
// 	fmt.Println(d(1) * d(10) * d(100) * d(1000) * d(10000) * d(100000) * d(1000000))
// }

// func d(n int) int {
// 	d, _ := strconv.Atoi(string(str[n-1]))
// 	return d
// }
