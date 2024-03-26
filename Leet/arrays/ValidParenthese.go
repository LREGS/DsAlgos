package main

//solved an entirely different problem / read it wrong - use a stack

func isValid(s string) bool {

	answer := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}

	if len(s)%2 == 0 {

		i := 0
		for i < (len(s) - 1) {
			if string(s[i+1]) != answer[string(s[i])] {
				return false
			}
			i = i + 2
		}

	} else {
		return false
	}
	return true

}
