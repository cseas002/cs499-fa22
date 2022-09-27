package wordutil

import (
	"strings"
)

// Counts occurrences of each word in a string.
//
// Returns a map that stores each unique word in the string s as the key and
// its count as the corresponding value.
// Matching is case insensitive, e.g. "Orange" and "orange" is considered the
// same word.
func WordCount(s string) map[string]int {
	s = strings.ToLower(s) // make the string s to lowercase
	stringMap := make(map[string]int) // initializing the map with string keys and int values
	stringList := strings.Fields(s) // seperating the input string to a list of strings
	for _, str := range stringList { // for every string in the list
		str = strings.ToLower(str)   // make it to lowercase
		stringMap[str]++ // add +1 to the map with the string as the key
	}
	return stringMap
}

