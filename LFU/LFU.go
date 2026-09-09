package main

import "fmt"

// Node in doubly linked list
type Node struct {
	key   int
	value int
	cnt   int

	prev *Node
	next *Node
}

// Constructor for Node
func NewNode(key, value int) *Node {
	return &Node{
		key:   key,
		value: value,
		cnt:   1,
	}
}

// Doubly Linked List
type List struct {
	size int

	head *Node
	tail *Node
}

// Constructor
func NewList() *List {
	head := &Node{}
	tail := &Node{}

	head.next = tail
	tail.prev = head

	return &List{
		head: head,
		tail: tail,
	}
}

// Add node at front
func (l *List) AddFront(node *Node) {

	temp := l.head.next

	node.next = temp
	temp.prev = node

	l.head.next = node
	node.prev = l.head

	l.size++
}

// Remove node
func (l *List) RemoveNode(node *Node) {

	prev := node.prev
	next := node.next

	prev.next = next
	next.prev = prev

	l.size--
}

// ---------------- LFU ----------------

type LFUCache struct {

	// key -> node
	keyNode map[int]*Node

	// frequency -> doubly linked list
	freqListMap map[int]*List

	maxSizeCache int
	minFreq      int
	curSize      int
}

// Constructor
func NewLFUCache(capacity int) *LFUCache {

	return &LFUCache{
		keyNode:      make(map[int]*Node),
		freqListMap:  make(map[int]*List),
		maxSizeCache: capacity,
	}
}

// Update frequency
func (l *LFUCache) updateFreqListMap(node *Node) {

	delete(l.keyNode, node.key)

	oldList := l.freqListMap[node.cnt]

	oldList.RemoveNode(node)

	if node.cnt == l.minFreq && oldList.size == 0 {
		l.minFreq++
	}

	node.cnt++

	newList, ok := l.freqListMap[node.cnt]
	if !ok {
		newList = NewList()
	}

	newList.AddFront(node)

	l.freqListMap[node.cnt] = newList

	l.keyNode[node.key] = node
}

// Get
func (l *LFUCache) Get(key int) int {

	node, ok := l.keyNode[key]
	if !ok {
		return -1
	}

	value := node.value

	l.updateFreqListMap(node)

	return value
}

// Put
func (l *LFUCache) Put(key, value int) {

	if l.maxSizeCache == 0 {
		return
	}

	// already exists
	if node, ok := l.keyNode[key]; ok {

		node.value = value

		l.updateFreqListMap(node)

		return
	}

	// cache full
	if l.curSize == l.maxSizeCache {

		list := l.freqListMap[l.minFreq]

		lru := list.tail.prev

		delete(l.keyNode, lru.key)

		list.RemoveNode(lru)

		l.curSize--
	}

	l.curSize++

	l.minFreq = 1

	list, ok := l.freqListMap[1]
	if !ok {
		list = NewList()
	}

	node := NewNode(key, value)

	list.AddFront(node)

	l.keyNode[key] = node

	l.freqListMap[1] = list
}

func lfu_example() {

	cache := NewLFUCache(2)

	cache.Put(1, 1)
	cache.Put(2, 2)

	fmt.Print(cache.Get(1), " ")

	cache.Put(3, 3)

	fmt.Print(cache.Get(2), " ")
	fmt.Print(cache.Get(3), " ")

	cache.Put(4, 4)

	fmt.Print(cache.Get(1), " ")
	fmt.Print(cache.Get(3), " ")
	fmt.Print(cache.Get(4), " ")
}
