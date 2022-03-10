package main 


func findAnagrams(s string, p string) []int {
    ms := [26]byte{}
    mp := [26]byte{}
    
    for _, c := range p {
        mp[c - 'a'] += 1
    }
    
    res := []int{}
    l := 0 
    for r := range s {
        ms[s[r] - 'a'] += 1
        if r < len(p) - 1 {
            continue
        }
        if is_equal(ms, mp) {
            res = append(res, l)
        }
        ms[s[l] - 'a'] -= 1
        l++
    }
    
    return res
}

func is_equal(ms, mp [26]byte) bool {
    for i := range mp {
        if ms[i] != mp[i] {
            return false
        }
    }
    return true
}