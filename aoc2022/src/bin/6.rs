extern crate core;

use std::collections::HashSet;

mod inputs;

fn main() {
    let pzl = inputs::read("6");
    let pzl_chars: Vec<char> = pzl.chars().collect();

    let mut p1 = 4;
    while p1 < pzl_chars.len() {
        let hs: HashSet<char> = pzl_chars[p1-4..p1].iter().cloned().collect();
        if hs.len() == 4 { break; }
        p1 += 1
    }

    let mut p2: usize = 14 + p1;
    while p2 < pzl_chars.len() {
        let hs: HashSet<char> = pzl_chars[p2-14..p2].iter().cloned().collect();
        if hs.len() == 14 { break; }
        p2 += 1
    }

    println!("p1: {}\np2: {}", p1, p2);
}