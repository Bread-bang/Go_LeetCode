func majorityElement(nums []int) int {
    var m map[int]int
    m = make(map[int]int)

    majority := 0
    maxCnt := 0
    for _, n := range nums {
        m[n] += 1
        if m[n] > maxCnt {
            maxCnt = m[n]
            majority = n
        }
    }

    return majority
}