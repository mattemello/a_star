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

type h func(Tree) int

type PriorityQueue struct {
	value int
	next  *PriorityQueue
}

func CreatePriorityQueue(v int, n *PriorityQueue) *PriorityQueue {
	return &PriorityQueue{
		value: v,
		next:  n,
	}
}

func (x *PriorityQueue) insert(v int) *PriorityQueue {

	if x.value > v {
		var try = CreatePriorityQueue(v, x)
		x = try
		return x
	}

	start := x

	for x.next != nil {
		if x.next.value >= v {
			var try = CreatePriorityQueue(v, x.next)
			x.next = try
			return start
		}

		x = x.next
	}

	if x.value > v {
		var try = CreatePriorityQueue(v, x)
		x.next = try
		return start
	} else {
		var try = CreatePriorityQueue(v, nil)
		x.next = try
		return start
	}

}
