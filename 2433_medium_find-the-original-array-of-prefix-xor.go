func findArray(pref []int) []int {
	var arr []int
	for i, v := range pref {
		if i == 0 {
			arr = append(arr, v)
			continue
		}
		arr = append(arr, v^pref[i-1])
	}
	return arr
}
