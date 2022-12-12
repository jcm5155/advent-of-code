use std::fs::read_to_string;


pub fn read(filename: &str) -> String {
    let full_path = format!("src/bin/inputs/{}.input", filename);
    println!("input = {}", full_path);
    let pzl = read_to_string(full_path)
        .expect("input file not found");
    return String::from(pzl.trim());
}
