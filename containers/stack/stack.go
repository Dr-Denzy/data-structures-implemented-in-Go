package stack

import "github.com/drdenzy/containers"

type Stack interface {
	containers.Container      // include Size, Clear, and IsEmpty operations or methods
	Push(e any)               // place an item at the top of a stack
	Pop() (e any, err error)  // remove and return the item at the top of a non-empty stack
	Peek() (e any, err error) // return the item at the top of a non-empty stack
}
