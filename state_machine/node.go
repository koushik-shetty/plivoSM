package state_machine

import "errors"

type Node struct {
	Name        string
	transitions map[string]string
}

func NewNode(name string) (*Node, error) {
	if name == "" {
		return nil, errors.New("Node name cannot be empty")
	}

	return &Node{
		Name:        name,
		transitions: map[string]string{},
	}, nil
}

func (n *Node) AddTransition(nodeName string) *Node {
	n.transitions[nodeName] = nodeName
	return n
}

func (n *Node) GetTransitions() []string {
	nodeNames := []string{}
	for nodeName := range n.transitions {
		nodeNames = append(nodeNames, nodeName)
	}
	return nodeNames
}
