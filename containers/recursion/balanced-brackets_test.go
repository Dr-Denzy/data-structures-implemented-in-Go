package recursion

import "testing"

func TestBalancedBrackets(t *testing.T) {
	tableTests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"empty", "", true},
		{"invalid or unsupported character", "/{", false},
		{"simple balanced", "[]", true},
		{"simple unbalanced", "][", false},
		{"nested balanced", "[[]]", true},
		{"nested unbalanced", "[[[]", false},
		{"deeply nested balanced", "[[[[]]]]", true},
		{"deeply nested unbalanced", "[[[[[]]]", false},
		{"simple sequence balanced", "[[][][]]", true},
		{"simple sequence unbalanced", "[[]]][][]]", false},
		{"complex sequence balanced", "[[[[]]][[[][]]][]]", true},
		{"complex sequence unbalanced", "[[[[]]][[[][]]][[]]", false},
	}

	t.Run("IsBalancedRecursive", func(t *testing.T) {
		for _, tt := range tableTests {
			result := IsBalancedRecursive(tt.input)

			if result != tt.expected {
				t.Errorf("IsBalancedRecursive(%s): expected %v, actual %v", tt.input, tt.expected, result)
			}
		}
	})

	t.Run("IsBalancedStack", func(t *testing.T) {
		for _, tt := range tableTests {
			result := IsBalancedStack(tt.input)

			if result != tt.expected {
				t.Errorf("IsBalancedStack(%s): expected %v, actual %v", tt.input, tt.expected, result)
			}
		}
	})
}
