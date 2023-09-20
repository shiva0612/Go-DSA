package main

import (
	"fmt"
	"math"
)

type Node struct {
	Left     *Node
	Right    *Node
	Val      int
	ColIndex int
	RowIndex int
}

// ----------------------------------------------
type Stack struct {
	arr []*Node
}

func (s *Stack) Push(n *Node) {
	s.arr = append(s.arr, n)
}
func (s *Stack) Top() *Node {
	if s.Len() > 0 {
		return s.arr[s.Len()-1]
	}
	return nil
}
func (s *Stack) Pop() *Node {
	l := s.Len()
	if l == 0 {
		return nil
	}
	n := s.arr[l-1]
	s.arr = s.arr[:l-1]
	return n
}
func (s *Stack) Len() int {
	return len(s.arr)
}

// ----------------------------------------------
type Queue struct {
	arr []*Node
}

func (q *Queue) Push(node *Node) {
	q.arr = append([]*Node{node}, q.arr...)
}
func (q *Queue) Pop() *Node {
	if q.Len() > 0 {
		top := q.arr[0]
		q.arr = q.arr[1:]
		return top
	}
	return nil
}
func (q *Queue) Top() *Node {
	if q.Len() > 0 {
		return q.arr[0]
	}
	return nil
}
func (q *Queue) Len() int {
	return len(q.arr)
}

// ----------------------------------------------
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ----------------------------------------------

func main() {

}

func PreR(node *Node) {
	//root left right
	if node == nil {
		return
	}
	fmt.Println(node.Val)
	PreR(node.Left)
	PreR(node.Right)

}

func InR(node *Node) {
	//left root right
	if node == nil {
		return
	}
	InR(node.Left)
	fmt.Println(node.Val)
	InR(node.Right)

}

func PostR(node *Node) {
	//root left right
	if node == nil {
		return
	}
	PostR(node.Right)
	PostR(node.Left)
	fmt.Println(node.Val)

}

// ----------------------------------------------
func PreI(node *Node) []int {
	ans := []int{}
	s := new(Stack)
	if node == nil {
		return ans
	}
	s.Push(node)
	for s.Len() > 0 {
		node := s.Pop()
		ans = append(ans, node.Val)
		if node.Right != nil {
			s.Push(node.Right)
		}
		if node.Left != nil {
			s.Push(node.Left)
		}
	}
	return ans
}

func InI(node *Node) []int {
	s := new(Stack)
	ans := []int{}
	for {
		if node != nil {
			s.Push(node)
			node = node.Left
		} else {
			if s.Len() == 0 {
				break
			}
			top := s.Pop()
			ans = append(ans, top.Val)
			node = top.Right
		}
	}
	return ans
}

func PostI(node *Node) {
	s1 := new(Stack)
	s2 := new(Stack) //ans

	s1.Push(node)
	for s1.Len() > 0 {
		top := s1.Pop()
		if top.Left != nil {
			s1.Push(top.Left)
		}
		if top.Right != nil {
			s1.Push(top.Right)
		}
		s2.Push(top)
	}
}

// ----------------------------------------------
func Level(node *Node) [][]int {
	ans := [][]int{}
	if node == nil {
		return ans
	}
	q := new(Queue)
	q.Push(node)
	for q.Len() > 0 {
		l := q.Len()
		temp := []int{}
		for i := 0; i < l; i++ {
			top := q.Pop()
			temp = append(temp, top.Val)
			if top.Left != nil {
				q.Push(top.Left)
			}
			if top.Right != nil {
				q.Push(top.Right)
			}
		}
		ans = append(ans, temp)
	}
	return ans
}

func Height(node *Node) int {
	if node == nil {
		return 0
	}
	left := Height(node.Left)
	right := Height(node.Right)
	return 1 + max(left, right)
}

func CheckIfBalanced1(node *Node) bool {
	if node == nil {
		return true
	}

	left := Height(node.Left)
	right := Height(node.Right)
	if math.Abs(float64(left)-float64(right)) > 1 {
		return false
	}

	leftBool := CheckIfBalanced1(node.Left)
	rightBool := CheckIfBalanced1(node.Right)
	if leftBool || rightBool {
		return false
	}
	return true
}

// CheckIfBalanced2 == -1
func CheckIfBalanced2(node *Node) int {
	if node == nil {
		return 0
	}

	left := CheckIfBalanced2(node.Left)
	if left == -1 {
		return -1
	}
	right := CheckIfBalanced2(node.Right)
	if right == -1 {
		return -1
	}

	if math.Abs(float64(left)-float64(right)) > 1 {
		return -1
	}
	return 1 + max(left, right)
}

func Diameter(node *Node, d *int) int {
	if node == nil {
		return 0
	}
	left := Diameter(node.Left, d)
	right := Diameter(node.Right, d)

	*d = max(*d, left+right)
	return 1 + max(left, right)
}

func MaxPathSum(node *Node, sum *int) int {
	if node == nil {
		return 0
	}
	left := MaxPathSum(node.Left, sum)
	right := MaxPathSum(node.Right, sum)

	*sum = max(*sum, left+right+*sum)
	return node.Val + max(left, right)
}

func IsSame(p, q *Node) bool {
	if p == nil || q == nil {
		return p == q
	}
	return p.Val == q.Val && IsSame(p.Left, q.Left) && IsSame(p.Right, q.Right)
}

// -----------------------------------------------
func topView(node *Node) {
	m := map[int]int{}
	q := new(Queue)
	node.ColIndex = 0
	q.Push(node)

	for q.Len() > 0 {
		l := q.Len()
		for i := 0; i < l; i++ {
			top := q.Pop()
			if top.Left != nil {
				top.Left.ColIndex = top.ColIndex - 1
				q.Push(top.Left)
			}
			if top.Right != nil {
				top.Right.ColIndex = top.ColIndex + 1
				q.Push(top.Right)
			}
			if _, ok := m[top.ColIndex]; !ok {
				m[top.ColIndex] = top.Val
			}

		}
	}

}
func bottomView(node *Node) {
	m := map[int]int{}
	q := new(Queue)
	node.ColIndex = 0
	q.Push(node)

	for q.Len() > 0 {
		l := q.Len()
		for i := 0; i < l; i++ {
			top := q.Pop()
			if top.Left != nil {
				top.Left.ColIndex = top.ColIndex - 1
				q.Push(top.Left)
			}
			if top.Right != nil {
				top.Right.ColIndex = top.ColIndex + 1
				q.Push(top.Right)
			}
			m[top.ColIndex] = top.Val
		}
	}

}

func verticalTraversal(node *Node) {
	q := new(Queue)
	node.ColIndex = 0
	node.RowIndex = 0
	q.Push(node)
	m := map[int]map[int][]int{} //col -> row -> [elems]
	for q.Len() > 0 {
		l := q.Len()
		for i := 0; i < l; i++ {
			top := q.Pop()
			if top.Left != nil {
				top.Left.ColIndex = top.ColIndex - 1
				top.Left.RowIndex = top.RowIndex + 1
				q.Push(top.Left)
			}
			if top.Right != nil {
				top.Right.ColIndex = top.ColIndex - 1
				top.Right.RowIndex = top.RowIndex + 1
				q.Push(top.Right)
			}
			row := m[top.ColIndex]
			row[top.RowIndex] = append(row[top.RowIndex], top.Val)
		}
	}
}

func rightView(node *Node, level int, ds *[]int) {
	if node == nil {
		return
	}

	if len(*ds) == level {
		*ds = append(*ds, node.Val)
	}
	rightView(node.Right, level+1, ds)
	rightView(node.Left, level+1, ds)
}

func leftView(node *Node, level int, ds *[]int) {
	if node == nil {
		return
	}

	if len(*ds) == level {
		*ds = append(*ds, node.Val)
	}
	leftView(node.Left, level+1, ds)
	leftView(node.Right, level+1, ds)

}
