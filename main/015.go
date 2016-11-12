package main

// import (
// 	"fmt"
// 	"time"
// )

// // var cache map[string]int

// func main() {
// 	defer timeTrack(time.Now())
// 	// cache = make(map[string]int)
// 	fmt.Println(countRoutes(4, 4))
// }

// // func countRoutes(m, n int) int {
// // 	if n == 0 || m == 0 {
// // 		return 1
// // 	}
// // 	return countRoutes(m, n-1) + countRoutes(m-1, n)
// // }

// // func countRoutes(m, n int) int {
// // 	if n == 0 || m == 0 {
// // 		return 1
// // 	}
// // 	key := fmt.Sprintf("%d %d", m, n)
// // 	if v, ok := cache[key]; ok {
// // 		return v
// // 	}
// // 	cache[key] = countRoutes(m, n-1) + countRoutes(m-1, n)
// // 	return cache[key]
// // }

// func countRoutes(m, n int) int {
// 	grid := make([][]int, m+1)
// 	for i := range grid {
// 		grid[i] = make([]int, n+1)
// 	}
// 	for i := 0; i <= m; i++ {
// 		grid[i][0] = 1
// 	}
// 	for j := 0; j <= n; j++ {
// 		grid[0][j] = 1
// 	}
// 	for i := 1; i <= m; i++ {
// 		for j := 1; j <= n; j++ {
// 			grid[i][j] = grid[i-1][j] + grid[i][j-1]
// 		}
// 	}
// 	return grid[m][n]
// }

// // 2x2 : 6 = 4! / (2! * 2!) = 24 / 4
// // 3x3 : 20 = 6! / (3! * 3!) = 720 / 36
// // 4x4 : 70 = 8! / (4! * 4!) = 40320 / 576 = 70
// // 5x5 : 252
// // 6x6 : 924
// // 15x15 : 155117520
// // 17x17 : 2333606220
// // 20x20 : 137846528820 = 40! / (20! * 20!)

// // func main() {
// // 	defer timeTrack(time.Now())
// // 	i := 0
// // 	j := 0
// // 	end := 2
// // 	count := 0
// // 	path(i, j, end, &count)
// // 	fmt.Println(count)
// // }

// // func path(i, j, end int, count *int) {
// // 	if i == end && j == end {
// // 		*count++
// // 		return
// // 	}
// // 	if i < end {
// // 		path(i+1, j, end, count)
// // 	}
// // 	if j < end {
// // 		path(i, j+1, end, count)
// // 	}
// // }
