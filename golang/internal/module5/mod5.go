package module5

type Node struct {
	left  *Node
	right *Node
	data  int
}

func (s *Node) SetValue(val int) {
	s.data = val
}

func (s *Node) GetValue() int {
	return s.data
}

type Tree struct {
	root *Node
}

func NewNode(data int) *Node {
	return &Node{
		data: data,
	}
}

func NewTree(n Node) *Tree {
	return &Tree{
		root: &n,
	}
}

func (self *Tree) Insert(item int) {
	self.root = self.innerInsert(item, self.root)
}

func (self *Tree) innerInsert(value int, current *Node) *Node {
	if current == nil {
		return NewNode(value)
	}
	if current.GetValue() > value {
		current.left = self.innerInsert(value, current.left)
	} else if current.GetValue() < value {
		current.right = self.innerInsert(value, current.right)
	}
	return current
}

func (self *Tree) GetElements() []int {
	elements := make([]int, 0)
	elements = self.innerTraversal(self.root, elements)
	return elements
}

func (self *Tree) innerTraversal(current *Node, elements []int) []int {
	if current == nil {
		return elements
	}
	elements = self.innerTraversal(current.left, elements)
	elements = append(elements, current.data)
	elements = self.innerTraversal(current.right, elements)
	return elements
}

func (self *Tree) Find(value int) bool {
	return self.innerFind(value, self.root)
}

func (self *Tree) innerFind(value int, current *Node) bool {
	if current == nil {
		return false
	}
	if current.data == value {
		return true
	}
	if current.data > value {
		return self.innerFind(value, current.left)
	}
	return self.innerFind(value, current.right)
}
