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

type BinaryTree struct {
	root *Node
	size int
}

func NewNode(data int) Node {
	return Node{
		data: data,
	}
}

func MakeTree(n Node) BinaryTree {
	return BinaryTree{
		root: &n,
		size: 1,
	}
}

func (self BinaryTree) Insert(n Node) {

}
