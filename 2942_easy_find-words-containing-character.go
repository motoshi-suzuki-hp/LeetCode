func findWordsContaining(words []string, x byte) []int {
    var ans []int
    for i, w := range words {
        if (strings.Contains(w, string(x))) {
            ans = append(ans, i)
        }
    }
    return ans
}