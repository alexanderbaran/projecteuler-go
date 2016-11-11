package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"time"
// )

// // type edge struct {
// // 	weight int
// // 	to     int
// // }

// func main() {
// 	defer timeTrack(time.Now())
// 	ints := []int{}
// 	f, err := os.Open("018.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	scanner := bufio.NewScanner(f)
// 	for scanner.Scan() {
// 		for _, v := range strings.Split(scanner.Text(), " ") {
// 			d, _ := strconv.Atoi(v)
// 			ints = append(ints, d)
// 		}
// 	}
// 	if err := scanner.Err(); err != nil {
// 		fmt.Println(err)
// 	}

// 	// adj := make([][]edge, len(ints))
// 	// test := 0
// 	// incr := 1
// 	// for i := 0; i <= len(adj); {
// 	// 	if i >= test && i < test+incr {
// 	// 		left := edge{to: i + incr, weight: ints[i+incr]}
// 	// 		right := edge{to: i + incr + 1, weight: ints[i+incr+1]}
// 	// 		edges := []edge{left, right}
// 	// 		adj[i] = edges
// 	// 		i++
// 	// 	} else {
// 	// 		test += incr
// 	// 		incr++
// 	// 	}
// 	// 	if test+incr == len(adj) {
// 	// 		break
// 	// 	}
// 	// }
// 	// spew.Dump(adj)

// 	// // Cannot use Dijkstra for longest path
// 	// notKnown := map[int]bool{}
// 	// dist := make([]int, len(ints))
// 	// // prev := make([]int, len(ints))
// 	// for i := 0; i < len(ints); i++ {
// 	// 	dist[i] = -1
// 	// 	prev[i] = -1 // Undefined
// 	// 	notKnown[i] = false
// 	// }
// 	// dist[0] = ints[0]

// 	// for len(notKnown) > 0 {
// 	// 	u := maxOfNotKnown(dist, notKnown)
// 	// 	delete(notKnown, u)

// 	// 	for _, edge := range adj[u] {
// 	// 		if _, ok := notKnown[edge.to]; !ok {
// 	// 			continue
// 	// 		}
// 	// 		alt := dist[u] + edge.weight
// 	// 		if alt > dist[edge.to] {
// 	// 			dist[edge.to] = alt
// 	// 			prev[edge.to] = u
// 	// 		}
// 	// 	}
// 	// }

// 	// dist := make([]int, len(adj))
// 	// dist[0] = ints[0]
// 	// for i := 0; i < len(adj); i++ {
// 	// 	for _, edge := range adj[i] {
// 	// 		alt := dist[i] + edge.weight
// 	// 		if alt > dist[edge.to] {
// 	// 			dist[edge.to] = alt
// 	// 		}
// 	// 	}
// 	// }

// 	// max := 0
// 	// for _, v := range dist {
// 	// 	if v > max {
// 	// 		max = v
// 	// 	}
// 	// }

// 	dist := make([]int, len(ints))
// 	dist[0] = ints[0]
// 	test := 0
// 	incr := 1
// 	max := 0
// 	for i := 0; i < len(ints); {
// 		if i >= test && i < test+incr {
// 			for k := 0; k <= 1; k++ {
// 				alt := dist[i] + ints[i+incr+k]
// 				if alt > dist[i+incr+k] {
// 					dist[i+incr+k] = alt
// 					if alt > max {
// 						max = alt
// 					}
// 				}
// 			}
// 			i++
// 		} else {
// 			test += incr
// 			incr++
// 		}
// 		if test+incr == len(ints) {
// 			break
// 		}
// 	}
// 	fmt.Println(max)
// }

// // func maxOfNotKnown(dist []int, notKnown map[int]bool) int {
// // 	maxDist := -1
// // 	maxOfNotKnown := -1
// // 	for vertex := range notKnown {
// // 		if dist[vertex] > maxDist {
// // 			maxDist = dist[vertex]
// // 			maxOfNotKnown = vertex
// // 		}
// // 	}
// // 	return maxOfNotKnown
// // }
