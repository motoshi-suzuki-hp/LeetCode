func scoreOfString(s string) int {
	ans := 0
	for i := 0; i < len(s)-1; i++ {
		s := int(s[i+1]) - int(s[i])
		if s < 0 {
			ans -= s
		} else {
			ans += s
		}
	}
	return ans
}