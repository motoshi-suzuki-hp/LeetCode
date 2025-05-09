func numJewelsInStones(jewels string, stones string) int {
	var ans int
	for _, j := range jewels {
		for _, s := range stones {
			if s == j {
				ans++
			}
		}
	}
	return ans
}
