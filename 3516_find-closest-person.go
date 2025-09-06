func findClosest(x int, y int, z int) int {
    var gapXZ, gapYZ int
    if z-x < 0 {
        gapXZ = x-z
    } else {
        gapXZ = z-x
    }

    if z-y < 0 {
        gapYZ = y-z
    } else {
        gapYZ = z-y
    }

    if gapXZ == gapYZ {
        return 0
    } else if gapXZ < gapYZ {
        return 1
    } else {
        return 2
    }
}
