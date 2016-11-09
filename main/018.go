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
	ints = ints[1:len(ints)]

	test := 0
	incr := 1
	for i := 0; i <= len(adj); i++ {
		if i > test {
			test += incr
			incr++
		}
		e1 := edge{to: i + 1}
		adj[i]
	}
}
