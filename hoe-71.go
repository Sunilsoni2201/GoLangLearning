package main

import "fmt"

func main() {
	v := 2022
	i := 1988
	for {
		if i > v {
			break
		}
		fmt.Printf("%v\n", i)
		i++
	}

}
