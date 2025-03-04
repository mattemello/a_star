package main

import (
	"fmt"
	"testing"
)

func TestPiorityQueue(t *testing.T) {
	var prir = CreatePriorityQueue(4, nil)
	prir.insert(5)
	prir.insert(3)
	prir.insert(4)

	if prir == nil {
		t.Errorf("it sucks")
	}

	fmt.Println(prir.value)

	var Wanted = []int{3, 4, 4, 5}

	for i, _ := range Wanted {
		if Wanted[i] != prir.value {
			t.Errorf("It's not working wanted %d, recived %d", Wanted[i], prir.value)
		}

		prir = prir.next
	}

	fmt.Println("All good!")
}
