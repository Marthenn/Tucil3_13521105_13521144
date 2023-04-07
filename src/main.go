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
	//var matrix = [][]float32{
	//	{0, 5, 3, 0, 0},
	//	{5, 0, 0, 2, 0},
	//	{3, 0, 0, 8, 10},
	//	{0, 2, 8, 0, 2},
	//	{0, 0, 10, 2, 0},
	//}
	g := lib.NewGraph(matrix)
	lib.PrintGraphInfos(g)
	res := lib.UCS(*g, 0, 3)
	fmt.Println("pr", res.Priority)
	for e := res.PassedNode.Front(); e != nil; e = e.Next() {
		fmt.Println("v", e.Value)
	}
	println()
}
