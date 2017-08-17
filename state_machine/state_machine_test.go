package state_machine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStateMachinAddNodeReturnsErrorIfNodeNameIsEmpty(t *testing.T) {
	sm := New()

	err := sm.AddNode("")

	assert.Error(t, err)
}

func TestStateMachinAddNodeReturnsErrorIfNodeAlreadyExists(t *testing.T) {
	sm := New()

	err := sm.AddNode("A")
	assert.NoError(t, err, "expected no error while adding initial node")

	err = sm.AddNode("A")
	assert.Error(t, err)
}

func TestStateMachinAddNodeReturnsErrorIfNodeCreationFails(t *testing.T) {
	sm := New()

	err := sm.AddNode("")
	assert.Error(t, err, "expected no error while adding initial node")
}

func TestStateMachinAddNodeAddsANodeToTheStateMachine(t *testing.T) {
	sm := New()

	err := sm.AddNode("A")
	assert.NoError(t, err, "expected no error while adding initial node")

	err = sm.AddNode("B")
	assert.NoError(t, err)

	err = sm.AddTransition("A", "B")
	assert.NoError(t, err)

	err = sm.TransitionTo("B")
	assert.NoError(t, err)
}

func TestStateMachineTransitionToTransitionsToANewState(t *testing.T) {
	sm := New()

	err := sm.AddNode("A")
	assert.NoError(t, err)
	currentNode := sm.GetCurrentNode()
	assert.Equal(t, "A", currentNode.Name)

	err = sm.AddNode("B")
	assert.NoError(t, err)
	currentNode = sm.GetCurrentNode()
	assert.Equal(t, "A", currentNode.Name)

	err = sm.AddTransition("A", "B")
	assert.NoError(t, err)

	err = sm.TransitionTo("B")
	assert.NoError(t, err)
	currentNode = sm.GetCurrentNode()
	assert.Equal(t, "B", currentNode.Name)
}

func TestStateMachineTransitionToErrorsIfStateIsEmpty(t *testing.T) {
	sm := New()

	err := sm.TransitionTo("B")
	assert.Error(t, err)
}

func TestStateMachineTransitionToErrorsIfStateDoesNotExists(t *testing.T) {
	sm := New()

	err := sm.AddNode("A")
	assert.NoError(t, err)
	currentNode := sm.GetCurrentNode()
	assert.Equal(t, "A", currentNode.Name)

	err = sm.TransitionTo("B")
	assert.Error(t, err)
}

func TestStateMachineGetCurrentNodeReturnsNilIfThereAreNoNodes(t *testing.T) {
	sm := New()

	node := sm.GetCurrentNode()
	assert.Nil(t, node)
}

func TestStateMachineGetCurrentNodeReturnsTheCurrentNode(t *testing.T) {
	sm := New()

	err := sm.AddNode("A")
	assert.NoError(t, err)

	err = sm.AddNode("B")
	assert.NoError(t, err)

	currentNode := sm.GetCurrentNode()
	assert.Equal(t, "A", currentNode.Name)
}

func TestStateMachinePrintsEmptyIfStateIsEmpty(t *testing.T) {
	sm := New()
	stateString := sm.FormatStateMachine()
	assert.Equal(t, "(Empty)", stateString)
}

func TestStateMachinePrintsStateWithTransitionsIfStateIsNotEmpty(t *testing.T) {
	sm := New()

	err := sm.AddNode("A")
	assert.NoError(t, err)

	err = sm.AddNode("B")
	assert.NoError(t, err)

	err = sm.AddTransition("A", "B")
	assert.NoError(t, err)

	err = sm.TransitionTo("B")
	assert.NoError(t, err)

	expectedStateString := "A(A -> B)\nB()"
	stateString := sm.FormatStateMachine()
	assert.Equal(t, expectedStateString, stateString)
}

func TestStateMachinAddTransitionAddsATranstionToTheStateMachineNode(t *testing.T) {
	sm := New()

	err := sm.AddNode("A")
	assert.NoError(t, err, "expected no error while adding initial node")

	err = sm.AddNode("B")
	assert.NoError(t, err)

	err = sm.AddTransition("A", "B")
	assert.NoError(t, err)

	err = sm.TransitionTo("B")
	assert.NoError(t, err)
}

func TestStateMachinAddTransitionFailsIfTheNodesDoesNotExists(t *testing.T) {
	sm := New()

	err := sm.AddNode("A")
	assert.NoError(t, err)

	err = sm.AddTransition("A", "B")
	assert.Error(t, err)
}
