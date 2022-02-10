package main

import (
	"fmt"
)

func main() {
	favSport := "Match"

	switch favSport {
	case "Football":
		fmt.Println("Fav sport is Football")
	case "Cricket":
		fmt.Println("Fav sport is Cricket")
	case "Match":
		fmt.Println("Fav sport is Match")
	default:
		fmt.Println("No Fav sport")
	}
}
