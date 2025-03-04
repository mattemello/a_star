package main

type Tree struct {
	value     int
	neighbors *[]Edge
	next      *[]Tree
}

type Edge struct {
	node  *Tree
	node2 *Tree
	value int
}

type h func(*Tree) int

type PriorityQueue struct {
	value int
	tree  *Tree
	next  *PriorityQueue
}

func CreatePriorityQueue(v int, tree *Tree, n *PriorityQueue) *PriorityQueue {
	return &PriorityQueue{
		value: v,
		tree:  tree,
		next:  n,
	}
}

func (x *PriorityQueue) pop() (PriorityQueue, *PriorityQueue) {
	return *x, x.next
}

func (x *PriorityQueue) insert(v int, tree *Tree) *PriorityQueue {

	if x.value > v {
		var try = CreatePriorityQueue(v, tree, x)
		x = try
		return x
	}

	start := x

	for x.next != nil {
		if x.next.value >= v {
			var try = CreatePriorityQueue(v, tree, x.next)
			x.next = try
			return start
		}

		x = x.next
	}

	if x.value > v {
		var try = CreatePriorityQueue(v, tree, x)
		x.next = try
		return start
	} else {
		var try = CreatePriorityQueue(v, tree, nil)
		x.next = try
		return start
	}

}
