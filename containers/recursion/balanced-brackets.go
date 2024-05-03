package recursion

import (
	"github.com/drdenzy/containers/stack"
	"strings"
	"text/scanner"
)

// IsBalancedRecursive uses recursion to check for balanced brackets
func IsBalancedRecursive(s string) bool {
	var scan scanner.Scanner
	var ch rune
	var isBalanced func() bool

	isBalanced = func() bool {
		// check if the current character in the input string is EOF
		if ch == scanner.EOF {
			return true
		}

		// if the current character is not an opening bracket, there is no point continuing; return false
		// otherwise, consume the character and proceed to scan the next character in the input string
		if ch != '[' {
			return false
		}

		// scan the next character in the input string
		ch = scan.Next()

		if ch == '[' {
			// call self recursively to check for nested balanced brackets
			if !isBalanced() {
				return false
			}
		}

		// check for a matching closing bracket if balanced nested brackets were found during the recursive call
		if ch != ']' {
			return false
		}

		ch = scan.Next()

		// this covers for when there exists sequence of balanced brackets in the input string
		if ch == '[' {
			return isBalanced()
		}

		return true
	}

	scan.Init(strings.NewReader(s))
	ch = scan.Next()
	return isBalanced() && ch == scanner.EOF
}

// IsBalancedStack uses stack data structure to check for balanced brackets
func IsBalancedStack(s string) bool {
	var scan scanner.Scanner
	var stk stack.LinkedStack[rune]

	scan.Init(strings.NewReader(s))

	for ch := scan.Next(); ch != scanner.EOF; ch = scan.Next() {
		switch ch {
		case '[':
			stk.Push(ch)
		case ']':
			if stk.IsEmpty() {
				// if stack is empty, it means there's a closing bracket without a corresponding opening one
				return false
			}
			// it pops (removes) the top element from the stack, assuming it matches the current closing bracket.
			_, _ = stk.Pop()
		default:
			// the current character is neither an opening nor a closing bracket
			return false

		}
	}
	// After processing all characters in the input string, if the stack is empty, it means all opening brackets have been matched with closing ones
	return stk.IsEmpty()
}
