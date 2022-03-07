package main

import (
	"math"
)

func main() {
	s := "a"
	t := "aa"

	println(minWindow(s, t))
}

func minWindow(s string, t string) string {
	if t == "" {
		return ""
	}
	window, counter := map[byte]int{}, map[byte]int{}
	for i := range t {
		counter[t[i]]++
	}
	need, has := len(counter), 0
	res, resLen := []int{-1, -1}, math.MaxInt

	for r, l := 0, 0; r < len(s); r++ {
		c := s[r]
		window[c]++
		if window[c] == counter[c] {
			has++
		}
		for ; need == has; l++ {
			t := r - l + 1
			if resLen > t {
				resLen = t
				res = []int{l, r}
			}
			c := s[l]

			if window[c]--; window[c] < counter[c] {
				has--
			}
		}
	}
	if res[0] == -1 {
		return ""
	}
	return s[res[0] : res[1]+1]
}
