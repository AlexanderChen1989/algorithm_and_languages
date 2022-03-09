class Solution:
    def minFlips(self, s: str) -> int:
        s2 = s * 2
        alt1 = ''.join([str(i % 2) for i in range(len(s2))])
        alt2 = ''.join([str((i + 1) % 2) for i in range(len(s2))])

        diff1 = 0
        diff2 = 0

        res = float('inf')

        l = 0
        for r in range(len(s2)):
            # 先更新窗口
            if s2[r] != alt1[r]:
                diff1 += 1
            if s2[r] != alt2[r]:
                diff2 += 1
            # 当r == len(s) - 1时就可以取值了
            if r < len(s) - 1:
                continue
            res = min(res, diff1, diff2)

            if s2[l] != alt1[l]:
                diff1 -= 1
            if s2[l] != alt2[l]:
                diff2 -= 1

            l += 1

        return res


ss = ["111000", "010", "1110"]

for s in ss:
    print(Solution().minFlips(s))