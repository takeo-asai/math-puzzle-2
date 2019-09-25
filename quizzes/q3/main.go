package q3

import (
	"unicode/utf8"
)

// Main : q3 solver
func Main() int {
	cnt := 0
	for i := 0; i < 4000; i++ {
		r := RomanNumerals(i).Value()
		if utf8.RuneCountInString(r) == 12 {
			cnt++
		}
	}
	return cnt
}
