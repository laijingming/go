package tree

func (n *TN) Traverse() {
	n.TraverseFunc(func(tn *TN) {
		tn.Print()
	})
}

func (n *TN) TraverseFunc(f func(*TN)) {
	if n == nil {
		return
	}
	n.Left.TraverseFunc(f)
	f(n)
	n.Right.TraverseFunc(f)
}

func (n *TN) TraverseWithChannel() chan *TN {
	c := make(chan *TN)
	go func() {
		n.TraverseFunc(func(n *TN) {
			c <- n
		})
		close(c)
	}()
	return c
}
