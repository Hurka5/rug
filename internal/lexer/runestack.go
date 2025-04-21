package lexer

type runeNode struct {
	r    rune
	next *runeNode
}

type runeStack struct {
	start *runeNode
}

func newRuneStack() runeStack {
	return runeStack{}
}

func (s *runeStack) push(r rune) {
	node := &runeNode{r: r}
	if s.start != nil {
		node.next = s.start
	}
	s.start = node
}

func (s *runeStack) pop() rune {
	if s.start == nil {
		return EOF
	} else {
		n := s.start
		s.start = n.next
		return n.r
	}
}

func (s *runeStack) clear() {
	s.start = nil
}
