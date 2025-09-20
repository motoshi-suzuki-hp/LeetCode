// R: 82
// L: 76

func balancedStringSplit(s string) int {
    ans := 0

    rlMap := map[rune]int{
        'R': 0,
        'L': 0,
    }

    for _, v := range s {
        rlMap[v]++
        if rlMap['R'] == rlMap['L'] {
            ans++
            rlMap['R'] = 0
            rlMap['L'] = 0
        }
    }

    return ans
}
