extern crate core;

mod inputs;

fn main() {
    let pzl = inputs::read("test");
    // let pzl = inputs::read("13");
    let (p1, p2) = (0, 0);


    println!("p1: {}\np2: {}", p1, p2);
}