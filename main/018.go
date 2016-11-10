package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type edge struct {
	to     int
	weight int
}

func main() {
	ints := []int{}
	f, err := os.Open("018.txt")
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		for _, v := range strings.Split(scanner.Text(), " ") {
			d, _ := strconv.Atoi(v)
			ints = append(ints, d)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	adj := make([][]edge, len(ints))

	test := 0
	incr := 1
	for i := 0; i <= len(adj); {
		if i >= test && i < test+incr {
			// fmt.Println(i, i+incr)
			// fmt.Println(i, i+incr+1)
			left := edge{to: i + incr, weight: ints[i+incr]}
			right := edge{to: i + incr + 1, weight: ints[i+incr+1]}
			edges := []edge{left, right}
			adj[i] = edges
			i++
		} else {
			test += incr
			incr++
		}
		if test+incr == len(adj) {
			break
		}
	}
	// spew.Dump(adj)

	// Dijkstra (longest path instead of shortest path)
	Q := map[int]bool{}
	dist := make([]int, len(ints))
	prev := make([]int, len(ints))
	for i := 0; i < len(ints); i++ {
		dist[i] = -1
		prev[i] = -1 // Undefined
		Q[i] = true
	}
	dist[0] = 0
	// spew.Dump(Q)

	// for len(Q) > 0 {
	// 	u := maxI(dist)
	// 	delete(Q, u)

	// 	// for _, v := range adj[u] {
	// 	// 	fmt.Println(v)
	// 	// }
	// }
}

func maxI(a []int) int {
	maxV := -1
	maxI := 0
	for i, v := range a {
		if v > maxV {
			maxV = v
			maxI = i
		}
	}
	return maxI
}
