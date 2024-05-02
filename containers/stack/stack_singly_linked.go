package stack

import (
	"fmt"
	"github.com/drdenzy/containers"
)

type node[T any] struct {
	data T
	next *node[T]
}

type LinkedStack[T any] struct {
	top  *node[T]
	size int
}

func NewLinkedStack[T any]() *LinkedStack[T] {
	return &LinkedStack[T]{}
}

func (s *LinkedStack[T]) Size() int {
	return s.size
}

func (s *LinkedStack[T]) IsEmpty() bool {
	return s.size == 0
}

func (s *LinkedStack[T]) Clear() {
	s.size = 0
	s.top = nil
}

func (s *LinkedStack[T]) Push(data T) {
	s.top = &node[T]{data: data, next: s.top} // the linking happens here
	s.size++
}

func (s *LinkedStack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue, &containers.CustomError{
			Msg:  fmt.Sprintf("Stack is empty. Zero value of the data type returned: %v", zeroValue),
			Item: zeroValue,
		}
	}
	poppedItem := s.top.data
	s.top = s.top.next
	s.size--
	return poppedItem, nil
}

func (s *LinkedStack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue, &containers.CustomError{
			Msg:  fmt.Sprintf("Stack is empty. Zero value of the data type returned: %v", zeroValue),
			Item: zeroValue,
		}
	}
	peekedItem := s.top.data
	return peekedItem, nil
}

func (s *LinkedStack[T]) Display() []T {
	dataFromNodes := make([]T, 0, s.size)
	currentNode := s.top
	for currentNode != nil {
		dataFromNodes = append(dataFromNodes, currentNode.data)
		currentNode = currentNode.next
	}
	return dataFromNodes
}
