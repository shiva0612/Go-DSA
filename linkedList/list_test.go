package main

import (
	"testing"

	"github.com/go-playground/assert"
)

func TestReverse(t *testing.T) {
	head := NewListFromArray([]int{1, 2, 3, 4, 5})
	assert.Equal(t, []int{5, 4, 3, 2, 1}, reverseList(head).arr())
}

func TestMiddle(t *testing.T) {
	head := NewListFromArray([]int{1, 2, 3})
	assert.Equal(t, findMiddleOfList(head).Val, 2)
}

func TestMergeSorted(t *testing.T) {

	type MergeTest struct {
		ip1 []int
		ip2 []int
		op  []int
	}

	mts := []MergeTest{
		{
			[]int{3, 7, 9},
			[]int{8, 9, 10},
			[]int{3, 7, 8, 9, 9, 10},
		},
		{
			[]int{1, 2, 5, 8, 10},
			[]int{3, 7, 10},
			[]int{1, 2, 3, 5, 7, 8, 10, 10},
		},
	}
	for _, mt := range mts {
		p := NewListFromArray(mt.ip1)
		q := NewListFromArray(mt.ip2)
		assert.Equal(t, mergeSortedLists2(p, q).arr(), mt.op)
	}
}

func TestRemoveNode(t *testing.T) {
	head := NewListFromArray([]int{1, 2})
	removeNode(head.find(1))
	assert.Equal(t, head.arr(), []int{2})
}

func TestMathAdd(t *testing.T) {
	p := NewListFromArray([]int{9, 9, 9, 9, 9, 9, 9})
	q := NewListFromArray([]int{9, 9, 9, 9})
	assert.Equal(t, mathAdd(p, q).arr(), []int{8, 9, 9, 9, 0, 0, 0, 1})
}
