func findPermutationDifference(s string, t string) int {
    var ans, diff int
    for i, u := range s {
        for j, v := range t {
            if u == v {
                diff = i-j
                continue
            }
        }
        if diff < 0 {
            ans -= diff
        } else {
            ans += diff
        }
    }
    return ans
}
