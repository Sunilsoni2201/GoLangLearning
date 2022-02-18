package main

import (
	"fmt"
)

func main() {
	an_struct := struct {
		title  string
		author string
		year   int
	}{
		title:  "Truth",
		author: "Unknown",
		year:   0,
	}
	fmt.Println(an_struct)
}
