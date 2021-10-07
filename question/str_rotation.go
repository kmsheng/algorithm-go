package question

import "strings"

// 1.9
// String rotation: Assume you have a method
// isSubstring which checks if one word is a
// substring of another. Given two strings, s1 and s2,
// write code to check if s2 is a rotation of s1 using
// only one call to isSubstring.
// (e.g. "waterbottle" is a rotation of "erbottlewat")

func isSubstring(s1 string, s2 string) bool {
	return strings.Contains(s1, s2)
}

func StrRotation(s1 string, s2 string) bool {
	i := 0
	j := 0
	for i < len(s1) && j < len(s2) {
		c1 := s1[i]
		c2 := s2[j]
		if (c2 == c1) {
			j++
		}
		i++
	}
	s3 := s2[j + 1:]
	return isSubstring(s1, s3)
}
