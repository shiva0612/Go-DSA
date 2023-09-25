package main

func removeNode(node *List) {

	t := node
	for t.Nxt.Nxt != nil {
		t.Val = t.Nxt.Val
		t = t.Nxt
	}
	t.Val = t.Nxt.Val
	t.Nxt = nil
}
