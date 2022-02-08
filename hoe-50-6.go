package main

import "fmt"

const (
	y1 = 2018 + iota
	y2 = 2018 + iota
	y3 = 2018 + iota
	y4 = 2018 + iota
)

func main() {
	fmt.Println("")
	fmt.Println(y1)
	fmt.Println(y2)
	fmt.Println(y3)
	fmt.Println(y4)
}
