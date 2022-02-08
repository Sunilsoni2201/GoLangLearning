package main

import "fmt"

const NUMBER = 100

func main() {
	fmt.Println("Testing")

	for i := 0; i <= NUMBER; i++ {
		fmt.Printf("%v\t", i)
		fmt.Printf("%b\t", i)
		fmt.Printf("%x\t", i)
		fmt.Println("")
	}
}
