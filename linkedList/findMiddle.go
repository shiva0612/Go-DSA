package main

func findMiddleOfList(head *List) *List {
	var s, f *List
	s, f = head, head

	for {
		if f.Nxt == nil {
			return s
		} else if f.Nxt.Nxt == nil {
			return s.Nxt
		}
		s = s.Nxt
		f = f.Nxt.Nxt
	}

}
