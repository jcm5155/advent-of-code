extern crate core;

mod inputs;

fn main() {
    let pzl = inputs::read("1");
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
    let p1 = elf_total_calories[0];
    let p2: i32 = elf_total_calories[0..3].iter().sum();
    println!("p1: {}\np2: {}", p1, p2);
}