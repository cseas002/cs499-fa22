// package main
package wordutil

import (
	// "fmt"
	"strings"
)

// Finds all occurrences of each word in a string.
//
// Returns a map that stores each unique word in the string s as the key and
// a slice that contains the index of each occurence of the word in the input
// string as the corresponding value.
// Matching is case insensitive, e.g. "Orange" and "orange" is considered the
// same word.
func WordIndexAll(s string) map[string][]int {
	// TODO: implement me
	s = strings.ToLower(s)
	stringMap := make(map[string][]int)
	stringList := strings.Fields(s)
	for _, str := range stringList {
		str = strings.ToLower(str) // make it to lowercase
		index := strings.Index(s, str)
		listLength := len(stringMap[str]) // the list in stringMap[str] length
		if listLength > 0 { // if there's already another element
			previousIndex := stringMap[str][listLength - 1]
			// fmt.Println(str, previousIndex, s[previousIndex:])
			index = strings.Index(s[previousIndex + 1:], str) + 1 // search on the right of that string
			index += previousIndex // the index is the previous index + 1 + the index on the right of that string
		}
		stringMap[str] = append(stringMap[str], index) // add to that key the index
	}
	// HINT: You may find the `strings.Index` and `strings.ToLower` functions helpful
	return stringMap
}

// func main() {
// 	fmt.Println(WordIndexAll("asdnas adhashd  ahdfuiahah sdhas art arq aa aa  art asd"))
// }
