package main

import (
	"fmt"
)

type person struct {
	first       string
	last        string
	favIceCream []string
}

func main() {

	p1 := person{
		first:       "AGyani",
		last:        "Buddha",
		favIceCream: []string{"Vanilla", "butterScotch"},
	}

	p2 := person{
		first:       "James",
		last:        "Bond",
		favIceCream: []string{"Strawberry", "Chocolate"},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println(m)

	for k, v := range m {
		fmt.Printf("%v\t%v\t%v\n", k, v.first, v.last)
		for i, v := range v.favIceCream {
			fmt.Printf("%v\t%v\n", i, v)
		}

	}

	/*
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
	*/

}
