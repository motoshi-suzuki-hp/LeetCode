func findThePrefixCommonArray(A []int, B []int) []int {
	var indexArray []int
	for i, a := range A {
		for j, b := range B {
			if a == b {
				indexArray = append(indexArray, max(i, j))
				break
			}
		}
	}

	var ans = make([]int, len(A))
	for _, index := range indexArray {
        ans[index]++
	}
    for i, a := range ans {
        if (i == len(ans)-1) {
            break
        }

        ans[i+1] += a
    }

    return ans
}
