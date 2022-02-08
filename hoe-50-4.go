package main

import "fmt"

func main() {
	a := 42
	fmt.Printf("%v\t %b\t %x\t %T\n", a, a, a, a)

	b := a << 1
	fmt.Printf("%v\t %b\t %x\t %T\n", b, b, b, b)
}
