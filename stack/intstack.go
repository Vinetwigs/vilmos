package stack

type IntStack struct {
	items []int
}

func NewIntStack() *IntStack {
	stack := make(IntStack)
	stack.items = []int

	return &stack
}

func (stack *IntStack) Push(val int) {
	if stack.items == nil {
		stack.items = []int
	}

	stack.items = append(stack.items, val)
}

func (stack *IntStack) Pop() int {
	if len(stack.items) == 0 {
		return nil
	}

	item := stack.items[len(stack.items)-1]

	return item
}

func (stack *IntStack) Size() int {
	return len(stack.items)
}

func (stack *IntStack) IsEmpty() bool {
	if stack.Size() == 0 {
		return true
	} else {
		return false
	}
}
