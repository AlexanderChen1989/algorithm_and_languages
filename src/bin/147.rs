use nice::{print_list, vec_to_list, ListNode};

fn main() {
    let head = vec_to_list(vec![-1, 5, 3, 4, 0]);

    print_list(&Solution::insertion_sort_list(head));
}

/*
!!!注意！！！: unwrap != ?
 */

struct Solution;

impl Solution {
    pub fn insertion_sort_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut head = head;
        let mut sorted = Some(Box::new(ListNode::new(0)));
        while let Some(mut node) = head {
            head = node.next;
            node.next = None;

            let mut pre = &mut sorted;

            if pre.as_ref().unwrap().next.is_none() {
                pre.as_mut().unwrap().next = Some(node);
                continue;
            }

            while pre.as_ref().unwrap().next.is_some()
                && pre.as_ref().unwrap().next.as_ref().unwrap().val > node.val
            {
                pre = &mut pre.as_mut().unwrap().next;
            }

            node.next = std::mem::take(&mut pre.as_mut().unwrap().next);
            pre.as_mut().unwrap().next = Some(node);
        }

        let mut head = sorted?.next;

        let mut reversed = None;
        while let Some(mut node) = head {
            head = node.next;
            node.next = reversed;
            reversed = Some(node);
        }

        return reversed;
    }
}
