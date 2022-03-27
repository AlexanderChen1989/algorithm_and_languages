#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }
}

pub fn vec_to_list(vals: Vec<i32>) -> Option<Box<ListNode>> {
    let mut list = None;
    for val in vals.into_iter().rev() {
        list = Some(Box::new(ListNode { val, next: list }));
    }
    return list;
}

pub fn print_list(head: &Option<Box<ListNode>>) {
    let mut head = head;
    while let Some(node) = head {
        print!("{}\t", node.val);
        head = &node.next;
    }
    println!();
}
