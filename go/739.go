package main

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))

	result[len(result)-1] = 0
	stack := []int{len(temperatures) - 1}

	for i := len(temperatures) - 2; i >= 0; i-- {
		t := temperatures[i]

		for len(stack) > 0 && t >= temperatures[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			result[i] = 0
		} else {
			result[i] = stack[len(stack)-1] - i
		}

		stack = append(stack, i)
	}

	return result
}
