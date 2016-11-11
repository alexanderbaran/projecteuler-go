package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	defer timeTrack(time.Now())
	ints := []int{}
	f, err := os.Open("067.txt")
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
	dist := make([]int, len(ints))
	dist[0] = ints[0]
	test := 0
	incr := 1
	max := 0
	for i := 0; i < len(ints); {
		if i >= test && i < test+incr {
			for k := 0; k <= 1; k++ {
				alt := dist[i] + ints[i+incr+k]
				if alt > dist[i+incr+k] {
					dist[i+incr+k] = alt
					if alt > max {
						max = alt
					}
				}
			}
			i++
		} else {
			test += incr
			incr++
		}
		if test+incr == len(ints) {
			break
		}
	}
	fmt.Println(max)
}
