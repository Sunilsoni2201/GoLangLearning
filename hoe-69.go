package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 126; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%v\t\t\t%x\t\t\t%#U\n", i, i, i)
		}
		fmt.Println("")
	}
}
