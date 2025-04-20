func intToRoman(num int) string {
    type pair struct {
		value  int
		symbol string
	}

	pairs := []pair{
		{1000, "M"}, {900, "CM"},
		{500, "D"}, {400, "CD"},
		{100, "C"}, {90, "XC"},
		{50, "L"}, {40, "XL"},
		{10, "X"}, {9, "IX"},
		{5, "V"}, {4, "IV"},
		{1, "I"},
	}

	var ans string
	for _, p := range pairs {
		for num >= p.value {
			ans += p.symbol
			num -= p.value
		}
	}
	return ans
}