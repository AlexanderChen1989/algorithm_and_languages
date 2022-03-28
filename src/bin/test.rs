use nice::{print_list, vec_to_list, ListNode};

fn main() {
    let head = vec_to_list(vec![-1, 5, 3, 4, 0]);
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
