package main

import "fmt"

const (
	A = 10
	B = "Agyani"
)

const (
	C int    = 20
	D string = "Buddha"
)

func main() {
	fmt.Println("")
	fmt.Printf("%v\t %T\n", A, A)
	fmt.Printf("%v\t %T\n", B, B)
	fmt.Printf("%v\t %T\n", C, C)
	fmt.Printf("%v\t %T\n", D, D)
}
