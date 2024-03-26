// package main

// import (
// 	"fmt"
// 	"strings"
// 	"unicode"
// )

// func isPalindrome(s string) bool {

// 	s = stripString(s)
// 	start := 0
// 	end := (len(s) - 1)

// 	if len(s) == 2 {
// 		fmt.Print("test")
// 		if s[0] != s[1] {
// 			return false
// 		}
// 		return true
// 	}

// 	strings.Map()

// 	for start < end {
// 		if strings.ToLower(string(s[start])) != strings.ToLower(string(s[end])) {
// 			return false
// 		}
// 		start = start + 1
// 		end = end - 1
// 		continue

// 	}
// 	return true
// }

// func stripString(s string) string {

// 	var result string

// 	for _, char := range s {
// 		if unicode.IsLetter(char) {
// 			result += string(char)
// 		} else {
// 			continue
// 		}
// 	}
// 	return result

// }

// func main() {
// 	string := "0a"
// 	if isPalindrome(string) {
// 		fmt.Print("true")
// 	}
// }

import "unicode"

func isPalindrome(s string) bool {

	f := func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}
		return unicode.ToLower(r)
	}

	s = strings.Map(f, s)

	start, end := 0, (len(s) - 1)

	for start < end {

		if s[start] != s[end] {
			return false
		}
		start = start + 1
		end = end - 1
	}
	return true
}

