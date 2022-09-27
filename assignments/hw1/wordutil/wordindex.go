package wordutil

import (
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
	s = strings.ToLower(s) // make the string s to lowercase
	stringMap := make(map[string]int) // initializing the map with string keys and int values
	stringList := strings.Fields(s)   // seperating the input string to a list of strings
	for _, str := range stringList {  // for every string in the list
		str = strings.ToLower(str) // make it to lowercase
		stringMap[str] = strings.Index(s, str) /* find the index of the first substring str in
		the string s and add the index of the string as the value with the string str as the key */
	}
	return stringMap
}
