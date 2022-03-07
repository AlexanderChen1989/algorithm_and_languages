class Solution:
    def characterReplacement(self, s: str, k: int) -> int:        
        counter = {}
        res = 0
        l = 0
        for r in range(len(s)):
            c = s[r]
            counter[c] = 1 + counter.get(c, 0)
            if r - l + 1 - max(counter.values()) <= k:
                res = max(res, r - l + 1)
            while r - l + 1 - max(counter.values()) > k:
                c = s[l]
                counter[c] -= 1 
                l += 1
        
        return res
				
        
print(Solution().characterReplacement('ABAB', 2))
print(Solution().characterReplacement('AABABBA', 1))
        
