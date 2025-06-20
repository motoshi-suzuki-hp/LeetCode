func subsetXORSum(nums []int) int {
    n := len(nums)
    var ans int

    for mask := 0; mask < (1 << n); mask++ {
        xor := 0

        for i := 0; i < n; i++ {
            if mask&(1<<i) != 0 {
                xor ^= nums[i]
            }
        }
        ans += xor
    }

    return ans
}
