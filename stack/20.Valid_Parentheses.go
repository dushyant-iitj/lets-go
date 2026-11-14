package stack

import "fmt"

func isValid(s string) bool {
	parenthesesTable := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}
	stack := []rune{}

	for _, value := range s {
		if v, exists := parenthesesTable[value]; exists {
			// closing bracket found but stack empty, return false
			if len(stack) == 0 {
				return false
			}

			// if opening parentheses just before then pop
			if stack[len(stack)-1] == v {
				stack = stack[:len(stack)-1] // pop
			} else {
				return false
			}
		} else {
			stack = append(stack, value) // push
		}
	}

	// return true if stack empty
	return len(stack) == 0
}

func ValidParentheses() {
	var s string
	// Example 1:
	// Input:
	// s := "()"
	// Output: true
	s = "()"
	fmt.Println(s)
	fmt.Println(isValid(s))

	// Example 2:
	// Input:
	// s := "()[]{}"
	// Output: true
	s = "()[]{}"
	fmt.Println(s)
	fmt.Println(isValid(s))

	// Example 3:
	// Input:
	// s := "(]"
	// Output: false
	s = "(]"
	fmt.Println(s)
	fmt.Println(isValid(s))

	// Example 4:
	// Input:
	// s := "([])"
	// Output: true
	s = "([])"
	fmt.Println(s)
	fmt.Println(isValid(s))

	// Example 5:
	// Input:
	// s := "([)]"
	// Output: false
	s = "([)]"
	fmt.Println(s)
	fmt.Println(isValid(s))
}
