package main

func (n *Edge) Insert(value string, data *Item) *Edge {
	// log.Println("Creating:", value)
	if n == nil {
		return nil
	}

	switch {
	case value == n.Value:
		if data != nil {
			n.ItemList = append(n.ItemList, data)
		}
		return n
	case value < n.Value:
		if n.Left == nil {
			n.Left = &Edge{Value: value}
			if data != nil {
				n.Left.ItemList = append(n.Left.ItemList, data)
			}
			return n.Left
		}
		return n.Left.Insert(value, data)
	case value > n.Value:
		if n.Right == nil {
			n.Right = &Edge{Value: value}
			if data != nil {
				n.Right.ItemList = append(n.Right.ItemList, data)
			}

			return n.Right
		}
		return n.Right.Insert(value, data)
	}
	return nil
}
func (t *Tree) TreeFind(params []string) (*Edge, bool) {

	if t.Root == nil {
		return nil, false
	}
	var edge *Edge
	edge = t.Root
	var lastValidEdge *Edge
	var matchCount int
	var totalMatches int
	for _, v := range params {
		if v == "" {
			continue
		}
		totalMatches++
		edge = edge.Find(v)
		if edge == nil {
			break
		} else {
			lastValidEdge = edge
			matchCount++
		}
	}
	if matchCount == totalMatches {
		return lastValidEdge, true
	}
	return nil, false
}
func (t *Tree) TreeFindAndInsert(params []string, bc *Item) bool {

	if t.Root == nil {
		return false
	}
	var edge *Edge
	edge = t.Root
	var matchCount int
	var edgeList []*Edge

	var totalMatches int

	for _, v := range params {
		if v == "" {
			continue
		}
		totalMatches++
		edge = edge.Find(v)
		if edge == nil {
			break
		} else {
			edgeList = append(edgeList, edge)
			matchCount++
		}
	}
	if matchCount == totalMatches {
		for _, v := range edgeList {
			v.ItemList = append(v.ItemList, bc)
		}
		return true
	}
	return false
}
func (n *Edge) Find(s string) *Edge {

	if n == nil {
		// log.Println("nill edge", n)
		return nil
	}
	// log.Println("switching on:", n.Value, "looking for", s)
	switch {

	case s == n.Value:
		// log.Println("found it !", n.Value)
		return n
	case s < n.Value:
		// log.Println("going left on", n.Value, " looking for ", s)
		return n.Left.Find(s)

	default:
		// log.Println("going right on", n.Value, " looking for ", s)
		return n.Right.Find(s)

	}
}

func (n *Edge) findMax(parent *Edge) (*Edge, *Edge) {
	if n == nil {
		return nil, parent
	}
	if n.Right == nil {
		return n, parent
	}
	return n.Right.findMax(n)
}

func (n *Edge) replaceEdge(parent, replacement *Edge) *Edge {
	if n == parent.Left {
		parent.Left = replacement
		return n
	}
	parent.Right = replacement
	return n
}

func (n *Edge) Delete(s string, parent *Edge) *Edge {

	switch {
	case s < n.Value:
		return n.Left.Delete(s, n)
	case s > n.Value:
		return n.Right.Delete(s, n)
	default:
		if n.Left == nil && n.Right == nil {
			return n.replaceEdge(parent, nil)
		}

		if n.Left == nil {
			return n.replaceEdge(parent, n.Right)
		}
		if n.Right == nil {
			return n.replaceEdge(parent, n.Left)
		}

		replacement, replParent := n.Left.findMax(n)

		n.Value = replacement.Value
		n.ItemList = replParent.ItemList

		return replacement.Delete(replacement.Value, replParent)
	}
}

func (t *Tree) Insert(value string, data *Item) *Edge {
	if t.Root == nil {
		t.Root = &Edge{Value: value}
		if data != nil {
			t.Root.ItemList = append(t.Root.ItemList, data)
		}
		return t.Root
	}
	return t.Root.Insert(value, data)
}

func (t *Tree) Find(s string) *Edge {
	if t.Root == nil {
		return nil
	}
	Edge := t.Root.Find(s)
	if Edge == nil {
		return nil
	}
	return Edge
}

func (t *Tree) FindMax() (*Edge, *Edge) {
	if t.Root == nil {
		return nil, nil
	}
	return t.Root.findMax(t.Root)
}

func (t *Tree) Delete(s string) *Edge {
	fakeParent := &Edge{Right: t.Root}
	if t.Root == nil {
		return fakeParent
	}

	return t.Root.Delete(s, fakeParent)
}

func (t *Tree) Traverse(parent string, level int, n *Edge, f func(string, int, *Edge)) {
	if n == nil {
		return
	}

	f(parent, level, n)
	level++
	t.Traverse(n.Value, level, n.Left, f)
	t.Traverse(n.Value, level, n.Right, f)
}

var TREE = &Tree{}

type Tree struct {
	Root *Edge
}
type Edge struct {
	Value    string
	ItemList []*Item
	Left     *Edge
	Right    *Edge
}
