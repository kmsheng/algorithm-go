package question

import (
	"strconv"
)

// 1.6
// String Compression: Implement a method to perform basic
// string compression using the counts of repeated characters.
// For example, the string aabcccccaaa would become a2b1c5a3.
// If the "compressed" string would not become smaller than
// the original string, your method should return the original
// string. You can assume the string has only uppercase and
// lowercase letters (a - z)

func Compress(str string) string {
	type Row struct {
		char string
		count int
	}
	rows := make([]Row, 0)
	prev := ""
	var currRow Row


	for _, c := range str {
		char := string(c)
		if char == prev {
		  rows[len(rows) - 1].count++
		} else {
			currRow = Row{ char: char, count: 1 }
			rows = append(rows, currRow)
			prev = char
		}
	}
	res := ""
	for _, row := range rows {
		res += row.char + strconv.Itoa(row.count)
	}
	if len(res) < len(str) {
		return res
	}
	return str
}
