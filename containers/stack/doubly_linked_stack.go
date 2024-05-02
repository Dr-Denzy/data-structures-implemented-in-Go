package stack

import (
	"fmt"
	"github.com/drdenzy/containers"
)

// dnode represents a node in the doubly linked list used by DoublyLinkedStack.
type dnode[T any] struct {
	data T         // Data stored in the node
	prev *dnode[T] // Pointer to the previous node in the stack
	next *dnode[T] // Pointer to the next node in the stack
}

// DoublyLinkedStack represents a stack implemented using a doubly linked list.
type DoublyLinkedStack[T any] struct {
	top  *dnode[T] // Pointer to the top of the stack
	size int       // Number of elements in the stack
}

// NewDoublyLinkedStack creates a new empty DoublyLinkedStack.
func NewDoublyLinkedStack[T any]() *DoublyLinkedStack[T] {
	return &DoublyLinkedStack[T]{}
}

// Size returns the number of elements in the stack.
func (s *DoublyLinkedStack[T]) Size() int {
	return s.size
}

// IsEmpty checks if the stack is empty.
func (s *DoublyLinkedStack[T]) IsEmpty() bool {
	return s.size == 0
}

// Push adds an element to the top of the stack.
func (s *DoublyLinkedStack[T]) Push(data T) {
	newNode := &dnode[T]{data: data, prev: nil, next: s.top}

	if s.top != nil {
		s.top.prev = newNode
	}
	s.top = newNode
	s.size++
}

// Pop removes and returns the element from the top of the stack.
func (s *DoublyLinkedStack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue, &containers.CustomError{
			Msg:  fmt.Sprintf("Stack is empty. Zero value of the data type returned: %v", zeroValue),
			Item: zeroValue,
		}
	}
	poppedItem := s.top.data
	s.top = s.top.next
	if s.top != nil {
		s.top.prev = nil
	}
	s.size--
	return poppedItem, nil
}

// Peek returns the element from the top of the stack without removing it.
func (s *DoublyLinkedStack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue, &containers.CustomError{
			Msg:  fmt.Sprintf("Stack is empty. Zero value of the data type returned: %v", zeroValue),
			Item: zeroValue,
		}
	}

	return s.top.data, nil
}

// Display returns a slice containing all elements in the stack.
func (s *DoublyLinkedStack[T]) Display() []T {
	var elements []T

	currentNode := s.top
	for currentNode != nil {
		elements = append(elements, currentNode.data)
		currentNode = currentNode.next
	}
	return elements
}

// Clear clears the stack, removing all elements.
func (s *DoublyLinkedStack[T]) Clear() {
	s.top = nil
	s.size = 0
}
