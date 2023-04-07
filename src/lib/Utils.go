package lib

import "container/list"

func itemInList(item Item, list list.List) bool {
	for e := list.Front(); e != nil; e = e.Next() {
		if item.Value == e.Value {
			return true
		}
	}
	return false
}
