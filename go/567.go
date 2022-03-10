package main

func checkInclusion(s1 string, s2 string) bool {
	ms1 := [26]int{}
	ms2 := [26]int{}

	for _, c := range s1 {
		ms1[c-'a']++
	}
	l := 0
	for r := range s2 {
		ms2[s2[r]-'a']++
		if r < len(s1)-1 {
			continue
		}

		if is_equal(ms1, ms2) {
			return true
		}
		ms2[s2[l]-'a']--
		l++
	}

	return false
}

func is_equal(ms1, ms2 [26]int) bool {
	for i := range ms1 {
		if ms1[i] != ms2[i] {
			return false
		}
	}

	return true
}
