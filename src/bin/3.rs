use std::collections::HashSet;

fn main() {
    let ss = vec!["abcabcbb", "bbbbb", "pwwkew"];
    for s in ss {
        let res = Solution::length_of_longest_substring(s.into());
        println!("{res}");
    }
}

struct Solution;

impl Solution {
    pub fn length_of_longest_substring(s: String) -> i32 {
        let (mut l, mut r) = (0usize, 0usize);
        let s = s.as_bytes();
        let mut set: HashSet<u8> = HashSet::new();
        let mut res: i32 = 0;

        while r < s.len() {
            while set.contains(&s[r]) {
                set.remove(&s[l]);
                l += 1;
            }
            set.insert(s[r]);
			r += 1;
            res = res.max(set.len() as i32);
        }

        return res;
    }
}
