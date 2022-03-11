#![feature(let_chains)]

use std::collections::VecDeque;

fn main() {
    let nums = vec![1, 3, -1, -3, 5, 3, 6, 7];
    let k = 3;

    println!("{:?}", Solution::max_sliding_window(nums, k));
}

struct Solution;

impl Solution {
    pub fn max_sliding_window(
        nums: Vec<i32>,
        k: i32,
    ) -> Vec<i32> {
        let mut res: Vec<i32> = Vec::new();

        let mut counter: VecDeque<usize> = VecDeque::new();

        let mut l = 0usize;
        for (r, num) in nums.iter().enumerate() {
            while let Some(v) = counter.back() && *num >= nums[*v] {
               	counter.pop_back();
            }
            counter.push_back(r);
            if r < (k - 1) as usize {
                continue;
            }
            while let Some(v) = counter.front() && *v < l {
				counter.pop_front();
            }
            if r - l + 1 == k as usize {
                if let Some(v) = counter.front() {
                    res.push(nums[*v]);
                }
            }
            l += 1;
        }

        return res;
    }
}
