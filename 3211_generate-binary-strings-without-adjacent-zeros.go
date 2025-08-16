func validStrings(n int) []string {
	ans := make([]string, 0)
	generateBinaryStrings(n, "", 0, &ans)
	return ans
}

func generateBinaryStrings(n int, bit string, i int, ans *[]string) {
	if i == n {
		*ans = append(*ans, bit)
		return
	}

	if i == 0 || bit[i-1] != '0' {
		generateBinaryStrings(n, bit+"0", i+1, ans)
	}

	generateBinaryStrings(n, bit+"1", i+1, ans)
}
