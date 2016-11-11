package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type edge struct {
	weight int
	to     int
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
	notKnown := map[int]bool{}
	dist := make([]int, len(ints))
	prev := make([]int, len(ints))
	for i := 0; i < len(ints); i++ {
		dist[i] = -1
		prev[i] = -1 // Undefined
		notKnown[i] = false
	}
	dist[0] = ints[0]
	// dist[2] = 3
	// delete(notKnown, 2)
	// fmt.Println(maxOfNotKnown(dist, notKnown))

	for len(notKnown) > 0 {
		u := maxOfNotKnown(dist, notKnown)
		delete(notKnown, u)

		for _, edge := range adj[u] {
			// fmt.Println(u, edge)
			// if _, ok := notKnown[edge.to]; ok {
			alt := dist[u] + edge.weight
			if alt > dist[edge.to] {
				dist[edge.to] = alt
				prev[edge.to] = u
			}
			// }
		}
	}

	max := 0
	for _, v := range dist {
		if v > max {
			max = v
		}
	}
	fmt.Println(dist)
}

func maxOfNotKnown(dist []int, notKnown map[int]bool) int {
	maxDist := -1
	maxOfNotKnown := -1
	for vertex := range notKnown {
		if dist[vertex] > maxDist {
			maxDist = dist[vertex]
			maxOfNotKnown = vertex
		}
	}
	return maxOfNotKnown
}
