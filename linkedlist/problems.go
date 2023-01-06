package linkedlist

func (a *List) rotateByK(k int) {

	length := a.length()
	if k >= length {
		return
	}
	temp := a.head
	var last *Node
	curlen := 0
	for (length - curlen) != k {
		curlen += 1
		if (length - curlen) == k {
			last = temp
		}
		temp = temp.next
	}
	last.next = nil

	if temp.next == nil {
		temp.next = a.head
		a.head = temp
	} else {
		p, c, n := temp, temp.next, temp.next.next
		temp.next = nil
		for {
			c.next = p
			p = c
			c = n
			if n == nil {
				break
			}
			n = n.next
		}

		for p != nil {
			move := p
			p = p.next
			move.next = a.head
			a.head = move
		}
	}

}

func (a *List) revereBySize(size int) {

	k := a.head
	arr := make([]int, 0)
	hopTillSize := func() bool {
		arr = make([]int, 0)
		for i := size; i > 0 && k != nil; i-- {
			arr = append(arr, k.val)
			k = k.next
		}
		if k == nil && len(arr) < 3 {
			return false
		}
		return true
	}

	temp := a.head
	for hopTillSize() {
		for i := len(arr) - 1; i >= 0; i-- {
			temp.val = arr[i]
			temp = temp.next
		}
	}
	a.printlist()
}

func (a *List) checkIfPalindrome() bool {

	b := a.copy()
	b.reverse()
	i, j := a.head, b.head
	for i != nil && j != nil {
		if i.val != j.val {
			return false
		}
		i, j = i.next, j.next
	}
	if i == nil && j == nil {
		return true
	}

	return false
}
func (a *List) checkIfPalindrome2() bool {

	mid := a.findmid()
	secHalf := &List{head: mid.next}
	secHalf.reverse()

	tempFirHalf := a.head
	tempSecHalf := secHalf.head
	for tempSecHalf != nil {

		if tempSecHalf.val != tempFirHalf.val {
			return false
		}

		tempSecHalf = tempSecHalf.next
		tempFirHalf = tempFirHalf.next
	}

	return true
}

func mergeList(a, b *List) *List {
	l := new(List)

	i, j, k := a.head, b.head, l.head

	for i != nil && j != nil {
		node := new(Node)

		if i.val <= j.val {
			node.val = i.val
			i = i.next
		} else {
			node.val = j.val
			j = j.next
		}
		if l.head == nil {
			l.head = node
			k = l.head
		} else {
			k.next = node
			k = k.next
		}
	}
	if i == nil {
		for j != nil {
			k.next = j
			j = j.next
			k = k.next
		}
	}
	if j == nil {
		for i != nil {
			k.next = i
			i = i.next
			k = k.next
		}
	}
	return l
}
func addNumbers(a *List, b *List) *List {
	ans := new(List)

	i, j, k := a.head, b.head, ans.head
	for i != nil && j != nil {
		add := i.val + j.val
		if add/10 != 0 {
			if i.next == nil {
				node := new(Node)
				i.next = node
			}
			i.next.val += add / 10
		}
		node := &Node{val: add % 10}
		if ans.head == nil {
			ans.head = node
			k = ans.head
		}
		k.next = node
		k = k.next
		i = i.next
		j = j.next
	}
	if i == nil {
		for j != nil {
			node := &Node{val: j.val}
			k.next = node
			j = j.next
			k = k.next
		}
	}
	if j == nil {
		for i != nil {
			node := &Node{val: i.val}
			k.next = node
			i = i.next
			k = k.next
		}
	}

	return ans
}

func findYjunction(a, b *List) bool {

	i := a.head
	for i != nil {
		j := b.head
		for j != nil {
			if j == i {
				return true
			}
			j = j.next
		}
		i = i.next
	}

	return false
}

func findIfLoop(a *List) bool {

	s, f := a.head, a.head
	for s != nil {
		if s == f {
			return true
		}
		s = s.next
		f = f.next.next
	}
	return false
}

func findIfLoop2(a *List) bool {

	m := map[*Node]int{}

	i := 0
	temp := a.head
	for temp != nil {

		if _, present := m[temp]; present {
			return true
		}
		m[temp] = i
		temp = temp.next
		i += 1
	}

	return false
}
