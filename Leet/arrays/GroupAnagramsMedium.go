// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func groupAnagrams(strs []string) [][]string {

// 	anagramMap := make(map[string][]string)

// 	for _, str := range strs {
// 		sortedString := SortString(str)
// 		anagramMap[sortedString] = append(anagramMap[sortedString], str)
// 	}

// 	response := make([][]string, 0)
// 	for key := range anagramMap {
// 		response = append(response, anagramMap[key])
// 	}

// 	return response
// }

// func SortString(input string) string {
// 	runes := []rune(input)

// 	sort.Slice(runes, func(i, j int) bool {
// 		return runes[i] < runes[j]
// 	})

// 	sortedString := string(runes)

// 	return sortedString
// }

// func main() {
// 	anagrams := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
// 	sorted := groupAnagrams(anagrams)
// 	fmt.Println(sorted)
// }
