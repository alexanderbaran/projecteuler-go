package main

import (
	"bytes"
	"strconv"
	"time"
)

func main() {
	defer timeTrack(time.Now())

	// i := 12345
	// o := splitInt(i)
	// s := splitInt(i)

	for i := 50000; i <= 99999; i++ {

		o := splitInt(i)
		s := splitInt(i)

		for j := 0; j < len(o); j++ {
			copy(s, o)
			s[j] = r
			// fmt.Println(joinInts(s))
			for i, _ := range s {
				if i != j {
					s[i] = r
				} else {
					s[i] = o[i]
				}
			}
			// fmt.Println(joinInts(s))
		}

		for j := 0; j < len(o); j++ {
			for k := j + 1; k < len(o); k++ {
				copy(s, o)
				s[j] = r
				s[k] = r
				// fmt.Println(joinInts(s))
				for i, _ := range s {
					if i != j {
						s[i] = r
					} else {
						s[i] = o[i]
					}
				}
				// fmt.Println(joinInts(s))
			}
		}
	}

	// for i := 50000; i <= 99999; i++ {
	// 	for r := 0; r <= 9; r++ {
	// 	}
	// }
}

func splitInt(n int) []int {
	s := make([]int, digitLength(n))
	for i := len(s) - 1; n > 0; i-- {
		s[i] = n % 10
		n = n / 10
	}
	return s
}

func joinInts(s []int) int {
	var buffer bytes.Buffer
	for _, v := range s {
		buffer.WriteString(strconv.Itoa(v))
	}
	i, _ := strconv.Atoi(buffer.String())
	return i
}

// *x
// x*

// *xx
// x*x
// xx*

// **x
// *x*
// x**

// *xxx
// x*xx
// xx*x
// xxx*

// **xx
// *x*x
// *xx*
// x**x
// x*x*
// xx**

// *xxxx
// x*xxx
// xx*xx
// xxx*x
// xxxx*

// **xxx
// *x*xx
// *xx*x
// *xxx*
// x**xx
// x*x*x
// x*xx*
// xx**x
// xx*x*
// xxx**
