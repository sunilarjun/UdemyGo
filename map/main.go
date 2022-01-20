package main

import "fmt"

func main() {
	//option 1 create map
	//var colors map[string]string

	//option 2
	//colors := make(map[int]string)

	//option 3
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"white": "#ffffff",
	}

	//colors[10] = "#ffffff"

	//delete(colors, 10)

	//fmt.Println(colors)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
