class Solution:
    def minWindow(self, s: str, t: str) -> str:
        if t == "":
            return ""
        tm = {}
        for c in t:
            tm[c] = 1 + tm.get(c, 0)
        need, has = len(tm), 0
        l, r = 0, 0
        wm = {}

        res, resLen = [-1, -1], float('infinity')

        for r in range(len(s)):
            c = s[r]
            wm[c] = 1 + wm.get(c, 0)
            if c in tm and tm[c] == wm[c]:
                has += 1
            while has == need:
                tLen = r - l + 1
                if tLen < resLen:
                    res, resLen = [l, r], tLen
                c = s[l]
                wm[c] -= 1
                if c in tm and tm[c] > wm[c]:
                    has -= 1
                l += 1
            r += 1

        [l, r] = res
        return s[l:r + 1]


print(Solution().minWindow("ADOBECODEBANC", "ABC"))