package main

import (
	"fmt"
)

func isValid(s string) bool {
	// Mapping of closing brackets to their corresponding opening brackets
	bracketMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	// Stack to keep track of opening brackets
	stack := []rune{}

	for _, char := range s {
		// If it's a closing bracket
		if matchingOpen, exists := bracketMap[char]; exists {
			// Check if the stack is empty or the top doesn't match
			if len(stack) == 0 || stack[len(stack)-1] != matchingOpen {
				return false
			}
			// Pop the stack
			stack = stack[:len(stack)-1]
		} else {
			// It's an opening bracket, push onto the stack
			stack = append(stack, char)
		}
	}

	// If the stack is empty, all brackets are matched
	return len(stack) == 0
}

func main() {
	// Test cases
	fmt.Println(isValid("()"))        // Output: true
	fmt.Println(isValid("()[]{}"))    // Output: true
	fmt.Println(isValid("(]"))        // Output: false
	fmt.Println(isValid("([)]"))      // Output: false
	fmt.Println(isValid("{[]}"))      // Output: true
}
