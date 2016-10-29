package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"sort"
// 	"strings"
// 	"time"
// )

// func main() {
// 	defer timeTrack(time.Now())
// 	b, err := ioutil.ReadFile("022.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	b = b[1 : len(b)-1]
// 	var strSlice sort.StringSlice
// 	strSlice = strings.Split(string(b), `","`)
// 	strSlice.Sort()
// 	aplhabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
// 	alphaMap := map[string]int{}
// 	i := 1
// 	for _, a := range aplhabet {
// 		alphaMap[string(a)] = i
// 		i++
// 	}
// 	totalScore := 0
// 	for i, v := range strSlice {
// 		namescore := 0
// 		for _, s := range v {
// 			namescore += alphaMap[string(s)]
// 		}
// 		totalScore += namescore * (i + 1)
// 	}
// 	fmt.Println(totalScore)
// }
