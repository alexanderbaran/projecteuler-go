package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"strconv"
// 	"strings"
// 	"time"
// )

// func main() {
// 	defer timeTrack(time.Now())
// 	b, err := ioutil.ReadFile("008.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	s := strings.Replace(string(b), "\n", "", -1)
// 	largest := 0
// 	n := 13
// 	prod := 1
// 	for i := 0; i < len(s)-n; i++ {
// 		prod = 1
// 		for k := 0; k < n; k++ {
// 			number, _ := strconv.Atoi(string(s[i+k]))
// 			prod *= number
// 		}
// 		if prod > largest {
// 			largest = prod
// 		}
// 	}
// 	fmt.Println(largest)
// }
