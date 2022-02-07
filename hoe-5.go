package main

import "fmt"

type myType int

func main() {

	var x myType
	var y int

	x = 42
	y = 84
	x = myType(y)

	fmt.Printf("value of x = %v and type of x = %T\n", x, x)
	fmt.Printf("value of y = %v and type of y = %T", y, y)

}
