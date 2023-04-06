package main

import (
	"fmt"
	"main/lib"
)

func main() {
	fmt.Println("Hello World!")
	matrix := [][]float32{
		{0, 1, 2, 3},
		{1, 0, 4, 5},
		{2, 4, 0, 6},
		{3, 5, 6, 0},
	}
	g := lib.NewGraph(matrix)
	lib.PrintGraphInfos(g)
}
