package stack

import (
	"fmt"
)

type dnode[T any] struct {
	data T
	prev *dnode[T]
	next *dnode[T]
}

type DoublyLinkedStack[T any] struct {
	top       *dnode[T]
	stackSize int
}

func NewDoublyLinkedStack[T any]() *DoublyLinkedStack[T] {
	return &DoublyLinkedStack[T]{}
}

func (s *DoublyLinkedStack[T]) Size() int {
	return s.stackSize
}

func (s *DoublyLinkedStack[T]) IsEmpty() bool {
	return s.stackSize == 0
}

func (s *DoublyLinkedStack[T]) Push(data T) {

	newNode := &dnode[T]{data: data, prev: nil, next: s.top}

	if s.top != nil {
		s.top.prev = newNode
	}
	s.top = newNode
	s.stackSize++
}

func (s *DoublyLinkedStack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue, &CustomStackError{
			Msg:  fmt.Sprintf("Stack is empty. Zero value of the data type returned: %v", zeroValue),
			Item: zeroValue,
		}
	}
	poppedItem := s.top.data
	s.top = s.top.next
	if s.top != nil {
		s.top.prev = nil
	}
	s.stackSize--
	return poppedItem, nil
}

func (s *DoublyLinkedStack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue, &CustomStackError{
			Msg:  fmt.Sprintf("Stack is empty. Zero value of the data type returned: %v", zeroValue),
			Item: zeroValue,
		}
	}

	return s.top.data, nil
}

func (s *DoublyLinkedStack[T]) Display() []T {
	var elements []T

	currentNode := s.top
	for currentNode != nil {
		elements = append(elements, currentNode.data)
		currentNode = currentNode.next
	}
	return elements
}

func (s *DoublyLinkedStack[T]) Clear() {
	s.top = nil
	s.stackSize = 0
}
