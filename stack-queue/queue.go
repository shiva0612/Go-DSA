package main

import (
	"fmt"
)

type Queue []interface{}

func (q *Queue) Enqueue(item interface{}) {
	*q = append(*q, item)
}

func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	item := (*q)[0]
	*q = (*q)[1:]
	return item
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func simple_queue() {
	queue := make(Queue, 0)

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	fmt.Println("Queue:", queue)

	dequeuedItem := queue.Dequeue()
	fmt.Println("Dequeued:", dequeuedItem)
	fmt.Println("Queue after dequeue:", queue)

	fmt.Println("Is Queue empty?", queue.IsEmpty())
}
