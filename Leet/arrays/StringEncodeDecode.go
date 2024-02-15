//NOT FINISHED

// package main

// import (
// 	"fmt"
// 	"reflect"
// 	"strconv"
// )

// //map the string with key = key index then to decode string turn string into list spli
// //use a char as a deliminater and then strip that when presenting the stirng - this will
// //act as a tell as to where the string was split to begin with

// func encode(strs []string) (rs string) {

// 	ss := ""

// 	for _, s := range strs {
// 		delim := strconv.Itoa(len(s)) + "!" + s
// 		ss += delim
// 	}
// 	return ss

// }

// func decode(inp string) []string {
// 	str := make([]string, 0)

// 	var newStr = ""

// 	for _, char := range inp {
// 		if reflect.TypeOf(int(char)) == reflect.TypeOf(int(0)) {
// 			newStr = inp[0 : int(char)+1]
// 			str = append(str, newStr)
// 		}
// 		break
// 	}

// 	return str

// }

// func main() {
// 	str := "4!dank4!dank4!dank"
// 	newStr := decode(str)

// 	fmt.Println(newStr)
// }
