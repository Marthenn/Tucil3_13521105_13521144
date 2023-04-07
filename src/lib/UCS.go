package lib

import (
	"container/heap"
	"container/list"
	"fmt"
)

func UCS(g Graph, startNode int, goalNode int) *Item {
	nodeQueue := make(PriorityQueue, 0)
	heap.Init(&nodeQueue)
	heap.Push(&nodeQueue, &Item{startNode, 0, list.List{}, 0})
	visited := list.New()
	fmt.Println("nq", nodeQueue.Len())
	for nodeQueue.Len() > 0 {
		currentItem := heap.Pop(&nodeQueue).(*Item)
		currentNodeNumber := currentItem.Value
		currentNodeCost := currentItem.Priority
		currentPassed := list.New()
		currentPassed.PushBackList(&currentItem.PassedNode)
		currentPassed.PushBack(currentNodeNumber)
		//fmt.Println("currnod", currentNodeNumber, currentNodeCost)
		//for e := currentPassed.Front(); e != nil && e.Value != nil; e = e.Next() {
		//	fmt.Println("v", e.Value)
		//}
		if currentNodeNumber == goalNode {
			return &Item{
				Value:      currentNodeNumber,
				Priority:   currentNodeCost,
				PassedNode: *currentPassed,
				Index:      0,
			}
		} else {
			if !itemInList(*currentItem, *visited) {
				visited.PushBack(currentItem)
				for e := g.adj[currentNodeNumber].Front(); e != nil; e = e.Next() {
					//fmt.Println("1", e.Value.(*Edge).dest)
					tmp := &Item{e.Value.(*Edge).dest, currentNodeCost + e.Value.(*Edge).weight, *currentPassed, len(nodeQueue)}
					heap.Push(&nodeQueue, tmp)
					nodeQueue.Update(tmp, tmp.Value, currentNodeCost+e.Value.(*Edge).weight)
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
	//PrintGraphInfos(g)
}
