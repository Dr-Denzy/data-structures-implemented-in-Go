package stack

import (
	"github.com/drdenzy/containers"
	"reflect"
	"testing"
)

func TestLinkedStackInt(t *testing.T) {

	testStack := NewLinkedStack[int]()

	t.Run("LinkedStack with int data type", func(t *testing.T) {
		testStack.Push(5)
		testStack.Push(10)

		t.Run("LinkedStack Size", func(t *testing.T) {
			expected := 2
			if got := testStack.Size(); got != expected {
				t.Errorf("LinkedStack.Size() = %v, want %v", got, expected)
			}
		})

		t.Run("LinkedStack Push", func(t *testing.T) {
			expected := []int{10, 5}
			if !reflect.DeepEqual(expected, testStack.Display()) {
				t.Errorf("expected: %v, got: %v", expected, testStack.Display())
			}

		})

		poppedItem, _ := testStack.Pop()

		t.Run("LinkedStack Pop", func(t *testing.T) {

			if poppedItem != 10 {
				t.Errorf("expected: %v, got: %v", 10, poppedItem)

			}
		})

		t.Run("LinkedStack Peek", func(t *testing.T) {
			peekedItem, _ := testStack.Peek()
			if peekedItem != 5 {
				t.Errorf("expected: %v, got: %v", 5, peekedItem)
			}
		})

		_, _ = testStack.Pop()

		t.Run("LinkedStack IsEmpty True", func(t *testing.T) {
			if !testStack.IsEmpty() {
				t.Errorf("expected: %v, got: %v", true, testStack.IsEmpty())
			}

		})

		testStack.Push(15)

		t.Run("LinkedStack IsEmpty False", func(t *testing.T) {
			if testStack.IsEmpty() != false {
				t.Errorf("expected: %v, got: %v", false, testStack.IsEmpty())
			}
		})

		t.Run("LinkedStack Display", func(t *testing.T) {
			expected := []int{15}
			if !reflect.DeepEqual(expected, testStack.Display()) {
				t.Errorf("expected: %v, got: %v", expected, testStack.Display())
			}
		})

		t.Run("LinkedStack Clear", func(t *testing.T) {
			expected := NewLinkedStack[int]()
			testStack.Clear()
			if !reflect.DeepEqual(expected, testStack) {
				t.Errorf("expected: %v, got: %v", expected, testStack)
			}
		})

		t.Run("LinkedStack Custom Error", func(t *testing.T) {
			expectedErrorMsg := "Stack is empty. Zero value of the data type returned"

			_, err := testStack.Pop()

			if !containers.ErrorContains(err, expectedErrorMsg) {
				t.Errorf("expected: %v, got: %v", expectedErrorMsg, err)
			}

			_, err = testStack.Peek()

			if !containers.ErrorContains(err, expectedErrorMsg) {
				t.Errorf("expected: %v, got: %v", expectedErrorMsg, err)
			}

		})
	})

}
