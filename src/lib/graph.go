package graph

import (
	"container/list"
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
}

func addEdge(g *Graph, src, dest int, weight float32) {
	g.adj[src].PushBack(&Edge{dest, weight})
	g.adj[dest].PushBack(&Edge{src, weight})
}
