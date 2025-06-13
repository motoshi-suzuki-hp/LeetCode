func numIdenticalPairs(nums []int) int {
    l := len(nums)
    ans := 0
    for i := 0; i < l; i++ {
        for j := i+1; j < l; j++ {
            if nums[i] == nums[j] {
                ans++
            }
        }
    }
    return ans
}