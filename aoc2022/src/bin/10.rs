extern crate core;

mod inputs;

const SCREEN_WIDTH: usize = 40;
const SCREEN_HEIGHT: usize = 6;

fn main() {
    let pzl = inputs::read("10");
    let mut screen: Vec<char> = vec![' '; SCREEN_WIDTH * SCREEN_HEIGHT];
    let (mut register, mut cycle_count, mut signal_strength) = (1, 0, 0);

    for line in pzl.lines() {
        let (next_op, amt) = match line {
            "noop" => (1, 0),
            _ => (2, line.split(" ").nth(1).unwrap().parse::<i32>().unwrap()),
        };

        for _ in 0..next_op {
            // part 2
            if (register - 1..register + 2).contains(&(cycle_count % SCREEN_WIDTH as i32)) {
                screen[cycle_count as usize] = '#';
            }

            cycle_count += 1;

            // part 1
            if let 20 | 60 | 100 | 140 | 180 | 220 = cycle_count {
                signal_strength += cycle_count * register;
            }
        }
        register += amt;
    }

    println!("p1: {}", signal_strength);
    print!("p2:");
    // print screen
    for (i, pixel) in screen.iter().enumerate() {
        if i % SCREEN_WIDTH == 0 {
            println!();
        }
        print!("{}", pixel);
    }
    println!();
}