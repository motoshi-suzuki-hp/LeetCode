func decode(encoded []int, first int) []int {
    var ans []int
    ans = append(ans, first)
    for _, e := range encoded {
        first ^= e
        ans = append(ans, first)
    }
    return ans
}
