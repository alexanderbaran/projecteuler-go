package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"time"
// )

// func main() {
// 	defer timeTrack(time.Now())
// 	file, err := os.Open("011.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer file.Close()
// 	arr := [][]int{}
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		row := []int{}
// 		str := scanner.Text()
// 		s := strings.Split(str, " ")
// 		for _, v := range s {
// 			d, _ := strconv.Atoi(string(v))
// 			row = append(row, d)
// 		}
// 		arr = append(arr, row)
// 	}
// 	if err := scanner.Err(); err != nil {
// 		fmt.Println(err)
// 	}
// 	result := 0
// 	maxh := 0
// 	prod := 0
// 	for i := 0; i < len(arr); i++ {
// 		for j := 0; j < len(arr[i])-3; j++ {
// 			prod = 1
// 			for k := 0; k < 4; k++ {
// 				prod *= arr[i][j+k]
// 				if prod > maxh {
// 					maxh = prod
// 				}
// 			}
// 		}
// 	}
// 	if maxh > result {
// 		result = maxh
// 	}
// 	maxv := 0
// 	for j := 0; j < len(arr); j++ {
// 		for i := 0; i < len(arr[j])-3; i++ {
// 			prod = 1
// 			for k := 0; k < 4; k++ {
// 				prod *= arr[i+k][j]
// 				if prod > maxv {
// 					maxv = prod
// 				}
// 			}
// 		}
// 	}
// 	if maxv > result {
// 		result = maxv
// 	}
// 	maxd1 := 0
// 	for i := 0; i < len(arr)-3; i++ {
// 		for j := 0; j < len(arr[i])-3; j++ {
// 			prod = 1
// 			for k := 0; k < 4; k++ {
// 				prod *= arr[i+k][j+k]
// 				if prod > maxd1 {
// 					maxd1 = prod
// 				}
// 			}
// 		}
// 	}
// 	if maxd1 > result {
// 		result = maxd1
// 	}
// 	maxd2 := 0
// 	for i := 3; i < len(arr); i++ {
// 		for j := 0; j < len(arr[i])-3; j++ {
// 			prod = 1
// 			for k := 0; k < 4; k++ {
// 				prod *= arr[i-k][j+k]
// 				if prod > maxd2 {
// 					maxd2 = prod
// 				}
// 			}
// 		}
// 	}
// 	if maxd2 > result {
// 		result = maxd2
// 	}
// 	fmt.Println(result)
// }
