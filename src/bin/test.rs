use nice::{ListNode, vec_to_list, print_list};

fn main() {
	let head = vec_to_list(vec![]);
	print_list(&Solution::insertion_sort_list(head));
}

struct Solution;

impl Solution {
    pub fn insertion_sort_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut head = head;

        let mut sorted = Some(Box::new(ListNode::new(0)));

        while let Some(mut node) = head {
            head = std::mem::take(&mut node.next);

            let mut p = &mut sorted;
            if p.as_ref().unwrap().next.is_none() {
                p.as_mut().unwrap().next = Some(node);
                continue;
            }

            while p.as_ref().unwrap().next.is_some()
                && p.as_ref().unwrap().next.as_ref().unwrap().val < node.val
            {
                p = &mut p.as_mut().unwrap().next;
            }

			node.next = std::mem::take(&mut p.as_mut().unwrap().next);

			p.as_mut().unwrap().next = Some(node);
        }

        sorted.unwrap().next
    }
}
