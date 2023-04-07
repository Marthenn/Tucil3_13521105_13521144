package lib

func astar(g Graph, euclid []float32, start, end int) (path []int, dist float32) {
	expanded := make([]float32, len(g.adj))
	for i := 0; i < len(g.adj); i++ {
		expanded[i] = euclid[i]
	}
	visited := make([]bool, len(g.adj))
	visited[start] = true
	previous := make([]int, len(g.adj))
	previous[start] = -1
	dist = 0
	for i := 0; i < len(g.adj); i++ {
		min := -1
		for j := 0; j < len(g.adj); j++ {
			if !visited[j] && (min == -1 || expanded[j] < expanded[min]) {
				min = j
			}
		}
		if min == -1 {
			break
		}
		visited[min] = true
		for e := g.adj[min].Front(); e != nil; e = e.Next() {
			edge := e.Value.(*Edge)
			if !visited[edge.dest] && (expanded[min]+edge.weight < expanded[edge.dest]) {
				expanded[edge.dest] = expanded[min] + edge.weight
				previous[edge.dest] = min
			}
		}
	}
	path = make([]int, 0)
	for i := end; i != -1; i = previous[i] {
		path = append(path, i)
	}
	dist = expanded[end]
	return
}
