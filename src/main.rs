use std::{env, fs::File, process};

fn main() {
    println!("Hello, world!");
    // read arguments from command line
    let args: Vec<String> = env::args().collect();
    // check if there are enough arguments
    if args.len() < 3 {
        println!("Please provide a file name");
        process::exit(1);
    }

    // read file
    let contents = read_file(&args[1]);
    println!("With text:\n{}", contents);
}

// function to read a file
fn read_file(input_file: &str) -> String {
    let mut file = File::open(input_file).expect("file not found");
    let mut contents = String::new();

    file.read_to_string(&mut contents)
        .expect("something went wrong reading the file");
    contents
}
