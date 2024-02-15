package main

import (
	"strconv"
	"unicode"
)

//map the string with key = key index then to decode string turn string into list spli
//use a char as a deliminater and then strip that when presenting the stirng - this will
//act as a tell as to where the string was split to begin with

func encode(strs []string) (rs string) {

	ss := ""

	for _, s := range strs {
		delim := strconv.Itoa(len(s)) + "!" + s
		ss += delim
	}
	return ss

}

func decode(inp string) []string {
	str := make([]string, 0)

	nextWord := 0

	for _, char := range inp {
		switch {
		case unicode.IsDigit(char):
			nextWord = int(char)
		}
	}
}
