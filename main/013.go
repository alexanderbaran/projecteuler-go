package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"time"
// )

// func main() {
// 	defer timeTrack(time.Now())
// 	file, err := os.Open("013.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer file.Close()
// 	sum := "0"
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		sum = add(sum, scanner.Text())
// 	}
// 	if err := scanner.Err(); err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(sum[:10])
// }
