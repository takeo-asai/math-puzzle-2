package q10_test

import (
	"testing"

	"puzzle/quizzes/q10"
)

func Test_ceil(t *testing.T) {
	if q10.Ceil(10, 3) != 4 {
		t.Error("ceil(10, 3) is 4")
	}
	if q10.Ceil(10, 2) != 5 {
		t.Error("ceil(10, 2) is 5")
	}
}
