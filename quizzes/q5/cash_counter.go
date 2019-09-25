package q5

import ()

// ICashCounter :
type ICashCounter interface {
	Count() int
}

// CashCounter : constructor for ICashCounter
func CashCounter(money int) ICashCounter {
	return &counter{money}
}

type counter struct {
	money int
}

func (c *counter) Count() int {
	rest := c.money
	cnt := 0

	// bills
	cnt += rest / 10000
	rest = rest - 10000*(rest/10000)
	cnt += rest / 5000
	rest = rest - 5000*(rest/5000)
	cnt += rest / 2000
	rest = rest - 2000*(rest/2000)
	cnt += rest / 1000
	rest = rest - 1000*(rest/1000)

	// coins
	cnt += rest / 500
	rest = rest - 500*(rest/500)
	cnt += rest / 100
	rest = rest - 100*(rest/100)
	cnt += rest / 50
	rest = rest - 50*(rest/50)
	cnt += rest / 10
	rest = rest - 10*(rest/10)
	cnt += rest / 5
	rest = rest - 5*(rest/5)
	cnt += rest

	return cnt
}
