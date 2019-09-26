package q9_test

import (
	"fmt"
	"testing"

	"puzzle/quizzes/q9"
)

func Test_Transform(t *testing.T) {

	r := q9.Numeral("122", 3).IsNarcissist()
	t.Logf("value %v", r)
}

func Test_Solve(t *testing.T) {
	for k := 8; k < 8*8*8*8*8; k++ {
		v := fmt.Sprintf("%v", k)
		n := q9.Numeral(v, 10).ChangeBase(8)
		if n.IsNarcissist() {
			t.Logf("%v", n.Value())
		}
	}
}

func Test_IsNarcissist(t *testing.T) {
	n1 := q9.Numeral("153", 10).IsNarcissist()
	if !n1 {
		t.Errorf("153(10) is narcissist")
	}
	n2 := q9.Numeral("12", 3).IsNarcissist()
	if !n2 {
		t.Errorf("12(3) is narcissist")
	}
	n3 := q9.Numeral("22", 3).IsNarcissist()
	if !n3 {
		t.Errorf("22(3) is narcissist")
	}

}
