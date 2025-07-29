func largestLocal(grid [][]int) [][]int {
    n := len(grid)
	size := n - 2

    ans := make([][]int, size)
	for i := range ans {
		ans[i] = make([]int, size)
	}

    for i := 0; i < n-2; i++ {
        for j := 0; j < n-2; j++ {
            maxLocal := 0
            for r := i; r < i+3; r++ {
                for c := j; c < j+3; c++ {
                    if maxLocal < grid[r][c] {
                        maxLocal = grid[r][c]
                    }
                }
            }
            ans[i][j] = maxLocal
        }
    }

    return ans
}

