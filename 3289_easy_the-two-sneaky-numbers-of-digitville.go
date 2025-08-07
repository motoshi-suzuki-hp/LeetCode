func getSneakyNumbers(nums []int) []int {
    var ans []int
    flg := false
    for i, n := range nums {
        for _, a := range ans {
            if a == n {
                flg = true
                break
            }
        }
        if !flg {
            for _, m := range nums[i+1:] {
                if n == m {
                    ans = append(ans, n)
                    break
                }
            }
        }
        flg = false
    }
    return ans
}
