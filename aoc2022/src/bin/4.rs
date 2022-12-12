extern crate core;

use regex::Regex;

mod inputs;

fn main() {
    let pzl = inputs::read("4");
    let (mut p1, mut p2) = (0, 0);

    let re = Regex::new(r"^(\d+)-(\d+),(\d+)-(\d+)$").unwrap();

    for line in pzl.split("\n") {
        let cap = re.captures(line).unwrap();
        let left_min = &cap[1].parse::<i32>().unwrap();
        let left_max = &cap[2].parse::<i32>().unwrap();
        let right_min = &cap[3].parse::<i32>().unwrap();
        let right_max = &cap[4].parse::<i32>().unwrap();

        // part 1
        if (left_min >= right_min && left_max <= right_max) ||
            (right_min >= left_min && right_max <= left_max) {
            p1 += 1;
        }

        // part 2
        if (left_min <= right_min && left_max >= right_min) ||
            (right_min <= left_min && right_max >= left_min) {
            p2 += 1;
        }
    }

    println!("p1: {}\np2: {}", p1, p2);
}