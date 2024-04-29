package stack

import (
	"fmt"
	"github.com/drdenzy/containers"
)

type Stack interface {
	containers.Container
	Push(e any)
	Pop() (e any, err error)
	Peek() (e any, err error)
}

type ContiguousStack[T any] struct {
	elements []T
}

// NewContiguousStack returns a new ContiguousStack (Array Stack) instance
func NewContiguousStack[T any]() *ContiguousStack[T] {
	return &ContiguousStack[T]{}
}

// Size returns the size of an instance of ContiguousStack (Array Stack)
func (s *ContiguousStack[T]) Size() int {
	return len(s.elements)
}

// IsEmpty checks if an instance of an ContiguousStack (Array Stack) is empty or not
func (s *ContiguousStack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

// Clear clears out the content of an ContiguousStack (Array Stack)
func (s *ContiguousStack[T]) Clear() {
	s.elements = []T{}
}

// Push appends an item to the top of the stack
func (s *ContiguousStack[T]) Push(value T) {
	s.elements = append(s.elements, value)
}

// Pop returns the last (or top) item on a stack.
// Returns the zero value of the element type and an error if the stack is empty
func (s *ContiguousStack[T]) Pop() (T, error) {
	if s.Size() == 0 {
		var zeroValue T
		//return zeroValue, errors.New("the stack cannot be empty when Pop() is called")
		return zeroValue, &CustomStackError{
			Msg:  fmt.Sprintf("Stack is empty. Zero value of the elements type returned: %v", zeroValue),
			Item: zeroValue,
		}
	}
	// removes the last element on the stack
	topElement := s.elements[s.Size()-1]

	// reassign the downsized slice back to elements variable
	s.elements = s.elements[:s.Size()-1]

	return topElement, nil
}

// Peek returns the top element of a stack without removing it from the stack
// if the stack is empty it returns an error
func (s *ContiguousStack[T]) Peek() (T, error) {
	if s.Size() == 0 {
		var zeroValue T
		return zeroValue, &CustomStackError{
			Msg:  fmt.Sprintf("Stack is empty. Zero value of the elements type returned: %v", zeroValue),
			Item: zeroValue,
		}
	}
	// returns the last element on the stack without removing it from the stack
	return s.elements[s.Size()-1], nil
}
