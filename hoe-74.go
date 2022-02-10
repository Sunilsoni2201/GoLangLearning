package main

import (
	"fmt"
)

func main() {
	i := 0
	if i != 0 {
		fmt.Println("Print it")
	} else if i < 1 {
		fmt.Println("Print it too")
	} else if i > 0 {
		fmt.Println("Don't print it")
	} else {
		fmt.Println("")
	}

}
