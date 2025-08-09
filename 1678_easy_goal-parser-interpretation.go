func interpret(command string) string {
    ans := ""
    flg := false
    for _, c := range command {
        if c == 71 {
            ans += "G"
        } else if c == 40 {
            flg = true
            continue
        } else if c == 41 {
            if flg {
                ans += "o"
            } else {
                continue
            }
        } else if c == 97 {
            ans += "al"
            flg = false
        } else if c == 108 {
            continue
        }
    }
    return ans
}
