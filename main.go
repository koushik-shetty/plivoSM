package main

import (
	"errors"
	"fmt"
	"plivoSM/state_machine"
	"strings"
)

const (
	AddNode       = "a"
	Transition    = "t"
	AddTransition = "+"
	Print         = "p"
	PrintCurrent  = "c"
	PrintState    = "s"
	AddHook       = "h"
)

func main() {
	fmt.Println("******************************************************\n\n")
	fmt.Println("Welcome to Plivo state machine client\n\n")
	fmt.Println("a <space> <node_name> to add node")
	fmt.Println("+ <space> <fromNode>,<toNode> to add transition to node")
	fmt.Println("t <space> <node_name> to transition to node")
	fmt.Println("p <space> c to print current node")
	fmt.Println("p <space> s to print state\n\n")
	fmt.Println("h <space> s to add hook\n\n")
	fmt.Println("******************************************************\n\n")

	sm := state_machine.New()
	for {
		command := ""
		arg := ""
		fmt.Scanf("%s%s\n", &command, &arg)
		if validCommand(command) {
			err := operate(sm, command, arg)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
		} else {
			fmt.Printf("invalid command: %v", arg)
		}
	}

}

func validCommand(c string) bool {
	if c == AddNode || c == Transition || c == Print || c == AddTransition || c == AddHook {
		return true
	}
	return false
}

func operate(sm state_machine.SM, command, arg string) error {
	if !validCommand(command) {
		return errors.New(fmt.Sprintf("Invalid command: [%s]", command))
	}
	switch command {
	case AddNode:
		err := sm.AddNode(arg)
		if err != nil {
			return err
		}
		fmt.Printf("Added node: %s\n", arg)
	case AddTransition:
		s := strings.Split(arg, ",")
		err := sm.AddTransition(strings.Trim(s[0], " "), strings.Trim(s[1], " "))
		if err != nil {
			return err
		}
		fmt.Printf("Added Transition: %s\n", arg)
	case Transition:
		node := sm.GetCurrentNode()
		err := sm.TransitionTo(arg)
		if err != nil {
			return err
		}
		fmt.Printf("Transitioned from %s -> %s\n", node.Name, arg)
	case Print:
		state := ""
		if arg == PrintState {
			state = sm.FormatStateMachine()
		} else if arg == PrintCurrent {
			currentNode := sm.GetCurrentNode()
			if currentNode != nil {
				state = fmt.Sprintf("Current node: %s", currentNode.Name)
			}
		} else {
			return errors.New(fmt.Sprintf("invalid Print arg: %c", arg))
		}
		fmt.Println(state)
		fmt.Println()
		return nil
	case AddHook:
		sm.AddHook(func(a, b string) { fmt.Printf("In hook, Transition from %s, %s\n", a, b) })
		fmt.Printf("Added Hook\n")
	default:
		return errors.New("Invalid command")
	}
	return nil
}

func validPrintArg(arg string) bool {
	if arg == PrintCurrent || arg == PrintState {
		return true
	}
	return false
}
