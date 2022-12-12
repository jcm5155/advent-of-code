extern crate core;

use regex::Regex;
use utils::SolutionPart;

mod inputs;
mod utils;


fn main() {
    let pzl = inputs::read("11");
    let p1 = solve(&pzl,SolutionPart::One);
    let p2 = solve(&pzl,SolutionPart::Two);
    println!("p1: {}\np2: {}", p1, p2);
}

pub enum MathOp<T> {
    Plus(Option<T>),
    Multiply(Option<T>),
}

struct Monkey {
    items: Vec<i64>,
    inspected: i64,
    true_target: usize,
    false_target: usize,
    divisor: i32,
    operation: MathOp<i64>,
}

fn solve(pzl: &String, part: SolutionPart) -> i64 {
    let mut barrel = parse_input(&pzl);
    let modulus: i64 = barrel.iter().map(|m| m.divisor as i64).product();

    let (rounds, worry_divisor) = match part {
        SolutionPart::One => (20, 3),
        SolutionPart::Two => (10000, 1)
    };

    for i in 0..barrel.len() * rounds {
        let monkey_n = i % barrel.len();
        let items = std::mem::take(&mut barrel[monkey_n].items);
        barrel[monkey_n].inspected += items.len() as i64;

        for mut item in items {
            // monkey see
            item = match &barrel[monkey_n].operation {
                MathOp::Plus(Some(amt)) => item + amt,
                MathOp::Multiply(Some(amt)) => item * amt,
                MathOp::Plus(None) => item + item,
                MathOp::Multiply(None) =>  item * item,
            } % modulus / worry_divisor;

            // monkey do
            let target = match item % barrel[monkey_n].divisor as i64 == 0 {
                true => barrel[monkey_n].true_target,
                false => barrel[monkey_n].false_target,
            };
            barrel[target].items.push(item);
        }
    }

    barrel.sort_by(|a, b| b.inspected.cmp(&a.inspected));
    return barrel[0].inspected * barrel[1].inspected;
}

fn parse_input(input: &str) -> Vec<Monkey> {
    let re = Regex::new(r"Monkey (\d):
  Starting items: (.*)
  Operation: new = old ([+,*]) (.+)
  Test: divisible by (\d+)
    If true: throw to monkey (\d)
    If false: throw to monkey (\d)").unwrap();

    let mut barrel: Vec<Monkey> = vec![];
    for instruction_block in input.split("\n\n") {
        let caps = re.captures(instruction_block.trim()).unwrap();

        let items: Vec<i64> = caps[2]
            .split(", ")
            .map(|e| e.parse::<i64>().unwrap())
            .collect();

        let operation = match (caps[3].chars().nth(0).unwrap(), &caps[4]) {
            ('+', "old") => MathOp::Plus(None),
            ('*', "old") => MathOp::Multiply(None),
            ('+', amt) => MathOp::Plus(Some(amt.parse::<i64>().unwrap())),
            ('*', amt) => MathOp::Multiply(Some(amt.parse::<i64>().unwrap())),
            _ => panic!("unknown operation")
        };

        let divisor = caps[5].parse::<i32>().unwrap();
        let true_target = caps[6].parse::<usize>().unwrap();
        let false_target = caps[7].parse::<usize>().unwrap();

        barrel.push(Monkey {
            items,
            divisor,
            true_target,
            false_target,
            operation,
            inspected: 0,
        });
    }

    return barrel;
}