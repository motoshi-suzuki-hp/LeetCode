func removeElement(nums []int, val int) int {
	d := 0
	l := len(nums)
	for i := 0; i < l; i++ {

		if (nums[i-d] == val) {
			if (i == l-1) {
				nums = nums[:i-d]
                d++
				continue
			}
			nums = append(nums[:i-d], nums[i-d+1:]...)
			d++
		}
	}

	return l-d
}