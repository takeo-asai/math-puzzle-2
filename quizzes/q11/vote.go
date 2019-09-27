package q11

// Vote :
func Vote(n int) int64 {
	if n <= 2 {
		return 1
	}
	cnt := 1 + Vote(n-1)
	for i := 2; i <= n-1; i++ { // 下位i国が並ぶ=2以上
		cnt += Vote(i) * Vote(n-1)
	}
	return cnt
}
