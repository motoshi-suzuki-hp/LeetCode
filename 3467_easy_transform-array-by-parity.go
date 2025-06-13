func transformArray(nums []int) []int {
    var even_slice, odd_slice []int
    for _, n := range nums {
        if (n % 2 == 0) {
            even_slice = append(even_slice, 0)
        } else {
            odd_slice = append(odd_slice, 1)
        }
    }

    ans := append(even_slice, odd_slice...)
    return ans
}
