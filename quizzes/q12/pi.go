package q12

import (
	"math"
)

// IPi :
type IPi interface {
	Value() int64
	N() int64
}

// Pi :
func Pi(n int64) IPi { return &pi{n} }

type pi struct{ n int64 }

func (p *pi) N() int64     { return p.n }
func (p *pi) Value() int64 { return int64(math.Pi * math.Pow(10, float64(p.n))) }
