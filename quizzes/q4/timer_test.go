package q4_test

import (
	"testing"

	"puzzle/quizzes/q4"
)

func Test_check(t *testing.T) {
	timer, _ := q4.Timer(12, 34, 56)
	if timer.Count() != 27 {
		t.Errorf("12:34:56 is 27")
	}
}
