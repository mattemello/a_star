package main

type Tree struct {
	value     string
	neighbors *[]Edge
	next      *[]Tree
}

type Edge struct {
	node  *Tree
	node2 *Tree
	value int
}

type h func(*Tree) int

type Queue struct {
	node  *Tree
	value int
	next  *Queue
}

func CreateQueue(st *Tree, v int, nx *Queue) Queue {
	return Queue{
		node:  st,
		value: v,
		next:  nx,
	}
}

func (q Queue) insert(nx Queue) {
	st := q

	for st.next != nil {
		st = *st.next
	}

	st.next = &nx
}

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

func (x *PriorityQueue) searchPriorityQueue(v int, tree *Tree) *PriorityQueue {
	start := x

	if x.tree.value == tree.value {
		x = x.next
		return x.insert(v, tree)
	}

	for x != nil {
		if x.next == nil {
			break
		}

		if x.next.tree.value == tree.value {
			x.next = x.next.next
			break
		}
		x = x.next
	}

	return start.insert(v, tree)
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
