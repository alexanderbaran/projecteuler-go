package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"math"
// 	"strings"
// 	"time"
// )

// func main() {
// 	defer timeTrack(time.Now())
// 	b, err := ioutil.ReadFile("042.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	b = b[1 : len(b)-1]
// 	arr := strings.Split(string(b), `","`)
// 	aplhabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
// 	alphaMap := map[string]int{}
// 	i := 1
// 	for _, a := range aplhabet {
// 		alphaMap[string(a)] = i
// 		i++
// 	}
// 	sum := 0
// 	count := 0
// 	for _, v := range arr {
// 		sum = 0
// 		for _, s := range v {
// 			sum += alphaMap[string(s)]
// 		}
// 		if isTriangular(sum) {
// 			count++
// 		}
// 	}
// 	fmt.Println(count)
// }

// func isTriangular(x int) bool {
// 	n := (math.Sqrt(float64(1+8*x)) - 1) / 2
// 	return math.Ceil(n) == n
// }
