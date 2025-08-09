func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
    ans := 0
    for _, h := range hours {
        if h < target {
            continue
        }
        ans++
    }

    return ans
}
