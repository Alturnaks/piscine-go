package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	c := 0
	for l != nil {
		if c == pos {
			return l
		}
		c++
		l = l.Next
	}
	return nil
}
