func minMovesToSeat(seats []int, students []int) int {
    var diff, ans int
    
    sort.Ints(seats)
    sort.Ints(students)

    for i := 0; i < len(seats); i++ {
        diff =  seats[i] - students[i]
        if diff < 0 {
            ans -= diff
        } else {
            ans += diff
        }
    }

    return ans
}
