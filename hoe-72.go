package main

import "fmt"

func main() {
	v := 100
	i := 10
	for i <= v {
		fmt.Printf("%v\n", i%4)
		i++
	}

}
