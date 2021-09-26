package question

import (
	"math"
)

func bitcount(n int) int {
	count := 0
	for n != 0 {
		count += (n & 1)
		n >>= 1
	}
	return count
}

// 1.1
// Is Unique: Implement an algorithm to determine
// if a string has all unique characters. What
// if you cannot use additional data structures ?

// use binary to represent each character
func IsUniq(str string) bool {
	length := len(str)
	base := 'a' - 1
	baseNum := 0
	for _, c := range str {
		n := float64(c - base)
		num := math.Pow(2, n - 1)
		baseNum = baseNum | int(num)
	}
	return bitcount(baseNum) == length
}
