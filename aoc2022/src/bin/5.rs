extern crate core;

use regex::Regex;

mod inputs;

fn main() {
    let pzl = inputs::read("5");
    let sections: Vec<&str> = pzl.split("\n\n").collect();

    let mut stacks: Vec<Vec<char>> = Vec::new();
    let mut stacks2: Vec<Vec<char>> = Vec::new();

    // look at the last row of the diagram to find how many Vec<char> to init
    for c in sections[0].split("\n").last().unwrap().chars() {
        if c.is_numeric() {
            stacks.push(Vec::new());
            stacks2.push(Vec::new());
        }
    }

    // populate initial positions
    let diagram: Vec<&str> = sections[0].split("\n").collect();
    for line in diagram.iter().rev() {
        let mut idx: usize = 0;
        let mut col: usize = 0;

        while idx < line.len() {
            let current_char = line.chars().nth(idx + 1).unwrap();
            if current_char.is_alphabetic() {
                stacks[col].push(current_char);
                stacks2[col].push(current_char);
            }
            idx += 4;
            col += 1;
        }
    }

    // do the stuff
    let re = Regex::new(r"^move (\d+) from (\d) to (\d)$").unwrap();
    for instruction in sections[1].lines() {
        let cap = match re.captures(instruction) {
            Some(c) => c,
            _ => continue,
        };

        let move_amount = cap[1].parse::<usize>().unwrap();
        let from = cap[2].parse::<usize>().unwrap() - 1;
        let to = cap[3].parse::<usize>().unwrap() - 1;

        // part 1
        for _ in 0..move_amount {
            let val = stacks[from].pop().unwrap();
            stacks[to].push(val);
        }

        // part 2
        let stacks2_len = stacks2[from].len();
        let mut bulk_move: Vec<char> = stacks2[from].drain(stacks2_len - move_amount..).collect();
        stacks2[to].append(&mut bulk_move);
    }

    let p1: String = stacks
        .iter()
        .map(|m| *m.last().unwrap())
        .collect();

    let p2: String = stacks2
        .iter()
        .map(|m| *m.last().unwrap())
        .collect();

    println!("p1: {}\np2: {}", p1, p2);
}