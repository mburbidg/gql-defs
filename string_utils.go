package gqldefs

import "unicode/utf8"

func FirstN(s string, n int) string {
	i := 0
	for j := range s {
		if i == n {
			return s[:j]
		}
		i++
	}
	return s
}

func LastN(s string, n int) string {
	j := len(s)
	for i := 0; i < 3 && j > 0; i++ {
		_, size := utf8.DecodeLastRuneInString(s[:j])
		j -= size
	}
	return s[j:]
}
