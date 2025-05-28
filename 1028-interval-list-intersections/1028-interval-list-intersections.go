func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
    p1, p2 := 0, 0
    ans := make([][]int, 0)
    for p1 < len(firstList) && p2 < len(secondList) {
        s1, s2 := firstList[p1], secondList[p2]
        
        l, h := max(s1[0], s2[0]), min(s1[1], s2[1])
        if l <= h {
            ans = append(ans, []int{l, h})
        }

        if s1[1] <= s2[1] {
            p1++
        } else {
            p2++
        }
    }
    return ans
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}