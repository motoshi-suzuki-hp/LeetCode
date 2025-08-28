func groupThePeople(groupSizes []int) [][]int {
    groupMap := make(map[int][]int)
    var ans [][]int
    for i, g := range groupSizes {
        groupMap[g] = append(groupMap[g], i)
        if len(groupMap[g]) == g {
            ans = append(ans, groupMap[g])
            groupMap[g] = []int{}
        }
    }
    return ans
}
