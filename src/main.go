package main

import (
	"fmt"
	"main/lib"
)

func main() {
	g, x, y := lib.FileNameParse()
	lib.PrintGraphInfos(&g)
	fmt.Println("Do you want to use a* or ucs?")
	fmt.Println("1. A*")
	fmt.Println("2. UCS")
	algo := lib.RangedInput(1, 2)
	fmt.Println("Below are the names of the nodes:")
	for i := 0; i < len(lib.GetName(g)); i++ {
		fmt.Println(i, lib.GetName(g)[i])
	}
	fmt.Println("Please enter the start node:")
	start := lib.RangedInput(0, len(lib.GetName(g))-1)
	fmt.Println("Please enter the end node:")
	end := lib.RangedInput(0, len(lib.GetName(g))-1)
	var res *lib.Item
	if algo == 1 {
		res = lib.Astar(g, x, y, start, end)
	} else {
		res = lib.UCS(g, start, end)
	}
	lib.PrintPath(g, *res)
}
