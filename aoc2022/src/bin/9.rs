extern crate core;

use std::collections::HashSet;

mod inputs;

#[derive(Debug, Clone)]
struct Segment {
    x: i32,
    y: i32,
}

impl Segment {
    fn coords(&self) -> (i32, i32) { (self.x, self.y) }

    fn move_toward(&mut self, x: i32, y: i32) {
        self.x += x;
        self.y += y;
    }

    fn follow(&mut self, leader: &Segment) {
        let diff_x = leader.x - self.x;
        let diff_y = leader.y - self.y;

        if diff_x.abs() < 2 && diff_y.abs() < 2 {
            return
        }

        let mut move_x = 0;
        if diff_x != 0 {
            move_x = diff_x / diff_x.abs();
        }

        let mut move_y = 0;
        if diff_y != 0 {
            move_y = diff_y / diff_y.abs();
        }

        self.move_toward(move_x, move_y);
    }
}

fn main() {
    let pzl = inputs::read("9");
    let mut rope: Vec<Segment> = vec![Segment {x: 0, y: 0}; 10];
    let mut visited_p1: HashSet<(i32, i32)> = HashSet::new();
    let mut visited_p2: HashSet<(i32, i32)> = HashSet::new();

    for line in pzl.lines() {
        let mut instructions = line.split(" ");
        let (x_dir, y_dir) = match instructions.next().unwrap() {
            "R" => (1,  0),
            "L" => (-1, 0),
            "U" => (0,  1),
            "D" => (0, -1),
            _ => panic!("unknown direction"),
        };
        let distance = instructions.next().unwrap().parse::<u8>().unwrap();

        for _ in 0..distance {
            let _ = &mut rope[0].move_toward(x_dir, y_dir);
            let mut prev = (&mut rope[0]).clone();

            for curr in rope[1..].iter_mut() {
                curr.follow(&prev);
                prev = curr.clone();
            }

            let _ = &mut visited_p1.insert(rope[1].coords());
            let _ = &mut visited_p2.insert(rope[9].coords());
        }
    }

    println!("p1: {}\np2: {}", visited_p1.len(), visited_p2.len());
}