package main

/*

You are given a string s consisting of the following characters: '(', ')', '{', '}', '[' and ']'.

The input string s is valid if and only if:

Every open bracket is closed by the same type of close bracket.
Open brackets are closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.
Return true if s is a valid string, and false otherwise.


vampire
*/

type Stack struct {
	stack []string
}

func (s *Stack) Add(sToAdd string) {
	s.stack = append(s.stack, sToAdd)
}

func (s *Stack) Pop() {
	if len(s.stack) > 0 {
		s.stack = s.stack[:len(s.stack)-1]
	}
}

func (s *Stack) Peek() string {
	if len(s.stack) > 0 {
		return s.stack[len(s.stack)-1]
	} else {
		return " "
	}
}

func isValid(ts string) bool {

	// will return true if s contains valid parentheses

	// Loop through
	// add items to the stack
	// compared item to stack, if correct, remove item from stack
	// else, continue
	//if stack empty parenthese = valid

	if !(len(ts) > 1) {
		return false
	}

	currentlyOpen := Stack{stack: make([]string, 0)}

	// Open := map[rune]bool{'(': true, '{': true, '[': true}
	Close := map[rune]rune{')': '(', '}': '{', ']': '['}

	_, ok := Close[rune(ts[0])]
	if ok {
		return false
	}

	for _, c := range ts {
		_, ok := Close[c]
		if ok {
			if currentlyOpen.Peek() == string(Close[c]) {
				currentlyOpen.Pop()
			} else {
				return false
			}
		} else {
			currentlyOpen.Add(string(c))
		}
	}

	if len(currentlyOpen.stack) > 0 {
		return false
	} else {
		return true
	}

}
