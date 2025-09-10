func majorityElement(nums []int) int {
    majority, cnt := nums[0], 1

    for i := 1; i < len(nums); i++ {
        if cnt == 0 {
            majority = nums[i]
        }

        if nums[i] == majority {
            cnt++
        } else {
            cnt--
        }
    }

    return majority
}