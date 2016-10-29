package main

// import (
// 	"fmt"
// 	"math"
// 	"time"
// )

// const (
// 	RIGHT = 0
// 	DOWN  = 1
// 	LEFT  = 2
// 	UP    = 3
// )

// var arr [1001][1001]int

// func main() {
// 	defer timeTrack(time.Now())
// 	n := 1001
// 	x := int(math.Floor(float64(n / 2)))
// 	y := x
// 	arr[x][y] = 1
// 	i, d, step := 2, 0, 0
// Outer:
// 	for {
// 		for j := 0; ; j++ {
// 			d = j % 4
// 			if j%2 == 0 {
// 				step++
// 			}
// 			for s := 1; s <= step; s++ {
// 				if i > n*n {
// 					break Outer
// 				}
// 				switch d {
// 				case RIGHT:
// 					x, y = right(x, y, i)
// 				case DOWN:
// 					x, y = down(x, y, i)
// 				case LEFT:
// 					x, y = left(x, y, i)
// 				case UP:
// 					x, y = up(x, y, i)
// 				}
// 				i++
// 			}
// 		}
// 	}
// 	sum := 0
// 	for i := 0; i < n; i++ {
// 		sum += arr[i][i]
// 		sum += arr[len(arr)-1-i][i]
// 	}
// 	sum--
// 	fmt.Println(sum)
// 	// arrPrintln()
// }

// func right(x, y int, i int) (int, int) {
// 	arr[x][y+1] = i
// 	return x, y + 1
// }

// func down(x, y int, i int) (int, int) {
// 	arr[x+1][y] = i
// 	return x + 1, y
// }

// func left(x, y int, i int) (int, int) {
// 	arr[x][y-1] = i
// 	return x, y - 1
// }

// func up(x, y int, i int) (int, int) {
// 	arr[x-1][y] = i
// 	return x - 1, y
// }

// // func arrPrintln() {
// // 	for _, v := range arr {
// // 		fmt.Println(v)
// // 	}
// // }
