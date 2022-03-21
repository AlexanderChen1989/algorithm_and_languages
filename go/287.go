package main

func main() {
	numss := [][]int{{2, 5, 9, 6, 9, 3, 8, 9, 7, 1}}//, {1, 3, 4, 2, 2}, {3, 1, 3, 4, 2}}
	for _, nums := range numss {

		println(findDuplicate(nums))
	}
}

func findDuplicate(nums []int) int {
	s, f := 0, 0
	for s == 0 || s != f {
		s = nums[s]
		f = nums[nums[f]]
	}

	s = 0
	for s == 0 || s != f {
		s = nums[s]
		f = nums[f]
	}

	return s
}
