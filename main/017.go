package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	defer timeTrack(time.Now())
// 	a := map[int]string{
// 		1:  "one",
// 		2:  "two",
// 		3:  "three",
// 		4:  "four",
// 		5:  "five",
// 		6:  "six",
// 		7:  "seven",
// 		8:  "eight",
// 		9:  "nine",
// 		10: "ten",
// 		11: "eleven",
// 		12: "twelve",
// 		13: "thirteen",
// 		14: "fourteen",
// 		15: "fifteen",
// 		16: "sixteen",
// 		17: "seventeen",
// 		18: "eighteen",
// 		19: "nineteen",
// 		20: "twenty",
// 	}
// 	c := map[int]string{
// 		20: "twenty",
// 		30: "thirty",
// 		40: "forty",
// 		50: "fifty",
// 		60: "sixty",
// 		70: "seventy",
// 		80: "eighty",
// 		90: "ninety",
// 	}
// 	h := "hundred"
// 	sum := 0
// 	for i := 1; i <= 999; i++ {
// 		if i <= 20 {
// 			// fmt.Println(a[i])
// 			sum += len(a[i])
// 		} else if i <= 99 {
// 			// fmt.Println(c[(i/10)*10], a[i%10])
// 			sum += len(c[(i/10)*10])
// 			sum += len(a[i%10])
// 		} else {
// 			if i%100 == 0 {
// 				// fmt.Println(a[i/100], h)
// 				sum += len(a[i/100])
// 				sum += len(h)
// 			} else {
// 				d := i % 100
// 				if d <= 20 {
// 					// fmt.Println(a[i/100], h, "and", a[i%100])
// 					sum += len(a[i/100])
// 					sum += len(h)
// 					sum += len("and")
// 					sum += len(a[i%100])
// 				} else {
// 					// fmt.Println(a[i/100], h, "and", c[(d/10)*10], a[d%10])
// 					sum += len(a[i/100])
// 					sum += len(h)
// 					sum += len("and")
// 					sum += len(c[(d/10)*10])
// 					sum += len(a[d%10])
// 				}
// 			}
// 		}
// 	}
// 	// fmt.Println("one thousand")
// 	sum += len("one")
// 	sum += len("thousand")
// 	fmt.Println(sum)
// }
