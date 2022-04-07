fn main() {
    let mut func: fn(i32, i32) -> i32 = add1;
    println!("{}", func(1, 2));
    func = add2;
    println!("{}", func(1, 2));

    let funcs: Vec<fn(i32, i32) -> i32> = vec![add1, add2, |a: i32, b: i32| a * b];

    for func in funcs {
        println!(">>> {}", func(10, 20));
    }
}

fn add1(
    a: i32,
    b: i32,
) -> i32 {
    a + b
}

fn add2(
    a: i32,
    b: i32,
) -> i32 {
    a + b + 100
}
