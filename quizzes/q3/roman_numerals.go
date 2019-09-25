package q3

// Numerals :
type Numerals interface {
	Value() string
}

// RomanNumerals : constructor for RomanNumerals
func RomanNumerals(arabic int) Numerals {
	return &roman{arabic}
}

// roman : ローマ数字
type roman struct {
	arabic int
}

func (r *roman) Value() string {
	return r.fourthDigit(r.arabic) + r.thirdDigit(r.arabic) + r.secondDigit(r.arabic) + r.firstDigit(r.arabic)
}

func (r *roman) firstDigit(n int) string {
	remainder := n % 10
	switch remainder {
	case 9:
		return "IX"
	case 8:
		return "VIII"
	case 7:
		return "VII"
	case 6:
		return "VI"
	case 5:
		return "V"
	case 4:
		return "IV"
	case 3:
		return "III"
	case 2:
		return "II"
	case 1:
		return "I"
	}
	return "" // case 0
}

func (r *roman) secondDigit(n int) string {
	remainder := (n % 100) / 10
	switch remainder {
	case 9:
		return "XC"
	case 8:
		return "LXXX"
	case 7:
		return "LXX"
	case 6:
		return "LX"
	case 5:
		return "L"
	case 4:
		return "XL"
	case 3:
		return "XXX"
	case 2:
		return "XX"
	case 1:
		return "X"
	}
	return "" // case 0
}

func (r *roman) thirdDigit(n int) string {
	remainder := (n % 1000) / 100
	switch remainder {
	case 9:
		return "CM"
	case 8:
		return "DCCC"
	case 7:
		return "DCC"
	case 6:
		return "DC"
	case 5:
		return "D"
	case 4:
		return "CD"
	case 3:
		return "CCC"
	case 2:
		return "CC"
	case 1:
		return "C"
	}
	return "" // case 0
}

func (r *roman) fourthDigit(n int) string {
	if n >= 4000 {
		panic("n >= 4000 is not supported yet")
	}
	remainder := (n % 10000) / 1000
	switch remainder {
	case 3:
		return "MMM"
	case 2:
		return "MM"
	case 1:
		return "M"
	}
	return "" // case 0
}
