fn main() {
    let v1 = vec![1, 2, 3, 4, 5];
    let head = vec_to_list(&v1);
    let vs = list_to_vec(Solution::reverse_between(head, 2, 4));

    println!("{vs:?}");
}

fn list_to_vec(mut list: Option<Box<ListNode>>) -> Vec<i32> {
    let mut vs = Vec::new();
    while let Some(node) = list {
        vs.push(node.val);
        list = node.next;
    }

    vs
}

fn vec_to_list(vs: &[i32]) -> Option<Box<ListNode>> {
    if vs.len() == 0 {
        None
    } else {
        Some(Box::new(ListNode {
            val: vs[0],
            next: vec_to_list(&vs[1..]),
        }))
    }
}

#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }
}

struct Solution;

impl Solution {
    pub fn reverse_between(
        head: Option<Box<ListNode>>,
        left: i32,
        right: i32,
    ) -> Option<Box<ListNode>> {
        let mut h = Some(Box::new(ListNode { val: 0, next: head }));
        let mut pre = &mut h;

        for _ in 0..left - 1 {
            pre = &mut pre.as_mut().unwrap().next;
        }

        let mut r = std::mem::take(&mut pre.as_mut().unwrap().next);
        let mut p = std::mem::take(&mut r.as_mut().unwrap().next);

        for _ in 0..(right - left) {
            let t = std::mem::take(&mut p.as_mut().unwrap().next);
            p.as_mut().unwrap().next = r.to_owned();
            r = p.to_owned();
            p = t;
        }

        pre.as_mut().unwrap().next = r.to_owned();

        while pre.is_some() && pre.as_ref().unwrap().next.is_some() {
            pre = &mut pre.as_mut().unwrap().next
        }

        pre.as_mut().unwrap().next = p;

        return h.unwrap().next;
    }
}

fn take(p: &mut Option<Box<ListNode>>) -> (Option<Box<ListNode>>, Option<Box<ListNode>>) {
    let mut node = std::mem::take(&mut p.as_mut().unwrap().next);
    let next = std::mem::take(&mut node.as_mut().unwrap().next);

    (node, next)
}
