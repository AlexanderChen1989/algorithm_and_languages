package main

import "fmt"

func main() {
	nums := []int{5, 4, 3, 6, 7, 8, 1, 0, 9}
	// nums = []int{}
	fmt.Println(nums)
	// qsort(nums, 0, len(nums) - 1);
	msort(nums, 0, len(nums)-1, make([]int, len(nums)))
	fmt.Println(nums)
}

func msort(nums []int, l, r int, tmp []int) {
	if l >= r {
		return
	}
	m := (l + r) / 2

	msort(nums, l, m, tmp)
	msort(nums, m+1, r, tmp)

	p, q := l, m+1

	i := l
	for ; p <= m && q <= r; i++ {
		if nums[p] <= nums[q] {
			tmp[i] = nums[p]
			p++
		} else {
			tmp[i] = nums[q]
			q++
		}
	}
	for p <= m {
		tmp[i] = nums[p]
		i++
		p++
	}
	for q <= r {
		tmp[i] = nums[q]
		i++
		q++
	}

	for i = l; i <= r; i++ {
		nums[i] = tmp[i]
	}
}

func qsort(nums []int, l, r int) {
	if l >= r {
		return
	}

	first, last := l, r
	key := nums[first]
	for first < last {
		for first < last && nums[last] >= key {
			last--
		}
		nums[first] = nums[last]
		for first < last && nums[first] <= key {
			first++
		}
		nums[last] = nums[first]
	}
	nums[first] = key

	qsort(nums, l, first-1)
	qsort(nums, first+1, r)
}
