func maximumWealth(accounts [][]int) int {
    ans := 0
    for _, account := range accounts {
        sum := 0
        for _, a := range account {
            sum += a
        }

        if ans < sum {
            ans = sum
        }
    }
    return ans
}
