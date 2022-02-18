package main

import (
	"fmt"
)

type Person struct {
	first       string
	last        string
	favIceCream []string
}

func main() {

	p1 := Person{
		first:       "AGyani",
		last:        "Buddha",
		favIceCream: []string{"Vanilla", "butterScotch"},
	}

	p2 := Person{
		first:       "James",
		last:        "Bond",
		favIceCream: []string{"Strawberry", "Chocolate"},
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.favIceCream {
		println(i, v)
	}

	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.favIceCream {
		println(i, v)
	}

}
