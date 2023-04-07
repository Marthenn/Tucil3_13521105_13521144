package lib

import (
	"container/heap"
	"container/list"
)

func euclideanCalculator(x, y []float32, end int) []float32 {
	euclid := make([]float32, len(x))
	for i := 0; i < len(x); i++ {
		euclid[i] = euclideanDistance(x[i], y[i], x[end], y[end])
	}
	return euclid
}

func Astar(g Graph, x, y []float32, start, end int) *Item {
	// init variables
	euclid := euclideanCalculator(x, y, end)
	expandedWeight := euclid
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{start, euclid[start], list.List{}, 0})
	visited := list.New()
	// start algorithm
	for pq.Len() > 0 {
		currentItem := heap.Pop(&pq).(*Item)
		currentNodeNumber := currentItem.Value
		currentNodeCost := currentItem.Priority
		currentPassed := list.New()
		currentPassed.PushBackList(&currentItem.PassedNode)
		currentPassed.PushBack(currentNodeNumber)

		// update expandedWeight
		for e := g.adj[currentNodeNumber].Front(); e != nil; e = e.Next() {
			if expandedWeight[e.Value.(*Edge).dest] > currentNodeCost+e.Value.(*Edge).weight+euclid[e.Value.(*Edge).dest] {
				expandedWeight[e.Value.(*Edge).dest] = currentNodeCost + e.Value.(*Edge).weight + euclid[e.Value.(*Edge).dest]
			}
		}

		if currentNodeNumber == end {
			var dist float32 = 0
			for e := currentPassed.Front(); e != nil && e.Next() != nil; e = e.Next() {
				for f := g.adj[e.Value.(int)].Front(); f != nil; f = f.Next() {
					if f.Value.(*Edge).dest == e.Next().Value.(int) {
						dist += f.Value.(*Edge).weight
					}
				}
			}
			return &Item{
				Value:      currentNodeNumber,
				Priority:   float32(dist),
				PassedNode: *currentPassed,
				Index:      0,
			}
		} else {
			if !itemInList(*currentItem, *visited) {
				visited.PushBack(currentItem)
				for e := g.adj[currentNodeNumber].Front(); e != nil; e = e.Next() {
					tmp := &Item{e.Value.(*Edge).dest, currentNodeCost + e.Value.(*Edge).weight, *currentPassed, len(pq)}
					heap.Push(&pq, tmp)
					pq.Update(tmp, tmp.Value, currentNodeCost+e.Value.(*Edge).weight+expandedWeight[tmp.Value])
				}
			}
		}
	}
	return &Item{
		Value:      -1,
		Priority:   -1,
		PassedNode: list.List{},
		Index:      -1,
	}
}
