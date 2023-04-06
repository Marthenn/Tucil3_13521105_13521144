package graph

import (
	"container/list"
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

func addEdge(g *Graph, src, dest int, weight float32) {
	g.adj[src].PushBack(&Edge{dest, weight})
	g.adj[dest].PushBack(&Edge{src, weight})
}

func newGraph(matrix [][]float32) *Graph {
	g := &Graph{len(matrix), make([]*list.List, len(matrix)), make([]string, len(matrix))}

	for i := 0; i < len(matrix); i++ {
		g.adj[i] = list.New()
		g.names[i] = strconv.Itoa(i)
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] != 0 {
				addEdge(g, i, j, matrix[i][j])
			}
		}
	}

	return g
}

func newGraphNamed(matrix [][]float32, names []string) *Graph {
	g := &Graph{len(matrix), make([]*list.List, len(matrix)), names}

	for i := 0; i < len(matrix); i++ {
		g.adj[i] = list.New()
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] != 0 {
				addEdge(g, i, j, matrix[i][j])
			}
		}
	}

	return g
}
