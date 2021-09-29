package question

type Row struct {
	char rune
	count int
}

func rowsToStr(rows []Row) string {
	chars := make([]rune, len(rows))
	for i, row := range rows {
		chars[i] = row.char
	}
	return string(chars)
}

func reverse(str string) string {
	res := ""
	for _, c := range str {
		res = string(c) + res
	}
	return res
}

func IsPalindromePermutation(str string) (bool, []string) {

	charMap := make(map[rune]int)

	for _, c := range str {
		count, existed := charMap[c]
		if existed {
			charMap[c] = count + 1
		} else {
			charMap[c] = 1
		}
	}
	rows := make([]Row, len(charMap))
	i := 0
	for k, v := range charMap {
		rows[i] = Row{ char: k, count: v }
		i++
	}

	evens := []Row{}
	odds := []Row{}

	for _, row := range rows {
		if (row.count % 2) == 0 {
			evens = append(evens, row)
		} else if row.count > 2 {
			char := row.char
			count := row.count
			evens = append(evens, Row{ char: char, count: count - 1 })
			odds = append(evens, Row{ char: char, count: 1 })
		} else {
			odds = append(odds, row)
		}
	}

	oddCount := 0
	for _, row := range odds {
		if (row.char != ' ') {
			oddCount++
		}
	}
	if oddCount > 1 {
		return false, nil
	}
	side := rowsToStr(evens)
	mid := rowsToStr(odds)
	sidesArr := permutation(side)

	permutations := make([]string, len(sidesArr))

    for i, s := range sidesArr {
		p := s + mid + reverse(s)
		permutations[i] = p
    }

	return true, permutations
}
