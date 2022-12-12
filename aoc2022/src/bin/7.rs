extern crate core;

use std::collections::HashMap;
use std::path::{Path, PathBuf};

mod inputs;

fn main() {
    let pzl = inputs::read("7");

    let mut directories: HashMap<PathBuf, i32> = HashMap::new();
    let mut current_directory = PathBuf::new();

    for line in pzl.split("\n") {
        if line.starts_with("$ cd ") {
            let next_directory = line
                .split(" ")
                .nth(2)
                .unwrap();

            if next_directory == ".." {
                current_directory.pop();
            } else {
                current_directory.push(next_directory);
                directories.insert(current_directory.clone(), 0);
            }
        } else {
            let file_size = match line.split(" ").nth(0).unwrap().parse::<i32>() {
                Ok(n) => n,
                _ => continue,
            };

            directories
                .entry(current_directory.clone())
                .and_modify( |v: &mut i32| *v += file_size);
        }
    }

    let mut directory_paths = directories
        .keys()
        .cloned()
        .collect::<Vec<PathBuf>>();

    // sort path names by how many "/" they contain (descending order)
    // to guarantee that children are fully calculated before being added to parent
    directory_paths.sort_by(|a, b|
        b.to_str().unwrap().matches("/").count().cmp(
            &a.to_str().unwrap().matches("/").count()));

    // add each directory's file size to its direct parent's file size
    for path_buf in directory_paths {
        match path_buf.parent() {
            Some(parent) => {
                let file_size = *directories.get(&path_buf).unwrap();
                directories.entry(parent.to_path_buf()).and_modify( |v: &mut i32| *v += file_size);
            }
            _ => continue,
        }
    }

    let all_files_size = *directories.get(Path::new("/")).unwrap();
    let p2_target_file_size = 30000000 - (70000000 - all_files_size);

    let (mut p1, mut p2) = (0, i32::MAX);
    for (_, v) in &directories {

        if *v <= 100000 {
            p1 += v;
        }

        if *v >= p2_target_file_size && v < &p2 {
            p2 = *v;
        }
    }

    println!("p1: {}\np2: {}", p1, p2);
}