func maxWidthOfVerticalArea(points [][]int) int {
    var xSlice []int
    for _, p := range points {
        xSlice = append(xSlice, p[0])
    }
    slices.Sort(xSlice)

    var ans int
    for i := 1; i < len(xSlice); i++ {
        if ans < xSlice[i] - xSlice[i-1] {
            ans = xSlice[i] - xSlice[i-1]
        }
    }

    return ans
}
