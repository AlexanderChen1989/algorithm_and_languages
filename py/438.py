from typing import *
import string

class Solution:
    def is_equal(self, ms, mp):
        for c in string.ascii_lowercase:
            if ms[c] != mp[c]:
                return False 
        return True

    def findAnagrams(self, s: str, p: str) -> List[int]:
        ms = {}
        mp = {}
        
        for c in string.ascii_lowercase:
            ms[c] = 0
            mp[c] = 0

        for c in p:
            mp[c] += 1

        res = []
        l = 0
        for r in range(len(s)):
            ms[s[r]] += 1
            if r - l + 1 < len(p):
                continue
            if self.is_equal(ms, mp):
                res.append(l)
                
            ms[s[l]] -= 1
            l += 1

        return res



solution = Solution()

for s, p in [("cbaebabacd", "abc"), ("abab", "ab")]:
    result = solution.findAnagrams(s, p)
    print(result)
