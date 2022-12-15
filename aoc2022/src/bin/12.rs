use pathfinding::prelude::astar;

mod inputs;

const START_SCORE: u8 = 45;
const LOWERCASE_A: u8 = 1;
const END_SCORE: u8 = 31;
const UPPERCASE_DIFF: u8 = 38;
const UPPERCASE_Z: u8 = 52;
const LOWERCASE_Z: u8 = 26;
const LOWERCASE_DIFF: u8 = 58;

#[derive(Clone, Debug, Eq, Hash, Ord, PartialEq, PartialOrd)]
pub struct Point {
    x: usize,
    y: usize,
}

impl Point {
    pub fn distance(&self, other: &Point) -> u32 {
        (self.x.abs_diff(other.x) + self.y.abs_diff(other.y)) as u32
    }

    pub fn successors(&self, rows: &[Vec<u8>]) -> Vec<(Point, u32)> {
        let y = self.y;
        let x = self.x;
        let mut successors: Vec<Point> = Vec::new();
        let item = rows[y][x];
        let is_z = item == LOWERCASE_Z;
        let row_length = rows[y].len();

        // left
        if x > 0 && rows[y][x - 1] <= item + 1 || is_z && rows[y][x - 1] == END_SCORE {
            successors.push(Point { x: (x - 1), y });
        }

        // right
        if x < row_length - 1 && rows[y][x + 1] <= item + 1 || is_z && rows[y][x + 1] == END_SCORE {
            successors.push(Point { x: (x + 1), y });
        }

        // up
        if y > 0 && rows[y - 1][x] <= item + 1 || is_z && rows[y - 1][x] == END_SCORE {
            successors.push(Point { x, y: (y - 1) });
        }

        // down
        if y < rows.len() - 1 && rows[y + 1][x] <= item + 1 || is_z && rows[y + 1][x] == END_SCORE {
            successors.push(Point { x, y: (y + 1) });
        }

        successors.into_iter().map(|p| (p, 1)).collect()
    }
}

pub fn score_item(item: char) -> u8 {
    let mut scored_item = item as u8 - UPPERCASE_DIFF;
    if scored_item > UPPERCASE_Z {
        scored_item -= LOWERCASE_DIFF;
    }
    scored_item as u8
}

pub fn parse(input: &str) -> Vec<Vec<u8>> {
    let rows = input
        .split('\n')
        .filter(|m| !m.is_empty())
        .map(|row| row.chars().map(score_item).collect::<Vec<u8>>())
        .collect();

    rows
}

pub fn get_start_and_end(rows: &[Vec<u8>]) -> (Point, Point) {
    let mut start = Point { x: 0, y: 0 };
    let mut end = Point { x: 0, y: 0 };
    for (y, row) in rows.iter().enumerate() {
        for (x, item) in row.iter().enumerate() {
            if *item == START_SCORE {
                start = Point { x, y };
            } else if *item == END_SCORE {
                end = Point { x, y };
            }
        }
    }
    return (start, end)
}

pub fn get_starting_locations(rows: &[Vec<u8>]) -> Vec<Point> {
    let mut start = vec![];
    for (row_index, row) in rows.iter().enumerate() {
        for (col_index, item) in row.iter().enumerate() {
            if *item == START_SCORE || *item == LOWERCASE_A {
                start.push(Point { x: col_index, y: row_index });
            }
        }
    }
    return start;
}

pub fn part_one(input: &str) -> u32 {
    let elevations = parse(input);
    let (start, end) = get_start_and_end(&elevations);
    let result = astar(
        &start,
        |pos: &Point| pos.successors(&elevations),
        |pos| pos.distance(&end),
        |pos| pos == &end
    );

    return result.unwrap().1;
}

pub fn part_two(input: &str) -> u32 {
    let elevations = parse(input);
    let (_start_pos, end_pos) = get_start_and_end(&elevations);
    let possible_starts = get_starting_locations(&elevations);

    let mut shortest: Vec<u32> = vec![];

    for start in possible_starts {
        let result = astar(
            &start,
            |pos: &Point| pos.successors(&elevations),
            |pos| pos.distance(&end_pos),
            |pos| pos == &end_pos
        );
        if let Some(dist) = result {
            shortest.push(dist.1);
        }
    }
    return *shortest.iter().min().unwrap();
}

fn main() {
    let input = &inputs::read("12");
    println!("p1: {}\np2: {}", part_one(input), part_two(input));
}
