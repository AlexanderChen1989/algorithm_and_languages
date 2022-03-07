package main 

func main() {
	ss := []string{"abcabcbb", "bbbbb", "pwwkew"}
	for _, s := range ss {
		println(lengthOfLongestSubstring(s))
	}
}

func lengthOfLongestSubstring(s string) int {
	l, r := 0, 0
	set := map[byte]bool{}
	res := 0 

	for r < len(s) {
		for set[s[r]] {
			delete(set, s[l])
			l++
		}
		set[s[r]] = true
		res = max(res, len(set));
		r++
	}
	return res
}

func max(a, b int) int {
	if a > b { return a }
	return b
}