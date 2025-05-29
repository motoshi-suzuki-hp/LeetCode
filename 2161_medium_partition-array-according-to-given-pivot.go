func pivotArray(nums []int, pivot int) []int {
    var left, medium, right []int
    for _, n := range nums {
        if (n < pivot) {
            left = append(left, n)
        } else if (pivot < n) {
            right = append(right, n)
        } else {
            medium = append(medium, n)
        }
    }
    return append(append(left, medium...), right...)
}