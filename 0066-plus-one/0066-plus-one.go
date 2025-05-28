func plusOne(digits []int) []int {
    p, j := 1, len(digits)-1
    for j >= 0 {
        v := digits[j] + p
        if v >= 10 {
            p = 1
            digits[j] = v % 10
        } else {
            p = 0
            digits[j] = v
        }
        j--
    }

    if p == 1 {
        return append([]int{1}, digits...)
    }

    return digits
}