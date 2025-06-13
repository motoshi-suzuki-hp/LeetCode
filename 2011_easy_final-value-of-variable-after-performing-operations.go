func finalValueAfterOperations(operations []string) int {
    ans := 0
    for _, x := range operations {
        if x == "++X" || x == "X++" {
            ans++
        } else {
            ans--
        }
    }
    return ans
}
