package state_machine

import "testing"
import "github.com/stretchr/testify/assert"

func TestNewNodeErrorsIfNodeNameIsEmpty(t *testing.T) {
	_, err := NewNode("")

	assert.Error(t, err)
}

func TestNewNodeReturnsNewNode(t *testing.T) {
	node, err := NewNode("A")

	assert.NoError(t, err)
	assert.Equal(t, "A", node.Name)
}

func TestNodeAddTransitionAddsTransitionToTheList(t *testing.T) {
	node, err := NewNode("A")

	assert.NoError(t, err)
	assert.Equal(t, "A", node.Name)

	node.AddTransition("B")
	assert.Equal(t, []string{"B"}, node.GetTransitions())
}
