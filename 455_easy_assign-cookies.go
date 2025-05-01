func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)
	ans := 0
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(s); j++ {
			if g[i] <= s[j] {
				ans++
                s[j] = 0
				break
			}
		}
	}
	return ans
}