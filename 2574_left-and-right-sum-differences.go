func leftRightDifference(nums []int) []int {
    ans := make([]int, 0, len(nums))
    for i := range nums {
        if i == 0 {
            ans = append(ans, absSubtract(nil, nums[1:]))
        } else if i == len(nums)-1 {
            ans = append(ans, absSubtract(nums[:len(nums)-1], nil))
        } else {
            ans = append(ans, absSubtract(nums[:i], nums[i+1:]))
        }
    }
    return ans
}

func absSubtract(a, b []int) int {
    sum := 0
    for _, u := range a {
        sum += u
    }
    for _, v := range b {
        sum -= v
    }
    if sum < 0 {
        return -sum
    }
    return sum
}
