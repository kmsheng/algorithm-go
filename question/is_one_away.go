package question

import (
	"math"
	"strings"
	"unicode/utf8"
)

func strLen(str string) float64 {
	return float64(utf8.RuneCountInString(str))
}

func getDiffChars(s1 string, s2 string) []rune {
	chars := make([]rune, 0)
	s1Len := int(strLen(s1))
	s2Len := int(strLen(s2))
	s1Chars := []rune(s1)
	s2Chars := []rune(s2)
	sameLen := s1Len == s2Len
	i := 0
	j := 0
	for (i < s1Len) && (j < s2Len) {
		equaled := s1Chars[i] == s2Chars[j]
		if (! equaled) {
			chars = append(chars, s2Chars[j])
		}
		if (equaled || sameLen) {
			i++
		}
		j++
	}
	return chars
}

// 1.4
// Palindrome Permutation: Given a string, write a function to check
// if it is a permutation of a palindrome. A palindrome is a word or
// phrase that is the same forwards and backwards. A permutation is
// a rearrangement of letters. The palindrome does not need to be limited
// to just dictionary words.
//
// EXAMPLE
// pale, ple -> true
// pales, pale -> true
// pale, bale -> true
// pale, bake -> false

func IsOneAway(str string) bool {
	arr := strings.Split(str, ",")
	s1 := strings.Trim(arr[0], " ")
	s2 := strings.Trim(arr[1], " ")

	s1Len := strLen(s1)
	s2Len := strLen(s2)

	if math.Abs(s1Len - s2Len) > 1 {
		return false
	}
	if (s1Len > s2Len) {
		s1, s2 = s2, s1
	}
	diffChars := getDiffChars(s1, s2)
	return len(diffChars) <= 1
}
