func getFinalState(nums []int, k int, multiplier int) []int {
    for i := 0; i < k; i++ {
        idx := 0
        min := nums[idx]
        for i, num := range nums {
            if i == 0 {
                continue
            }
            if num < min {
                idx = i
                min = num
            }
        }

        nums[idx] *= multiplier
    }

    return nums
}
