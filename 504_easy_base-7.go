func convertToBase7(num int) string {
    if num == 0 {
        return "0"
    }

    var q, r int
    var ans_slice []string
    minus_flg := false

    for num != 0 {
        q = num / 7
        r = num % 7
        if r < 0 {
            r = -r
            minus_flg = true
        }
        ans_slice = append(ans_slice, strconv.Itoa(r))
        num = q
    }

    for i, j := 0, len(ans_slice)-1; i < j; i, j = i+1, j-1 {
        ans_slice[i], ans_slice[j] = ans_slice[j], ans_slice[i]
    }

    var ans string

    if minus_flg {
        ans = "-" + strings.Join(ans_slice, "")
    } else {
        ans = strings.Join(ans_slice, "")
    }
    
    return ans
}