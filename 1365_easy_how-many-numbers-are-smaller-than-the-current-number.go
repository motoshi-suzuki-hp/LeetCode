func smallerNumbersThanCurrent(nums []int) []int {
    var s int
    var ans []int

    for i, u := range nums {
        s = 0
        for j, v := range nums {
            if i == j {
                continue
            }
            if u > v {
                s++
            }
        }
        ans = append(ans, s)
    }

    return ans
}
