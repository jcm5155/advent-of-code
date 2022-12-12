extern crate core;

use std::collections::HashSet;

mod inputs;

fn main() {
    let pzl = inputs::read("3");
    let (mut p1, mut p2) = (0, 0);

    let mut current_bag_idx = 0;
    let mut current_bags: Vec<HashSet<u8>> = Vec::with_capacity(3);

    let get_char_score = |c: &u8 | -> i32 {
        if *c > 97 {
            // lowercase
            return (*c as i32)  - 96;
        }
        // uppercase
        return (*c as i32) - 38;
    };

    for (idx, line) in pzl.split("\n").enumerate() {
        // part 1
        let line_bytes = line.as_bytes();
        let boundary = line_bytes.len() / 2;
        let first_bag: HashSet<u8> = HashSet::from_iter(line_bytes[0..boundary].iter().cloned());
        for c in line_bytes[boundary..].iter() {
            if first_bag.contains(c) {
                p1 += get_char_score(c);
                break;
            }
        }

        // part 2
        current_bag_idx = idx % 3;
        let both_bags = HashSet::from_iter(line_bytes.iter().cloned());

        // this is stupid, but was struggling to initialize the vector with 0 values lol
        if current_bags.len() < current_bag_idx + 1 {
            current_bags.push(both_bags);
        } else {
            current_bags[current_bag_idx] = both_bags;
        }

        if current_bag_idx == 2 {
            for shared_char in current_bags[0].intersection(&current_bags[1]) {
                if current_bags[2].contains(shared_char) {
                    p2 += get_char_score(shared_char);
                    break;
                }
            }
        }
    }

    println!("p1: {}\np2: {}", p1, p2);
}