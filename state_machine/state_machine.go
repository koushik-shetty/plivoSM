package state_machine

import (
	"errors"
	"fmt"
	"strings"
)

type SM interface {
	AddNode(string) error
	AddTransition(from, to string) error
	TransitionTo(name string) error
	GetCurrentNode() *Node
	FormatStateMachine() string
	AddHook(func(a, b string))
}

type StateMachine struct {
	nodes        []*Node
	currentNode  *Node
	nextLocation int
	hook         func(a, b string)
}

func New() *StateMachine {
	return &StateMachine{hook: func(a, b string) {}}
}

func (sm *StateMachine) AddNode(name string) error {
	if sm.getNode(name) != nil {
		return errors.New("Node with name already exists")
	}

	node, err := NewNode(name)
	if err != nil {
		return err
	}

	sm.nodes = append(sm.nodes, node)
	if sm.currentNode == nil {
		sm.currentNode = node
	}
	return nil
}

func (sm *StateMachine) AddTransition(from, to string) error {
	fromNode := sm.getNode(from)
	toNode := sm.getNode(to)
	if fromNode == nil || toNode == nil {
		return errors.New("Node(s) does not exists")
	}

	fromNode.AddTransition(toNode.Name)
	return nil
}

func (sm *StateMachine) TransitionTo(name string) error {
	if sm.currentNode == nil {
		return errors.New("State is empty and cannot tranition")
	}

	if sm.currentNode.Name == name {
		return nil
	}
	if !sm.isValidTransition(name) {
		return errors.New(fmt.Sprintf("Cannot transition from %s state to %s state", sm.currentNode.Name, name))
	}

	node := sm.getNode(name)
	if node == nil {
		return fmt.Errorf("Node %s does not exists", name)
	}

	sm.hook(sm.currentNode.Name, node.Name)
	sm.currentNode = node
	return nil
}

func (sm *StateMachine) isValidTransition(name string) bool {
	for _, v := range sm.currentNode.transitions {
		if v == name {
			return true
		}
	}
	return false
}

func (sm *StateMachine) GetCurrentNode() *Node {
	return sm.currentNode
}

func (sm *StateMachine) FormatStateMachine() string {
	if len(sm.nodes) == 0 {
		return "(Empty)"
	}
	state := []string{}
	for _, node := range sm.nodes {
		state = append(state, sm.formatTransitions(node))
	}
	return strings.Join(state, "\n")
}

func (sm *StateMachine) getNode(name string) *Node {
	for _, n := range sm.nodes {
		if n.Name == name {
			return n
		}
	}
	return nil
}

func (sm *StateMachine) AddHook(f func(a, b string)) {
	sm.hook = f
}

func (sm *StateMachine) formatTransitions(node *Node) string {
	// fmt.Println("---")
	if node != nil {
		transitionHistory := []string{}
		for _, transition := range node.GetTransitions() {
			transitionHistory = append(transitionHistory, fmt.Sprintf("%s -> %s", node.Name, transition))
		}
		// fmt.Printf("%v\n", transitionHistory)
		// fmt.Println("---")
		return fmt.Sprintf("%s(%s)", node.Name, strings.Join(transitionHistory, ", "))
	}
	return ""
}
