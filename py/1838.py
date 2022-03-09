from typing import *

class Solution:
    def maxFrequency(self, nums: List[int], k: int) -> int:
        nums = sorted(nums)
        
        res = 0
        l = r = 0
        total = 0
        while r < len(nums):
            while nums[r] * (r - l) - total > k:
                total -= nums[l]
                l += 1 # 注意：先使用数据，后移动！！
            total += nums[r]
            res = max(res, r - l + 1)
            r += 1 
        
        return res
    
