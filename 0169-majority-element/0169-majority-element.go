func majorityElement(nums []int) int {
    m := make(map[int]int)

    for _, n := range nums {
        m[n] += 1
        if m[n] > len(nums) / 2 {
            return n
        }
    }

    return 0
}