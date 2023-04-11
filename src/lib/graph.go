package lib

import (
	"container/list"
	"fmt"
	"strconv"
)

type Edge struct {
	dest   int
	weight float32
}

/*
Graph is a struct that represents a graph
Using an adjacency list
nodes are represented by integers to save the number of nodes
adj is the adjacency list
*/
type Graph struct {
	nodes int
	adj   []*list.List
	names []string
}

func GetName(g Graph) []string {
	return g.names
}

func AddEdge(g *Graph, src, dest int, weight float32) {
	g.adj[src].PushBack(&Edge{dest, weight})
}

func NewGraph(matrix [][]float32) *Graph {
	g := &Graph{len(matrix), make([]*list.List, len(matrix)), make([]string, len(matrix))}

	for i := 0; i < len(matrix); i++ {
		g.adj[i] = list.New()
		g.names[i] = strconv.Itoa(i)
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] != 0 {
				AddEdge(g, i, j, matrix[i][j])
			}
		}
	}

	return g
}

func NewGraphNamed(matrix [][]float32, names []string) *Graph {
	g := &Graph{len(matrix), make([]*list.List, len(matrix)), names}

	for i := 0; i < len(matrix); i++ {
		g.adj[i] = list.New()
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] != 0 {
				AddEdge(g, i, j, matrix[i][j])
			}
		}
	}

	return g
}

func PrintGraphInfos(g *Graph) {
	fmt.Println("Graph Infos:")
	fmt.Println("Number of nodes:", g.nodes)
	fmt.Println("Adjacency List:")
	for i := 0; i < g.nodes; i++ {
		fmt.Print(g.names[i], ": ")
		if g.adj[i].Len() == 0 {
			fmt.Println("No edge")
			continue
		}
		for e := g.adj[i].Front(); e != nil; e = e.Next() {
			fmt.Printf("%s(%f) ", g.names[e.Value.(*Edge).dest], e.Value.(*Edge).weight)
		}
		fmt.Println()
	}
}
