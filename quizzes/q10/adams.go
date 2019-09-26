package q10

type IRange interface {
	Center() int
	Min() int
	Max() int
}

func Range(min, max int) IRange {
	return &rang{min, max}
}

type rang struct {
	min int
	max int
}

func (r rang) Center() int { return (r.min + r.max) / 2 }
func (r rang) Min() int    { return r.min }
func (r rang) Max() int    { return r.max }

func Adams(threshold int, r IRange, prefectures []int) (int, int, IRange, []int) {
	sum := 0
	target := r.Center()
	for _, n := range prefectures {
		sum += Ceil(n, target)
	}

	// goal
	if sum == threshold {
		return target, threshold, r, prefectures
	}

	if sum > threshold {
		return Adams(threshold, Range(r.Center(), r.Max()), prefectures)
	}
	return Adams(threshold, Range(r.Min(), r.Center()), prefectures)
}
