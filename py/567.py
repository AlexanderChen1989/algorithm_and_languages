import string


class Solution:
    def checkInclusion(self, s1: str, s2: str) -> bool:
        ms1 = {}
        ms2 = {}

        for c in string.ascii_lowercase:
            ms1[c] = 0
            ms2[c] = 0

        for c in s1:
            ms1[c] += 1

        matches = 0
        for c in string.ascii_lowercase:
            if ms1[c] == ms2[c]:
                matches += 1

        l = 0
        for r in range(len(s2)):
            c = s2[r]
            ms2[c] += 1
            if ms2[c] == ms1[c]:
                matches += 1
            if r < len(s1) - 1:
                continue

            if matches == 26:
                return True

            c = s2[l]

            if ms1[c] == ms2[c]:
                matches -= 1
            ms2[c] -= 1
            l += 1

        return False
