package main

func generateParenthesis(n int) []string {
	result := []string{}

	stack := []Item{Item{}}

	for len(stack) > 0 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if item.Left == n && item.Right == n {
			result = append(result, item.Str)
			continue
		}

		if item.Left == item.Right {
			stack = append(
				stack,
				Item{Left: item.Left + 1, Right: item.Right, Str: item.Str + "("},
			)
			continue
		}
		stack = append(
			stack,
			Item{Left: item.Left, Right: item.Right + 1, Str: item.Str + ")"},
		)
		if item.Left+1 <= n {
			stack = append(
				stack,
				Item{Left: item.Left + 1, Right: item.Right, Str: item.Str + "("},
			)
		}
	}

	return result
}

type Item struct {
	Str   string
	Left  int
	Right int
}
