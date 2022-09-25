package wordutil
// package main

import (
	// "fmt"
	"strings"
)

// Counts occurrences of each word in a string.
//
// Returns a map that stores each unique word in the string s as the key and
// its count as the corresponding value.
// Matching is case insensitive, e.g. "Orange" and "orange" is considered the
// same word.
func WordCount(s string) map[string]int {
	// TODO: implement me
	stringMap := make(map[string]int)
	stringList := strings.Fields(s)
	for _, str := range stringList {
		str = strings.ToLower(str)              // make it to lowercase
		stringMap[str]++ // add to that key the index
	}
	// HINT: You may find the `strings.Index` and `strings.ToLower` functions helpful
	return stringMap
}

// func main() {
// 	fmt.Println(WordCount("sadas asdafa asda a a aa a dsa"))
// }
