package q13_test

import (
	"testing"

	"puzzle/quizzes/q13"
)

func Test_cards(t *testing.T) {
	cards := q13.Cards([]int{1, 2, 3, 4})

	b2 := []q13.ICards{}
	for _, v := range cards.Before() {
		t.Logf("n=1, %v", v.Values())
		b2 = append(b2, v)
	}

	b3 := []q13.ICards{}
	for _, v := range b2 {
		b3 = append(b3, v.Before()...)
	}
	for _, v := range b3 {
		t.Logf("n=2 %v\n", v)
	}

	b4 := []q13.ICards{}
	for _, v := range b3 {
		b4 = append(b4, v.Before()...)
	}
	for _, v := range b4 {
		t.Logf("n=3 %v\n", v)
	}
}
