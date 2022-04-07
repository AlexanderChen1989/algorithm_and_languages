use std::borrow::Borrow;

fn main() {
	let h = Hello{val: 10};
	let ref_h = h.borrow();
}


struct Hello {
	val: i32
}

// impl Borrow<i32> for Hello {
//     fn borrow(&self) -> &i32 {
//         &self.val
//     }
// }
