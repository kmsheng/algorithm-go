package question

func insert(arr []string, c rune) []string {
	length := len(arr)
	if length == 0 {
		return append(arr, string(c))
	}
	var res []string
	char := string(c)
	for i := 0; i < length; i++ {
		str := arr[i]
		for j := 0; j <= len(str); j++ {
			newStr := str[:j] + char + str[j:]
			res = append(res, newStr)
		}
	}
	return res
}

func permutation(str string) []string {
  // a                        f(1) = 1 * 1 = 1
  // ab ba                    f(2) = f(1) * 2 = 2
  // cab acb abc cba bca bac  f(3) = f(2) * 3 = 6
  //                          f(4) = f(3) * 4 = 24
  //                          f(n) = f(n - 1) * n
  //                          O(n!)
  var arr []string
  for _, c := range str {
	  arr = insert(arr, c)
  }
  return arr
}

// 1.2
// Check Permutation: Given two strings, write a method to
// decide if one is a permutation of the other.

func IsPermutation(s1 string, s2 string) bool {
	arr := permutation(s1)
	for _, s := range arr {
		if s2 == s {
			return true
		}
	}
	return false
}
