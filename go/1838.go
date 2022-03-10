package main

import (
	"sort"
)

func maxFrequency(nums []int, k int) int {
	res := 0

	sum := 0
	sort.Ints(nums)


	l := 0
	for r := range nums {
		sum += nums[r]
		if (r-l+1)*nums[r]-sum > k {
			sum -= nums[l]
			l++
		}
		d := r - l + 1
		if d > res {
			res = d
		}
	}

	return res
}