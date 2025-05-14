func addDigits(num int) int {
    var l int
    n := num
    for n > 9 {
        ns := strconv.Itoa(n)
        ns_list := strings.Split(ns, "")
        l = len(ns_list)
        n = 0
        for i := 0; i < l; i++ {
            ni , _ := strconv.Atoi(ns_list[i])
            n += ni
        }
    }
    return n
}