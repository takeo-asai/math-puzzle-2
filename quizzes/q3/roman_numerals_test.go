package q3_test

import (
	"testing"

	"puzzle/quizzes/q3"
)

func Test_Roman(t *testing.T) {
	r0 := q3.RomanNumerals(7).Value()
	if r0 != "VII" {
		t.Errorf("roman numerals(7) must be VII, but %v", r0)
	}

	r1 := q3.RomanNumerals(46).Value()
	if r1 != "XLVI" {
		t.Errorf("roman numerals(46) must be XLVI, but %v", r1)
	}

	r3 := q3.RomanNumerals(3888).Value()
	if r3 != "MMMDCCCLXXXVIII" {
		t.Errorf("roman numerals(3888) must be MMMDCCCLXXXVIII, but %v", r3)
	}

	r4 := q3.RomanNumerals(3000).Value()
	if r4 != "MMM" {
		t.Errorf("roman numerals(3000) must be MMM, but %v", r4)
	}
}
