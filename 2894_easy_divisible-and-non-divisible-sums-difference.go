func differenceOfSums(n int, m int) int {
	l := n / m
	num1 := n*(n+1)/2 - l*(l+1)/2*m
	num2 := l * (l + 1) / 2 * m
	return num1 - num2
}