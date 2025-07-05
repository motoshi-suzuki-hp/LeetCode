func countConsistentStrings(allowed string, words []string) int {
    ans := 0
    for _, word := range words {
        count := 0
        for _, w := range word {
            if strings.Contains(allowed, string(w)) {
                count++
            }
        }
        if count == len(word) {
            ans++
        }
    }
    return ans
}
