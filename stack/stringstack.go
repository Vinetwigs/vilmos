package stack

import "errors"

type StringStack struct {
	items []string
}

func NewStringStack() *StringStack {
	stack := new(StringStack)
	return stack
}

func (stack *StringStack) Push(val string) {
	stack.items = append(stack.items, val)
}

func (stack *StringStack) Pop() (string, error) {
	if len(stack.items) == 0 {
		return "", errors.New("error: stack is empty, cannot pop")
	}

	item := stack.items[len(stack.items)-1]

	return item, nil
}

func (stack *StringStack) Size() int {
	return len(stack.items)
}

func (stack *StringStack) IsEmpty() bool {
	if stack.Size() == 0 {
		return true
	} else {
		return false
	}
}
