package main

import (
	"fmt"
	"testing"
)

func TestPiorityQueue(t *testing.T) {
	var prir = CreatePriorityQueue(4, &Tree{value: "str"}, nil)
	prir = prir.insert(5, &Tree{value: "a"})
	prir = prir.insert(3, &Tree{value: "b"})
	prir = prir.insert(4, &Tree{value: "c"})

	if prir == nil {
		t.Errorf("it sucks")
	}

	var Wanted = []int{3, 4, 4, 5}

	for i, _ := range Wanted {
		if Wanted[i] != prir.value {
			t.Errorf("It's not working wanted %d, recived %d", Wanted[i], prir.value)
		}

		prir = prir.next
	}

	var pque = CreatePriorityQueue(4, &Tree{value: "str"}, nil)
	pque = pque.insert(5, &Tree{value: "a"})
	pque = pque.insert(3, &Tree{value: "b"})
	pque = pque.insert(4, &Tree{value: "c"})

	pque = pque.searchPriorityQueue(2, &Tree{value: "b"})
	var WantedSearch = []string{"b", "c", "str", "a"}

	for i, _ := range WantedSearch {
		if WantedSearch[i] != pque.tree.value {
			t.Errorf("It's not working wanted %s, recived %s", WantedSearch[i], pque.tree.value)
		}

		fmt.Println(pque.tree.value, pque)

		pque = pque.next
	}

	fmt.Println("All good!")
}
