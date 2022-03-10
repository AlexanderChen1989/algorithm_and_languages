package main

import "sort"

func maxSlidingWindow(nums []int, k int) []int {
	res := []int{}

	q := []int{}
	l := 0

	for r := range nums {
		for len(q) > 0 && nums[q[len(q)-1]] <= nums[r] {
			q = q[:len(q)-1]
		}
		q = append(q, r)
		for len(q) > 0 && q[0] < l {
			q = q[1:]
		}
		if r < k-1 {
			continue
		}
		if r-l+1 == k {
			res = append(res, nums[q[0]])
		}
		l++
	}

	return res
}
