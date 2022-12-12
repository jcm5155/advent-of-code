extern crate core;

mod inputs;

fn main() {
    let pzl = inputs::read("8");
    let pzl_nums: Vec<Vec<u32>> = pzl
        .split("\n")
        .map(|m| m.chars()
            .map(|c| c.to_digit(10).unwrap())
            .collect::<Vec<u32>>())
        .collect();

    let (mut visible_trees, mut top_view_score) = (0, 0);
    let (y_len, x_len) = (pzl_nums.len(), pzl_nums[0].len());

    for y in 0..y_len {
        for x in 0..x_len {
            // outer edge trees are automatically visible
            if x == 0 || x == x_len -1 || y == 0 || y == y_len -1 {
                visible_trees += 1;
                // take a gamble on the tree with the highest view score
                // not being one on the edge lmaooo
                continue;
            }

            let current_num = pzl_nums[y][x];
            let (mut left_view, mut right_view) = (0, 0);
            let (mut up_view, mut down_view) = (0, 0);
            let mut is_visible = false;

            // left
            for i in (0..x).rev() {
                left_view += 1;
                if pzl_nums[y][i] >= current_num {
                    break;
                }

                if i == 0 && !is_visible {
                    visible_trees += 1;
                    is_visible = true;
                }
            }

            // right
            for i in x+1..x_len {
                right_view += 1;
                if pzl_nums[y][i] >= current_num {
                    break;
                }

                if i == x_len - 1 && !is_visible {
                    visible_trees += 1;
                    is_visible = true;
                }
            }

            // up
            for i in (0..y).rev() {
                up_view += 1;
                if pzl_nums[i][x] >= current_num {
                    break;
                }

                if i == 0 && !is_visible {
                    visible_trees += 1;
                    is_visible = true;
                }
            }

            // down
            for i in y+1..y_len {
                down_view += 1;
                if pzl_nums[i][x] >= current_num {
                    break;
                }

                if i == y_len -1 && !is_visible {
                    visible_trees += 1;
                    is_visible = true;
                }
            }

            let view_score = left_view * right_view * up_view * down_view;
            if view_score > top_view_score {
                top_view_score = view_score;
            }
        }
    }

    println!("p1: {}\np2: {}", visible_trees, top_view_score);
}