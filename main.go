package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Welcome to calculator!")
	for {
		option, err := getOption([]Option{
			{"C", "Calculate a recipe chain"},
			{"I", "Browse Items"},
			{"R", "Browse Recipes"},
			{"X", "Exit"},
		})
		if err != nil {
			log.Fatal(err)
		}
		switch option {
		case "C":
			calculateRecipeChain()
		case "I":
			browseItems()
		case "R":
			calculateRecipeChain()
		case "X":
			return
		}
	}
}
