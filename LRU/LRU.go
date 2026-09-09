package main

import "fmt"

type Node struct {
	key  int
	val  int
	prev *Node
	next *Node
}

type LRUCache struct {
	cap  int
	head *Node
	tail *Node
	m    map[int]*Node
}

func NewLRUCache(capacity int) *LRUCache {
	head := &Node{key: -1, val: -1}
	tail := &Node{key: -1, val: -1}

	head.next = tail
	tail.prev = head

	return &LRUCache{
		cap:  capacity,
		head: head,
		tail: tail,
		m:    make(map[int]*Node),
	}
}

// Add node right after head (Most Recently Used)
func (l *LRUCache) addNode(newNode *Node) {
	temp := l.head.next

	newNode.next = temp
	temp.prev = newNode

	l.head.next = newNode
	newNode.prev = l.head
}

// Delete node from linked list
func (l *LRUCache) deleteNode(delNode *Node) {
	p := delNode.prev
	n := delNode.next

	p.next = n
	n.prev = p
}

// Get value from cache
func (l *LRUCache) Get(key int) int {

	node, ok := l.m[key]
	if !ok {
		return -1
	}

	value := node.val

	// Move node to front
	l.deleteNode(node)
	l.addNode(node)

	l.m[key] = l.head.next

	return value
}

// Put key-value into cache
func (l *LRUCache) Put(key, value int) {

	// Key already exists
	if node, ok := l.m[key]; ok {
		l.deleteNode(node)
		delete(l.m, key)
		//we are deleting bcz we need to add right after head as its LRU
	}

	// Cache full
	if len(l.m) == l.cap {

		lru := l.tail.prev

		delete(l.m, lru.key)
		l.deleteNode(lru)
	}

	newNode := &Node{
		key: key,
		val: value,
	}

	l.addNode(newNode)
	l.m[key] = newNode
}

func lru_example() {

	cache := NewLRUCache(2)

	cache.Put(1, 1)
	cache.Put(2, 2)

	fmt.Println(cache.Get(1)) // 1

	cache.Put(3, 3)

	fmt.Println(cache.Get(2)) // -1

	cache.Put(4, 4)

	fmt.Println(cache.Get(1)) // -1
	fmt.Println(cache.Get(3)) // 3
	fmt.Println(cache.Get(4)) // 4
}
