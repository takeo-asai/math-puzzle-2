package q1_test

import (
	"puzzle/quizzes/q1"

	"testing"
)

func Test_case4(t *testing.T) {
	tbl := q1.GameTable(4)
	if tbl.Count() != 12 {
		t.Error("when number of players is 4, then result is 12")
	}
}
