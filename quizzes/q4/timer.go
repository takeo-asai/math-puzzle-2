package q4

import (
	"errors"
)

// ITimer :
type ITimer interface{ Count() int }

// Timer : constructor for ITimer
func Timer(hours, minutes, seconds int) (ITimer, error) {
	if hours >= 24 || hours < 0 {
		return nil, errors.New("hours must be '0<=hours<24'")
	}
	if minutes >= 60 || minutes < 0 {
		return nil, errors.New("minutes must be '0<=minutes<60'")
	}
	if seconds >= 60 || seconds < 0 {
		return nil, errors.New("seconds must be '0<=seconds<60'")
	}
	return &timer{hours, minutes, seconds}, nil
}

type timer struct {
	hours   int
	minutes int
	seconds int
}

func (t *timer) Count() int {
	return t.countDigit(t.hours/10) + t.countDigit(t.hours%10) + t.countDigit(t.minutes/10) + t.countDigit(t.minutes%10) + t.countDigit(t.seconds/10) + t.countDigit(t.seconds%10)
}

func (t *timer) countDigit(n int) int {
	if n >= 10 || n < 0 {
		panic("countDigit(n): n must be 0<=n<=9")
	}
	switch n {
	case 9:
		return 6
	case 8:
		return 7
	case 7:
		return 3
	case 6:
		return 6
	case 5:
		return 5
	case 4:
		return 4
	case 3:
		return 5
	case 2:
		return 5
	case 1:
		return 2
	}
	return 6 // case 0
}
