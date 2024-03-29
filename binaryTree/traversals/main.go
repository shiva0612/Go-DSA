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

func NewNode(val int) *Node {
	return &Node{Val: val}
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
	q.arr = append(q.arr, node)
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

func (q *Queue) print() {
	fmt.Println("")
	fmt.Print("q = ")
	for _, node := range q.arr {
		fmt.Print(node.Val)
	}
	fmt.Println("")

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

	// node := cnstrTree_prein([]int{10, 20, 40, 50, 30, 60}, []int{40, 20, 50, 10, 60, 30})
	// PostR(node)
	// node := cnstrTree_postin([]int{60, 30, 50, 40, 20, 10}, []int{40, 20, 50, 10, 60, 30})
	// PostR(node)

	r := cnstrTree_prein([]int{3, 5, 6, 2, 7, 4, 1, 0, 8}, []int{6, 5, 7, 2, 4, 3, 0, 1, 8})
	ans := findNodesAtKDist(r, 5, 2)
	fmt.Println(ans)
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

// --------------------------------------------------------
func helpLeft(node *Node, ds *[]int) {
	for node != nil {
		if !isLeaf(node) {
			*ds = append(*ds, node.Val)
		}
		if node.Left != nil {
			node = node.Left
		} else {
			node = node.Right
		}
	}
}
func helpleaf(node *Node, ds *[]int) {

	if isLeaf(node) {
		*ds = append(*ds, node.Val)
		return
	}
	if node.Left != nil {
		helpleaf(node.Left, ds)
	}
	if node.Right != nil {
		helpleaf(node.Right, ds)
	}
}

func isLeaf(node *Node) bool {
	return node.Left == nil && node.Right == nil
}

func helpRightReverse(node *Node, ds *[]int) {
	stack := []int{}

	for node != nil {
		if !isLeaf(node) {
			stack = append(stack, node.Val)
		}
		if node.Right != nil {
			node = node.Right
		} else {
			node = node.Left
		}
	}

	//populate ds from stack
	for i := len(stack) - 2; i > 0; i-- {
		*ds = append(*ds, stack[i])
	}
}
func BoundaryClockWise(node *Node) []int {
	ans := []int{}
	if node == nil {
		return ans
	}
	if !isLeaf(node) {
		ans = append(ans, node.Val)
	}
	helpLeft(node.Left, &ans)
	helpleaf(node.Right, &ans)
	helpRightReverse(node, &ans)
	return ans
}

// --------------------------------------------------------
// verify again
func zigzag(node *Node) [][]int {
	ans := [][]int{}
	f := -1
	q := new(Queue)
	if node == nil {
		return ans
	}
	q.Push(node)

	for q.Len() > 0 {

		temp := []int{}
		for i := 0; i < q.Len(); i++ {
			top := q.Pop()
			if top.Left != nil {
				q.Push(top.Left)
			}
			if top.Right != nil {
				q.Push(top.Right)
			}

			if f < 0 {
				temp = append(temp, top.Val)
			} else {
				temp = append([]int{top.Val}, temp...)
			}
		}
		f *= -1
		ans = append(ans, temp)

	}
	return ans
}

// --------------------------------------------------------

func isSymetric(node *Node) bool {
	if node == nil {
		return true
	}
	return helpSymetric(node.Left, node.Right)
}

func helpSymetric(left, right *Node) bool {
	if left == nil || right == nil {
		return left == right
	}
	if left.Val != right.Val {
		return false
	}
	return helpSymetric(left.Left, right.Right) && helpSymetric(left.Right, right.Left)
}

// --------------------------------------------------------

func rootToNodePath(node *Node, arr *[]int, find int) bool {

	if node == nil {
		return false
	}
	*arr = append(*arr, node.Val)
	if node.Val == find {
		return true
	}
	if rootToNodePath(node.Left, arr, find) || rootToNodePath(node.Right, arr, find) {
		return true
	}
	*arr = (*arr)[:len(*arr)-1]
	return false

}

// --------------------------------------------------------

func lowestCommonAncestor(node *Node, p, q int) *Node {

	if node == nil {
		return node
	}
	if node.Val == p || node.Val == q {
		return node
	}
	left := lowestCommonAncestor(node.Left, p, q)
	right := lowestCommonAncestor(node.Right, p, q)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return node
}

// --------------------------------------------------------

func maxWidth(node *Node) int {

	if node == nil {
		return 0
	}
	q := new(Queue)
	q.Push(node)
	width := 0
	for q.Len() > 0 {
		l := q.Len()
		levelMin := q.Top().RowIndex
		var first, last int
		for i := 0; i < l; i++ {
			top := q.Pop()
			top.RowIndex -= levelMin
			if i == 0 {
				first = top.RowIndex
			}
			if i == l-1 {
				last = top.RowIndex
			}
			if node.Left != nil {
				node.RowIndex = 2*node.RowIndex + 1
				q.Push(node.Left)
			}
			if node.Right != nil {
				node.RowIndex = 2*node.RowIndex + 2
				q.Push(node.Right)
			}
		}
		if last-first+1 > width {
			width = last - first + 1
		}
	}
	return width
}

// --------------------------------------------------------

func findNodesAtKDist(root *Node, n, k int) []int {
	ans := []int{}
	parent := map[*Node]*Node{}
	node := findNodeWithVal(root, n)
	if node == nil {
		return ans //source node not found
	}
	helpGetNodeAssignParent(root, nil, parent)

	q := new(Queue)
	q.Push(node)
	vis := []int{}
	dist := k

	for q.Len() > 0 && dist > 0 {
		l := q.Len()
		for i := 0; i < l; i++ {
			top := q.Pop()
			if top.Left != nil && !isVisited(vis, top.Left.Val) {
				q.Push(top.Left)
			}
			if top.Right != nil && !isVisited(vis, top.Right.Val) {
				q.Push(top.Right)
			}
			if parent[top] != nil && !isVisited(vis, parent[top].Val) {
				q.Push(parent[top])
			}
			vis = append(vis, top.Val)
		}
		dist--
	}

	for q.Len() > 0 {
		top := q.Pop()
		ans = append(ans, top.Val)
	}

	return ans
}

func isVisited(vis []int, n int) bool {
	for _, v := range vis {
		if v == n {
			return true
		}
	}
	return false
}
func helpGetNodeAssignParent(node, parent *Node, m map[*Node]*Node) {
	if node == nil {
		return
	}
	m[node] = parent
	helpGetNodeAssignParent(node.Left, node, m)
	helpGetNodeAssignParent(node.Right, node, m)
}
func findNodeWithVal(node *Node, val int) *Node {
	if node == nil {
		return nil
	}
	if node.Val == val {
		return node
	}
	if l := findNodeWithVal(node.Left, val); l != nil {
		return l
	}
	if r := findNodeWithVal(node.Right, val); r != nil {
		return r
	}
	return nil
}

// --------------------------------------------------------

func cnstrTree_prein(pre, in []int) *Node {

	inMap := map[int]int{}
	for i, v := range in {
		inMap[v] = i
	}

	return helpCnstr_prein(pre, 0, len(pre)-1, in, 0, len(in)-1, inMap)
}

func helpCnstr_prein(pre []int, preStart, preEnd int, in []int, inStart, inEnd int, m map[int]int) *Node {
	if preStart > preEnd || inStart > inEnd {
		return nil
	}

	node := NewNode(pre[preStart])
	rootIdx := m[node.Val]
	numsLeft := rootIdx - inStart

	node.Left = helpCnstr_prein(pre, preStart+1, preStart+numsLeft, in, inStart, rootIdx-1, m)
	node.Right = helpCnstr_prein(pre, preStart+numsLeft+1, preEnd, in, rootIdx+1, inEnd, m)
	return node
}

// --------------------------------------------------------

func cnstrTree_postin(post, in []int) *Node {
	inMap := map[int]int{}
	for i, v := range in {
		inMap[v] = i
	}

	return helpCnstr_postin(post, 0, len(post)-1, in, 0, len(in)-1, inMap)
}

func helpCnstr_postin(post []int, postStart, postEnd int, in []int, inStart, inEnd int, m map[int]int) *Node {

	if inStart > inEnd || postStart > postEnd {
		return nil
	}
	node := NewNode(post[postEnd])
	rootIdx := m[node.Val]
	numsLeft := rootIdx - inStart

	node.Left = helpCnstr_postin(post, postEnd-numsLeft, postEnd-1, in, inStart, rootIdx-1, m)
	node.Right = helpCnstr_postin(post, postStart, postEnd-numsLeft-1, in, rootIdx+1, inEnd, m)
	return node
}

// --------------------------------------------------------
