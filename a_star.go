package main

import "errors"

func createReturn(cameFrom map[*Tree]*Tree, current PriorityQueue) *Queue {
	var queue = CreateQueue(current.tree, 0, nil)
	cur := current.tree

	for {
		cur, ok := cameFrom[cur]
		if !ok {
			break
		}

		queue.insert(CreateQueue(cur, 1, nil))
	}

	return &queue
}

func A_star(start *Tree, goal string, hFunction estimation) (Queue, error) {
	openSet := CreatePriorityQueue(hFunction(start), start, nil)

	var gScore map[*Tree]int
	gScore = make(map[*Tree]int)
	gScore[start] = 0

	var cameFrom map[*Tree]*Tree
	cameFrom = make(map[*Tree]*Tree)

	var current PriorityQueue

	for openSet != nil {
		current, openSet = openSet.pop()
		if current.tree.value == goal {
			return *createReturn(cameFrom, current), nil
		}

		subTree := current.tree
		for _, sons := range subTree.son {
			elem := sons.nextTree
			tentative_gScore := gScore[current.tree] + sons.valueMove

			if tentative_gScore < gScore[elem] {
				cameFrom[elem] = current.tree
				gScore[elem] = tentative_gScore
				openSet = openSet.searchPriorityQueue(tentative_gScore+hFunction(elem), elem)
			}
		}

	}

	return Queue{}, errors.New("Goal not founded")
}

func main() {

}
