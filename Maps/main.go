package main

import "fmt"

func main() {

	// first way to initialize maps
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#ef234",
	}

	// 2nd way
	// var colors map[string]string

	// 3rd way
	// colors := make(map[string]string)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, " is ", hex)
	}
}
