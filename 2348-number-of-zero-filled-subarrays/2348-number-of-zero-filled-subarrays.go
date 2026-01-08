func zeroFilledSubarray(nums []int) int64 {
 	zeroCount := []int{}
	cnt := 0
	for _, v := range nums {
		if v != 0 {
			if cnt > 0 {
				zeroCount = append(zeroCount, cnt)
			}
			cnt = 0
			continue
		}
		cnt++
	}

	if cnt > 0 {
		zeroCount = append(zeroCount, cnt)
	}

	sum := 0
	for _, v := range zeroCount {
		sum += v * (v + 1) / 2
	}

	return int64(sum)
}