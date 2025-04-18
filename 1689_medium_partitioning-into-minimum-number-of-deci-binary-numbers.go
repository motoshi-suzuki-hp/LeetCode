func minPartitions(n string) int {
	ans := 0
	for _, digit := range strings.Split(n, "") {
		d, _ := strconv.Atoi(digit)
		if ans < d {
			ans = d
		}
		if ans == 9 {
			break
		}
	}
	return ans
}