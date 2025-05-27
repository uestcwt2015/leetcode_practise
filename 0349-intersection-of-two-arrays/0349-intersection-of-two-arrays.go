func intersection(nums1 []int, nums2 []int) []int {
    m := map[int]bool{}
    result := []int{}
    for _, v := range nums1 {
        m[v] = true
    }
    
    for _, v := range nums2 {
        if m[v] {
            result = append(result, v)
            m[v] = false
        }
    }

    return result
}