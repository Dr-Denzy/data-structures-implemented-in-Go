package stack

import (
	"reflect"
	"testing"
)

func TestDoublyLinkedStackInt(t *testing.T) {

	testDLS := NewDoublyLinkedStack[int]()

	t.Run("Doubly Linked Stack containing Int Data type", func(t *testing.T) {
		testDLS.Push(10)
		testDLS.Push(20)

		t.Run("Doubly Linked Stack Size", func(t *testing.T) {
			if testDLS.Size() != 2 {
				t.Errorf("Doubly Linked Stack Size is incorrect, got: %d, want: %d.", testDLS.Size(), 2)
			}
		})

		t.Run("Doubly Linked Stack Push", func(t *testing.T) {
			expected := []int{20, 10}

			if !reflect.DeepEqual(testDLS.Display(), expected) {
				t.Errorf("expected: %v, got: %v", expected, testDLS.Display())
			}
		})

		poppedItem, _ := testDLS.Pop()

		t.Run("Doubly Linked Stack Pop", func(t *testing.T) {
			if poppedItem != 20 {
				t.Errorf("expected: %v, got: %v", 20, poppedItem)
			}
		})

		t.Run("Doubly Linked Stack Peek", func(t *testing.T) {
			peekedItem, _ := testDLS.Peek()

			if peekedItem != 10 {
				t.Errorf("expected: %v, got: %v", 10, peekedItem)
			}
		})

		_, _ = testDLS.Pop()

		t.Run("Doubly Linked Stack IsEmpty True", func(t *testing.T) {
			if !testDLS.IsEmpty() {
				t.Errorf("expected: %v, got: %v", true, testDLS.IsEmpty())
			}
		})

		testDLS.Push(30)

		t.Run("Doubly Linked Stack IsEmpty False", func(t *testing.T) {
			if testDLS.IsEmpty() {
				t.Errorf("expected: %v, got: %v", false, testDLS.IsEmpty())
			}
		})

		t.Run("Doubly Linked Stack Display", func(t *testing.T) {
			expected := []int{30}
			if !reflect.DeepEqual(testDLS.Display(), expected) {
				t.Errorf("expected: %v, got: %v", expected, testDLS.Display())
			}
		})

		t.Run("Doubly Linked Stack Clear", func(t *testing.T) {
			expected := NewDoublyLinkedStack[int]()
			testDLS.Clear()
			if !reflect.DeepEqual(testDLS, expected) {
				t.Errorf("expected: %v, got: %v", expected, testDLS.Display())
			}
		})

		t.Run("Doubly Linked Stack Custom Error ", func(t *testing.T) {
			expectedErrorMsg := "Stack is empty. Zero value of the data type returned"

			_, err := testDLS.Pop()

			if !ErrorContains(err, expectedErrorMsg) {
				t.Errorf("expected: %v, got: %v", expectedErrorMsg, err)
			}

			_, err = testDLS.Peek()

			if !ErrorContains(err, expectedErrorMsg) {
				t.Errorf("expected: %v, got: %v", expectedErrorMsg, err)
			}
		})

	})
}
