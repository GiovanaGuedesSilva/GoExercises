// calculator/calculator.go
package calculator

// node represents an element in the stack, holding an integer value
type node struct {
	Data int
}

// Stack represents the stack itself, containing a slice of nodes
type Stack struct {
	nodes []node
}

// MakeStack creates and returns a new empty stack
func MakeStack() Stack {
	return Stack{
		nodes: []node{},
	}
}

// makeNode creates a new node with the given value
func makeNode(data int) node {
	return node{Data: data}
}

// Push adds a new element to the top of the stack
func (s *Stack) Push(data int) {
	s.nodes = append(s.nodes, makeNode(data))
}

// Pop removes and returns the element at the top of the stack
func (s *Stack) Pop() *node {
	l := len(s.nodes)
	if l == 0 {
		return nil // empty stack
	}
	n := &s.nodes[l-1:][0] // get the last element
	s.nodes = s.nodes[:l-1] // remove the last element
	return n
}

// Top returns the element at the top of the stack without removing it
func (s *Stack) Top() *node {
	if len(s.nodes) == 0 {
		return nil // empty stack
	}
	n := s.nodes[len(s.nodes)-1]
	return &n
}