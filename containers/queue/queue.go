package queue

import "github.com/drdenzy/containers"

type Queue interface {
	containers.Container       // include Size, Clear, and IsEmpty operations or methods
	Front() (any, error)       // return the item at the frontIndex of a non-empty queue
	Back() (any, error)        // return the item at the back or tail of a non-empty queue
	Enqueue(any) (bool, error) // Places an item at the tail or back or rearIndex of a queue
	Dequeue() (any, error)     // remove and return the item at the back or tail of a non-empty queue
}

type StaticQueue interface {
	Queue
	IsFull() bool
}
