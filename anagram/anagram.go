package anagram

import (
	"sort"
	"strings"
)

// Rune is a wrapper to slice of runes, implements sort.Interface
type Rune []rune

func (r Rune) Len() int           { return len(r) }
func (r Rune) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r Rune) Less(i, j int) bool { return r[i] < r[j] }

func (r Rune) Equal(a []rune) bool {
	if len(r) != len(a) {
		return false
	}

	for index, elem := range r {
		if a[index] != elem {
			return false
		}
	}

	return true
}

// stringToRuneSlice returns rune slice from input string
func stringToRuneSlice(s string) Rune {
	// initialize result slice with capacity of string
	r := make([]rune, 0, len(s))

	// fill rune slice
	for _, runeValue := range s {
		r = append(r, runeValue)
	}

	return r
}

// Anagram returns a boolean indicating whether a is an anagram of b.
// Works with any count of spaces
func Anagram(a string, b string) bool {
	// strings to lower case without any spaces
	as := strings.ToLower(strings.Replace(a, " ", "", -1))
	bs := strings.ToLower(strings.Replace(b, " ", "", -1))

	// rune slices from a string to lower case without any spaces
	arune := stringToRuneSlice(as)
	brune := stringToRuneSlice(bs)

	// if lengths of slices are not equal, then input strings are not anagrams
	if len(arune) != len(brune) {
		return false
	}

	// sort both slices
	sort.Sort(arune)
	sort.Sort(brune)

	// if two sorted slices are equal, then a and b are anagram
	return arune.Equal(brune)
}
