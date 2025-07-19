func countMaxOrSubsets(nums []int) int {
    var ans int

    var bitwiseOR int
    for _, n := range nums {
        bitwiseOR |= n
    }

    size := len(nums)
    for i := 0; i < (1 << uint(size)); i++ {
        var subset int
        for j := 0; j < size; j++ {
            if i&(1<<uint(j)) != 0 {
                subset |= nums[j]
            }
        }

        if subset == bitwiseOR {
            ans++
        }
    }
    return ans
}
