package main

import "math"

func minFlips(s string) int {
	res := math.MaxInt

	ds := s + s
	alt1 := make([]byte, len(ds))
	alt2 := make([]byte, len(ds))

	for i := range ds {
		alt1[i] = '0' + byte(i%2)
		alt2[i] = '0' + byte((i+1)%2)
	}

	diff1 := 0
	diff2 := 0

	l := 0
	for r := range ds {
		if alt1[r] != ds[r] {
			diff1 += 1
		}
		if alt2[r] != ds[r] {
			diff2 += 1
		}
		if r < len(s)-1 {
			continue
		}
		if r-l+1 == len(s) {
			res = min(res, diff1)
			res = min(res, diff2)
		}
		if alt1[l] != ds[l] {
			diff1 -= 1
		}
		if alt2[l] != ds[l] {
			diff2 -= 1
		}
		l++
	}

	if res == math.MaxInt {
		return 0
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}