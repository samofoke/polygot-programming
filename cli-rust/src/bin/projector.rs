use clap::Parser;
use cli_rust::opts::Opts;

fn main() {
    let opts = Opts::parse();
    println!("{:?}", opts);
}
