package stack

import (
	"fmt"
	"github.com/drdenzy/containers"
)

// node represents a node in the linked list used by LinkedStack.
type node[T any] struct {
	data T        // Data stored in the node
	next *node[T] // Pointer to the next node in the stack
}

// LinkedStack represents a stack implemented using a singly linked list.
type LinkedStack[T any] struct {
	top  *node[T] // Pointer to the top of the stack
	size int      // Number of elements in the stack
}

// NewLinkedStack creates a new empty LinkedStack.
func NewLinkedStack[T any]() *LinkedStack[T] {
	return &LinkedStack[T]{}
}

// Size returns the number of elements in the stack.
func (s *LinkedStack[T]) Size() int {
	return s.size
}

// IsEmpty checks if the stack is empty.
func (s *LinkedStack[T]) IsEmpty() bool {
	return s.size == 0
}

// Clear clears the stack, removing all elements.
func (s *LinkedStack[T]) Clear() {
	s.size = 0
	s.top = nil
}

// Push adds an element to the top of the stack.
func (s *LinkedStack[T]) Push(data T) {
	s.top = &node[T]{data: data, next: s.top} // Create a new node and link it to the top
	s.size++
}

// Pop removes and returns the element from the top of the stack.
func (s *LinkedStack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue, &containers.CustomError{
			Msg:  fmt.Sprintf("Stack is empty. Zero value of the data type returned: %v", zeroValue),
			Item: zeroValue,
		}
	}
	poppedItem := s.top.data // Get the data from the top node
	s.top = s.top.next       // Move the top pointer to the next node
	s.size--
	return poppedItem, nil
}

// Peek returns the element from the top of the stack without removing it.
func (s *LinkedStack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue, &containers.CustomError{
			Msg:  fmt.Sprintf("Stack is empty. Zero value of the data type returned: %v", zeroValue),
			Item: zeroValue,
		}
	}
	peekedItem := s.top.data // Get the data from the top node
	return peekedItem, nil
}

// Display returns a slice containing all elements in the stack.
func (s *LinkedStack[T]) Display() []T {
	dataFromNodes := make([]T, 0, s.size)
	currentNode := s.top
	for currentNode != nil {
		dataFromNodes = append(dataFromNodes, currentNode.data) // Append data from each node to the slice
		currentNode = currentNode.next                          // Move to the next node
	}
	return dataFromNodes
}
