package q9

import (
	"fmt"
	"strconv"
	"strings"
)

// INumeral :
type INumeral interface {
	IsNarcissist() bool
	Value() string
	ChangeBase(base int) INumeral
}

// Numeral : constructor for INumeral
func Numeral(value string, base int) INumeral {
	return &numeral{value, base}
}

type numeral struct {
	value string
	base  int
}

func (n *numeral) ChangeBase(base int) INumeral {
	b10, _ := strconv.ParseInt(n.value, n.base, 0)
	n.value = strconv.FormatInt(b10, base)
	n.base = base
	return n
}

func (n *numeral) IsNarcissist() bool {
	var sum int64 = 0
	digits := len(strings.Split(n.value, ""))
	for _, r := range strings.Split(n.value, "") {
		v, _ := strconv.ParseInt(r, 10, 0)
		var w int64 = 1
		for i := 0; i < digits; i++ {
			w *= v
		}
		sum += w
	}

	b10, _ := strconv.ParseInt(n.value, n.base, 0)
	return sum == b10
}
func (n *numeral) Value() string {
	b10, _ := strconv.ParseInt(n.value, n.base, 0)
	b2 := strconv.FormatInt(b10, 2)
	b8 := strconv.FormatInt(b10, 8)
	return fmt.Sprintf("%v(%v) == %v(10), %v(8), %v(2)", n.value, n.base, b10, b8, b2)
}
