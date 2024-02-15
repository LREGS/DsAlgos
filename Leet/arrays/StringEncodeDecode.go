package main

import "strings"

//map the string with key = key index then to decode string turn string into list spli
//use a char as a deliminater and then strip that when presenting the stirng - this will
//act as a tell as to where the string was split to begin with

func encode(strs []string) (rs string) {

	ss := ""

	for _, s := range strs {
		delim := s + "!"
		ss += delim
	}
	return ss

}

func decode(inp string) []string {
	result := strings.Split(inp, "!")
	return result
}


func main(){
	
}