fn main() {
    let inputs = vec![("ABAB", 2), ("AABABBA", 1)];

    for (s, k) in inputs {
        let r = Solution::character_replacement(s.into(), k);
        println!(">>> {r}");
    }
}

struct Solution;

impl Solution {
    pub fn character_replacement(s: String, k: i32) -> i32 {
        let bs = s.as_bytes();
        let k = k as usize;
        let mut counter = [0usize; 26];

        let mut l = 0;
        let mut res = 0;

        for (r, c) in bs.iter().enumerate() {
            let idx = (c - b'A') as usize;
            counter[idx] += 1;
            if r - l + 1 - counter.iter().max().unwrap() <= k {
                res = res.max(r - l + 1);
            }
            while r - l + 1 - counter.iter().max().unwrap() > k {
                let c = bs[l];
                let idx = (c - b'A') as usize;
                counter[idx] -= 1;
                l += 1;
            }
        }

        return res as i32;
    }
}
