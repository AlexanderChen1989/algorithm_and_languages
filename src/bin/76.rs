use std::collections::HashMap;

fn main() {
    let ss = vec![("ADOBECODEBANC", "ABC"), ("a", "a"), ("a", "aa")];
    for (s, t) in ss {
        let res = Solution::min_window(s.into(), t.into());
        println!("{res}");
    }
}

struct Solution;

impl Solution {
    pub fn min_window(
        s: String,
        t: String,
    ) -> String {
        use std::collections::HashMap;

        if t == "" || s.len() < t.len() {
            return "".into();
        }
        let mut tm: HashMap<u8, i32> = HashMap::new();
        for c in t.as_bytes() {
            *tm.entry(*c).or_insert(0) += 1;
        }

        let bs = s.as_bytes();
        let mut wm: HashMap<u8, i32> = HashMap::new();

        let mut has_res = false;
        let (mut l, mut r, mut has, need) = (0usize, 0usize, 0usize, tm.len());
        let (mut res, mut res_len) = ((0, 0), usize::MAX);

        while r < s.len() {
            let c = &bs[r];
            *wm.entry(*c).or_insert(0) += 1;
            if tm.contains_key(c) && wm[c] == tm[c] {
                has += 1;
            }
            while has == need {
                let len = r - l + 1;
                if len < res_len {
                    has_res = true;
                    res = (l, r);
                    res_len = len;
                }
                let c = &bs[l];
                *wm.get_mut(c).unwrap() -= 1;
                if tm.contains_key(c) && wm[c] < tm[c] {
                    has -= 1;
                }
                l += 1;
            }
            r += 1;
        }
        if !has_res {
            return "".into();
        }
        let (l, r) = res;
        return s[l..r + 1].into();
    }
}
