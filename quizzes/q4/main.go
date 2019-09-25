package q4

// Main : solver
func Main() int {
	cnt := 0
	for h := 0; h < 24; h++ {
		for m := 0; m < 60; m++ {
			for s := 0; s < 60; s++ {
				timer, _ := Timer(h, m, s)
				if timer.Count() == 30 {
					cnt++
				}
			}
		}
	}
	return cnt
}
