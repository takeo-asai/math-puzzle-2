package q13

import (
//	"fmt"
)

// ICards :
type ICards interface {
	Values() []int
	Before() []ICards
}

// Cards :
func Cards(values []int) ICards { return &cards{values} }

type cards struct{ values []int }

func (cs *cards) Values() []int { return cs.values }
func (cs *cards) Before() []ICards {
	length := len(cs.Values())
	beforeCards := []ICards{}
	for idx, v := range cs.Values() {
		if idx >= 1 && idx+1 == v {
			//			fmt.Printf("idx=%v, v=%v\n", idx, v)
			before := make([]int, length)
			copy(before, cs.Values())
			//			fmt.Printf("%v \n", before)
			for i := 0; i < (idx+1)/2; i++ {
				a := before[idx-i]
				before[idx-i] = before[i]
				before[i] = a
			}
			//			fmt.Printf("%v \n", before)
			//			fmt.Print("----\n")
			beforeCards = append(beforeCards, Cards(before))
		}
	}
	return beforeCards
}
