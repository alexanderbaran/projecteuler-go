package main

import (
	"fmt"
	"strings"
)

func main() {
	x := "123"
	s := []string{}

	// for i := 0; i < len(x); i++ {
	// 	s = strings.Split(x, "")
	// 	s[i] = "*"
	// 	fmt.Println(s)
	// }

	for i := 0; i < len(x); i++ {
		for j := i + 1; j < len(x); j++ {
			s = strings.Split(x, "")
			s[i] = "*"
			s[j] = "*"
			fmt.Println(s)
		}
	}
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
