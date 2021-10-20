package stack

import (
	"errors"
)

type Stack struct {
	items []uint8
}

func NewStack() *Stack {
	stack := new(Stack)
	return stack
}

func (stack *Stack) Push(val uint8) {
	stack.items = append(stack.items, val)
}

func (stack *Stack) Pop() (uint8, error) {
	if len(stack.items) == 0 {
		return 0, errors.New("error: stack is empty, cannot pop")
	}

	index := len(stack.items) - 1 // Get the index of the top most element.
	item := stack.items[index]    // Index into the slice and obtain the element.
	stack.items = stack.items[:index]
	return item, nil
}

func (stack *Stack) Size() int {
	return len(stack.items)
}

func (stack *Stack) IsEmpty() bool {
	if stack.Size() == 0 {
		return true
	} else {
		return false
	}
}
