from typing import *
import collections

class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        q = collections.deque()

        res = []
        l = r = 0
        
        while r < len(nums):
            while q and nums[q[-1]] < nums[r]:
                q.pop()
            q.append(r)
            while q[0] < l:
                q.popleft()
            
            if r + 1 >= k:
                res.append(nums[q[0]])
                l += 1
            r += 1
        
        return res



solution = Solution()

nums = [1,3,-1,-3,5,3,6,7]
k = 3
print(solution.maxSlidingWindow(nums, k))

nums = [1]
k = 1
print(solution.maxSlidingWindow(nums, k))

