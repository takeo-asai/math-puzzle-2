package q10_test

import (
	"testing"

	"puzzle/quizzes/q10"
)

func Test_Solve(t *testing.T) {
	result, _, _, _ := q10.Adams(289, q10.Range(1, 1308265), q10.Prefectures)
	t.Logf("result integer is %v\n", result)
	for _, n := range q10.Prefectures {
		t.Logf("%v => %v\n", n, q10.Ceil(n, result))
	}
}
