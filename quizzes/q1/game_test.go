package q1_test

import (
	"testing"

	"puzzle/quizzes/q1"
)

func Test_Game_IsDraw(t *testing.T) {
	draws := []bool{
		q1.Game(0, 0, 0).IsDraw(),
		q1.Game(10, 10, 0).IsDraw(),
		q1.Game(0, 20, 20).IsDraw(),
		q1.Game(30, 0, 30).IsDraw(),
	}
	for idx, isDraw := range draws {
		if !isDraw {
			t.Errorf("line %v is not draw", idx)
		}
	}
}

func Test_Game_IsNotDraw(t *testing.T) {
	draws := []bool{
		q1.Game(0, 1, 2).IsDraw(),
		q1.Game(11, 12, 13).IsDraw(),
	}
	for idx, isDraw := range draws {
		if isDraw {
			t.Errorf("line %v is draw", idx)
		}
	}
}
