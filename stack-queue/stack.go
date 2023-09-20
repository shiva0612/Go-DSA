package main

import (
	"fmt"
)

type Stack []interface{}

func (s *Stack) Push(item interface{}) {
	*s = append(*s, item)
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	lastIndex := len(*s) - 1
	item := (*s)[lastIndex]
	*s = (*s)[:lastIndex]
	return item
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func simple_stack() {
	stack := make(Stack, 0)

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	fmt.Println("Stack:", stack)

	poppedItem := stack.Pop()
	fmt.Println("Popped:", poppedItem)
	fmt.Println("Stack after pop:", stack)

	fmt.Println("Is Stack empty?", stack.IsEmpty())
}
