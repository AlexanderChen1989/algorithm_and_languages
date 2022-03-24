package main

func main() {
	nums := []int{2,5,4,2,4,5,3,1,2,4}
	println(maxSumMinProduct(nums))
}

func maxSumMinProduct(nums []int) int {
	sums := make([]int, len(nums))

	sums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sums[i] = nums[i] + sums[i-1]
	}
	stack := []Pair{}
	result := 0
	for idx, num := range nums {
		newIdx := idx
		if len(stack) == 0 {
			stack = append(stack, Pair{Idx: newIdx, Num: num})
			continue
		}
		for len(stack) > 0 {
			pair := stack[len(stack)-1]
			if pair.Num < num {
				break
			}
			if pair.Num == num {
				newIdx = pair.Idx
				break
			}
			if pair.Idx == 0 {
				result = max(result, pair.Num*sums[idx-1])
			} else {
				result = max(result, pair.Num*(sums[idx-1]-sums[pair.Idx-1]))
			}
			newIdx = pair.Idx
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, Pair{Idx: newIdx, Num: num})
	}

	for _, pair := range stack {
		if pair.Idx == 0 {
			result = max(result, pair.Num*sums[len(sums)-1])
			continue
		}

		result = max(result, pair.Num*(sums[len(sums)-1]-sums[pair.Idx-1]))
	}

    return result % (1000000000 + 7)
}

type Pair struct {
	Idx int
	Num int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}