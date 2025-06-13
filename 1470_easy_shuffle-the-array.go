func shuffle(nums []int, n int) []int {
    var ans []int
    for i := range 2*n {
        if (i % 2 == 0) {
            ans = append(ans, nums[i/2])
        } else {
            ans = append(ans, nums[i/2+n])
        }
    }
    return ans
}
