package piscine

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}

	c := l.Head
	for c.Next != nil {
		c = c.Next
	}
	return c.Data
}
