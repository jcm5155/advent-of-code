extern crate core;

mod inputs;

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

const ATTACK_ROCK: Attack = Attack{name: ROCK, beats: SCISSORS, loses_to: PAPER, score: 1};
const ATTACK_PAPER: Attack = Attack{name: PAPER, beats: ROCK, loses_to: SCISSORS, score: 2};
const ATTACK_SCISSORS: Attack = Attack{name: SCISSORS, beats: PAPER, loses_to: ROCK, score: 3};

fn get_attack(attack_name: char) -> Attack {
    return match attack_name {
        ROCK | ROCK_OR_LOSE => ATTACK_ROCK,
        PAPER | PAPER_OR_DRAW => ATTACK_PAPER,
        SCISSORS | SCISSORS_OR_WIN => ATTACK_SCISSORS,
        _ => panic!("unknown attack"),
    }
}

fn main() {
    let pzl = inputs::read("2");
    let (mut p1, mut p2) = (0, 0);

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

    println!("p1: {}\np2: {}", p1, p2);
}
