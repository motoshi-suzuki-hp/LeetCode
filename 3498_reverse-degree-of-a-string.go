func reverseDegree(s string) int {
    ans := 0
    for i, c := range s {
        indexInReversedAlphabet := 27 - (int(c) - 96)
        ans += (i+1) * indexInReversedAlphabet
    }
    return ans
}
