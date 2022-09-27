package wordutil

import (
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
	s = strings.ToLower(s)              // make the string s to lowercase
	stringMap := make(map[string][]int) // initializing the map with string keys and int values
	stringList := strings.Fields(s)     // seperating the input string to a list of strings
	for _, str := range stringList {    // for every string in the list
		str = strings.ToLower(str)        // make it to lowercase
		index := strings.Index(s, str)    // find the index of the first substring str in the string s
		listLength := len(stringMap[str]) // the list in stringMap[str] length
		if listLength > 0 {               // if there's already another element
			previousIndex := stringMap[str][listLength-1]       // get the index of the last element
			index = strings.Index(s[previousIndex+1:], str) + 1 // search on the right of that string
			index += previousIndex                              // the index is the previous index + 1 + the index on the right of that string
		}
		stringMap[str] = append(stringMap[str], index) // add to the key = the substring, the index
	}
	return stringMap
}
