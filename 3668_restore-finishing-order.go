func recoverOrder(order []int, friends []int) []int {
    var ans []int
    for _, o := range order {
        if slices.Contains(friends, o) {
            ans = append(ans, o)
        }
    }

    return ans
}
