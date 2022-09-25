package wordutil

import (
	// "fmt"
	"strings"
)

// Finds first occurrence of each word in a string.
//
// Returns a map that stores each unique word in the string s as the key and
// the index of the first occurence of the word in the input string as the
// corresponding value.
// Matching is case insensitive, e.g. "Orange" and "orange" is considered the
// same word.
func WordIndex(s string) map[string]int {
	// TODO: implement me
	stringMap := make(map[string]int)
	stringList := strings.Fields(s)
	for _, str := range stringList {
		str = strings.ToLower(str)              // make it to lowercase
		// if _, found := stringMap[str]; !found { // if there's not an entry in stringMap[str] (first time it's found)
		// 	stringMap[str] = strings.Index(s, str) // add to that key the index
		// }
		stringMap[str] = strings.Index(s, str)
	}
	// HINT: You may find the `strings.Index` and `strings.ToLower` functions helpful
	return stringMap
}

// func main() {
// 	fmt.Println(WordIndex("testing aaa  asd  aaa alow?"))
// }
