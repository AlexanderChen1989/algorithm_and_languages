fn main() {
    /*
        AsRef/AsMut
    */

    let mut hello = Hello {
        val: 10,
        data: Data { val: 100 },
    };

    let ref_data_val: &u64 = hello.as_ref();
    println!("{ref_data_val}");

    let data_val: &mut u64 = hello.as_mut();
    *data_val = 1000;
    let ref_data_val: &u64 = hello.as_ref();
    println!("{ref_data_val}");

    let ref_f32: &f32 = hello.as_ref();
    println!("{ref_f32}");
}

fn print_i32(s: impl AsRef<String>) {
    println!(">>>> {}", s.as_ref());
}

struct Hello {
    val: i32,
    data: Data,
}

struct Data {
    val: u64,
}

impl AsRef<Hello> for Hello {
    fn as_ref(&self) -> &Hello {
        self
    }
}

impl AsRef<f32> for Hello {
    fn as_ref(&self) -> &f32 {
        &100.0
    }
}

impl AsRef<i32> for Hello {
    fn as_ref(&self) -> &i32 {
        &self.val
    }
}

impl AsRef<Data> for Hello {
    fn as_ref(&self) -> &Data {
        &self.data
    }
}

impl AsRef<u64> for Hello {
    fn as_ref(&self) -> &u64 {
        &self.data.val
    }
}

impl AsMut<i32> for Hello {
    fn as_mut(&mut self) -> &mut i32 {
        &mut self.val
    }
}

impl AsMut<Data> for Hello {
    fn as_mut(&mut self) -> &mut Data {
        &mut self.data
    }
}

impl AsMut<u64> for Hello {
    fn as_mut(&mut self) -> &mut u64 {
        &mut self.data.val
    }
}
