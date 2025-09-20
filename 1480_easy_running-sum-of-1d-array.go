func runningSum(nums []int) []int {
    var ans []int

    for i, v := range nums {
        if i == 0 {
            ans = append(ans, v)
        } else {
            ans = append(ans, ans[i-1]+v)
        }
    }

    return ans
}
