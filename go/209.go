package main

import "math"

func minSubArrayLen(target int, nums []int) int {
	res := math.MaxInt
	sum := 0
	l := 0
	for r := range nums {
		sum += nums[r]
		if sum < target {
			continue
		}
		for sum >= target {
			res = min(res, r-l+1)
			sum -= nums[l]
			l++
		}
	}

	if res == math.MaxInt {
		return 0
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}