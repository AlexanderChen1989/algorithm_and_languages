package main

func main() {
	println(characterReplacement("ABAB", 2))
	println(characterReplacement("AABABBA", 1))
}

func characterReplacement(s string, k int) int {
	counter := [26]int{}
	res := 0
	l := 0
	for r := range s {
		counter[s[r]-'A']++

		if (r-l+1)-maxCounter(counter) <= k {
			res = max(res, r-l+1)
		}
		for ; (r-l+1)-maxCounter(counter) > k; l++ {
			counter[s[l]-'A']--
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxCounter(counter [26]int) int {
	res := counter[0]
	for _, v := range counter {
		res = max(res, v)
	}
	return res
}
