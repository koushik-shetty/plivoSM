package main

import (
	"errors"
	"plivoSM/state_machine"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/mock"
)

type fakeStateMachine struct {
	mock.Mock
}

func (f *fakeStateMachine) AddNode(name string) error {
	args := f.Called(name)
	return args.Error(0)
}

func (f *fakeStateMachine) TransitionTo(name string) error {
	args := f.Called(name)
	return args.Error(0)
}

func (f *fakeStateMachine) GetCurrentNode() *state_machine.Node {
	args := f.Called()
	if args.Get(0) != nil {
		return args.Get(0).(*state_machine.Node)
	}
	return nil
}

func (f *fakeStateMachine) FormatStateMachine() string {
	args := f.Called()
	return args.String(0)
}

func (f *fakeStateMachine) AddTransition(from, to string) error {
	args := f.Called(from, to)
	return args.Error(0)
}

func (f *fakeStateMachine) AddHook(fn func(a, b string)) {
	f.Called(fn)
}

func TestOperateShouldPerformTheNecessaryOperationIfCommanValid(t *testing.T) {
	sm := &fakeStateMachine{}
	sm.On("AddNode", "A").Return(nil)
	err := operate(sm, AddNode, "A")

	assert.NoError(t, err)
}

func TestOperateShouldFailIfInvalidCommand(t *testing.T) {
	sm := &fakeStateMachine{}
	err := operate(sm, "blah", "A")

	assert.Error(t, err)
}

func TestOperateShouldFailIfInvalidArgsToPrint(t *testing.T) {
	sm := &fakeStateMachine{}
	err := operate(sm, Print, "A")

	assert.Error(t, err)
}

func TestOperateShouldCallGetCurrentNodeForPrintCurrent(t *testing.T) {
	sm := &fakeStateMachine{}

	sm.On("GetCurrentNode").Return(state_machine.NewNode("A"))
	err := operate(sm, Print, PrintCurrent)

	assert.NoError(t, err)
	sm.AssertExpectations(t)
}

func TestOperateShouldCallFormatStateMachineForPrintState(t *testing.T) {
	sm := &fakeStateMachine{}
	sm.On("FormatStateMachine").Return("STATE")

	err := operate(sm, Print, PrintState)

	assert.NoError(t, err)
	sm.AssertExpectations(t)
}

func TestOperateShouldCallAddNodeForCommand(t *testing.T) {
	sm := &fakeStateMachine{}

	sm.On("AddNode", "A").Return(nil)
	err := operate(sm, AddNode, "A")

	assert.NoError(t, err)
	sm.AssertExpectations(t)
}

func TestOperateShouldFailIfAddNodeFails(t *testing.T) {
	sm := &fakeStateMachine{}

	sm.On("AddNode", "A").Return(errors.New("Add Node Failed"))
	err := operate(sm, AddNode, "A")

	assert.Error(t, err)
	assert.Equal(t, "Add Node Failed", err.Error())
	sm.AssertExpectations(t)
}

func TestOperateShouldFailIfTransitionToFails(t *testing.T) {
	sm := &fakeStateMachine{}

	sm.On("GetCurrentNode").Return(state_machine.NewNode("A"))
	sm.On("TransitionTo", "A").Return(errors.New("Transition Node Failed"))
	err := operate(sm, Transition, "A")

	assert.Error(t, err)
	assert.Equal(t, "Transition Node Failed", err.Error())
	sm.AssertExpectations(t)
}

func TestOperateShouldCallTransitionToForValidCommand(t *testing.T) {
	sm := &fakeStateMachine{}

	sm.On("GetCurrentNode").Return(state_machine.NewNode("A"))
	sm.On("TransitionTo", "A").Return(nil)
	err := operate(sm, Transition, "A")

	assert.NoError(t, err)
	sm.AssertExpectations(t)
}
