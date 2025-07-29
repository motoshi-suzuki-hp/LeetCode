func convertDateToBinary(date string) string {
    ymd := strings.Split(date, "-")
    var ymd_bin []string
    for _, x := range ymd {
        x_int, _ := strconv.ParseInt(x, 10, 64)
        x_bin := strconv.FormatInt(x_int, 2)
        x_bin_str := string(x_bin)
        ymd_bin = append(ymd_bin, x_bin_str)
    }
    return strings.Join(ymd_bin[:], "-")
}
