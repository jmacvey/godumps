package stack

// Stack - Stack Data Structure
type Stack struct {
	Size int
	arr  []int
}

// NewStack - Parameterized Stack Constructor
func NewStack(size int, arr []int) *Stack {
	return &Stack{Size: size, arr: arr}
}

// DefaultStack - Default Stack Constructor
func DefaultStack() *Stack {
	return &Stack{Size: 0, arr: []int{}}
}

// Push - Pushes an item on the stack
func (stack *Stack) Push(item int) {
	stack.arr = append(stack.arr[0:stack.Size], item)
	stack.Size++
}

// Pop - Pops an item from the stack
func (stack *Stack) Pop() int {
	if stack.Size == 0 {
		panic("Tried to Pop From an empty stack")
	}
	stack.Size--
	return stack.arr[stack.Size]
}

// IsEmpty - Checks if stack is empty
func (stack *Stack) IsEmpty() bool {
	return stack.Size == 0
}

// Top - Gets the top of the stack
func (stack *Stack) Top() int {
	if stack.Size == 0 {
		panic("Tried to get top of an empty stack")
	}
	return stack.arr[stack.Size]
}
