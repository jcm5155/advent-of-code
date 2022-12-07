use std::collections::{HashMap, HashSet};
use regex::Regex;
use crate::inputs;

pub fn day7() -> (i32, i32) {
    let (mut p1, mut p2) = (0, 0);


    return (p1, p2);
}

pub fn day6() -> (usize, usize) {
    let pzl: Vec<char> = inputs::read("day6").chars().collect();

    let mut p1 = 4;
    while p1 < pzl.len() {
        let hs: HashSet<char> = pzl[p1-4..p1].iter().cloned().collect();
        if hs.len() == 4 { break; }
        p1 += 1
    }

    let mut p2: usize = 14 + p1;
    while p2 < pzl.len() {
        let hs: HashSet<char> = pzl[p2-14..p2].iter().cloned().collect();
        if hs.len() == 14 { break; }
        p2 += 1
    }

    return (p1, p2);
}


pub fn day5() -> (String, String) {
    let pzl = inputs::read("day5");
    let sections: Vec<&str> = pzl.split("\n\n").collect();

    let mut stacks: Vec<Vec<char>> = vec![Vec::new()];
    let mut stacks2: Vec<Vec<char>> = vec![Vec::new()];

    // dumb pre-initialize stacks
    for c in sections[0].split("\n").last().unwrap().chars() {
        if c.is_numeric() {
            stacks.push(Vec::new());
            stacks2.push(Vec::new());
        }
    }

    // populate initial positions
    for line in sections[0].split("\n") {
        let mut idx: usize = 0;
        let mut col: usize = 1;

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

    // reverse input stacks
    for i in 0..stacks.len() {
        stacks[i].reverse();
        stacks2[i].reverse();
    }

    // do the stuff
    let re = Regex::new(r"^move (\d+) from (\d) to (\d)$").unwrap();
    for instruction in sections[1].lines() {
        let cap = match re.captures(instruction) {
            Some(c) => c,
            _ => continue,
        };

        let move_amount = cap[1].parse::<usize>().unwrap();
        let from = cap[2].parse::<usize>().unwrap();
        let to = cap[3].parse::<usize>().unwrap();

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

    // part 1
    let p1: String = stacks[1..stacks.len()]
        .iter()
        .map(|m| *m.last().unwrap())
        .collect();

    // part 2
    let p2: String = stacks2[1..stacks2.len()]
        .iter()
        .map(|m| *m.last().unwrap())
        .collect();

    return (p1, p2);
}


pub fn day4() -> (i32, i32) {
    let pzl = inputs::read("day4");
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

    return (p1, p2);
}


pub fn day3() -> (i32, i32) {
    let pzl = inputs::read("day3");
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

    return (p1, p2);
}


pub fn day2() -> (i32, i32) {
    let pzl = inputs::read("day2");
    let (mut p1, mut p2) = (0, 0);

    const ROCK: char = 'A';
    const PAPER: char = 'B';
    const SCISSORS: char = 'C';
    const ROCK_OR_LOSE: char = 'X';
    const PAPER_OR_DRAW: char  = 'Y';
    const SCISSORS_OR_WIN: char = 'Z';

    struct Attack {
        name: char,
        beats: char,
        loses_to: char,
        score: i32,
    }

    const rock: Attack = Attack{name: ROCK, beats: SCISSORS, loses_to: PAPER, score: 1};
    const paper: Attack = Attack{name: PAPER, beats: ROCK, loses_to: SCISSORS, score: 2};
    const scissors: Attack = Attack{name: SCISSORS, beats: PAPER, loses_to: ROCK, score: 3};

    let get_attack = | name | -> Attack {
        match name {
            ROCK | ROCK_OR_LOSE => rock,
            PAPER | PAPER_OR_DRAW => paper,
            SCISSORS | SCISSORS_OR_WIN => scissors,
            _ => panic!("unknown attack"),
        }
    };

    for line in pzl.split("\n") {
        let choices: Vec<char> = line.chars().filter(|c| c.is_alphabetic()).collect();
        let (elf_choice, my_choice) = (choices[0], choices[1]);
        let (elf_attack, my_attack) = (get_attack(elf_choice), get_attack(my_choice));

        // part 1
        p1 += my_attack.score;
        if my_attack.beats == elf_choice {
            p1 += 6;
        } else if my_attack.name == elf_choice {
            p1 += 3;
        }

        // part 2
        match my_choice {
            SCISSORS_OR_WIN => {
                p2 += 6;
                p2 += get_attack(elf_attack.loses_to).score;
            },
            PAPER_OR_DRAW => {
                p2 += 3;
                p2 += elf_attack.score;
            },
            ROCK_OR_LOSE => {
                p2 += get_attack(elf_attack.beats).score;
            }
            _ => panic!("invalid choice"),
        }
    }
    return (p1, p2);
}


pub fn day1() -> (i32, i32) {
    let pzl = inputs::read("day1");
    let mut elf_total_calories: Vec<i32> = Vec::new();

    for elf_backpack in pzl.split("\n\n") {
        let foods = elf_backpack.split("\n");
        let mut calorie_sum = 0;

        for food in foods {
            calorie_sum += food.parse::<i32>().unwrap();
        }

        elf_total_calories.push(calorie_sum);
    }

    elf_total_calories.sort_by(|a, b| b.cmp(a));
    return (elf_total_calories[0], elf_total_calories[0..3].iter().sum());
}




