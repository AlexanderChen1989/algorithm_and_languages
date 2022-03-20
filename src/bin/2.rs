#![feature(let_chains)]

fn main() {
    let v1 = vec![9, 9, 9, 9, 9, 9, 9];
    let v2 = vec![9, 9, 9, 9];
    let l1 = vec_to_list(&v1);
    let l2 = vec_to_list(&v2);

    let vs = list_to_vec(Solution::add_two_numbers(l1, l2));

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

struct Solution;

impl Solution {
    pub fn add_two_numbers(
        l1: Option<Box<ListNode>>,
        l2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        let mut l1 = l1;
        let mut l2 = l2;

        let mut result = Some(Box::new(ListNode { val: 0, next: None }));
        let mut pre = &mut result;

        let mut rem = 0;

        loop {
            if let Some(ref mut n1) = l1 {
                if let Some(ref mut n2) = l2 {
                    rem += n1.val;
                    rem += n2.val;
                    n1.val = rem % 10;
                    rem /= 10;
                } else {
                    break;
                }
            } else {
                break;
            }
            pre.as_mut().unwrap().next = l1;
            pre = &mut pre.as_mut().unwrap().next;

            l1 = std::mem::take(&mut pre.as_mut().unwrap().next);

            l2 = l2.unwrap().next;
        }

        if l1.is_some() {
            pre.as_mut().unwrap().next = l1;
        }

        if l2.is_some() {
            pre.as_mut().unwrap().next = l2;
        }

        while rem != 0 {
            if pre.as_ref().unwrap().next.is_none() {
                pre.as_mut().unwrap().next = Some(Box::new(ListNode {
                    val: rem,
                    next: None,
                }));
                break;
            }

            rem += pre.as_ref().unwrap().next.as_ref().unwrap().val;
            pre.as_mut().unwrap().next.as_mut().unwrap().val = rem % 10;
            rem /= 10;
            pre = &mut pre.as_mut().unwrap().next;
        }

        return result.unwrap().next;
    }
}
