use std::f64::INFINITY;
use std::fs::File;
use std::io::BufRead;
use std::io::BufReader;
use std::io::{BufWriter, Write};

fn shortest_tour(
    nodes: &mut Vec<Vec<f64>>,
    present_node: i64,
    visited_nodes: i64,
    lookupMatrix: &mut Vec<Vec<f64>>,
) -> f64 {
    if visited_nodes == (1 << nodes.len()) - 1 {
        return nodes[present_node as usize][0];
    }
    if lookupMatrix[present_node as usize][visited_nodes as usize] != INFINITY {
        return lookupMatrix[present_node as usize][visited_nodes as usize];
    }

    for i in 0..nodes.len() {
        if (i == present_node as usize) || ((visited_nodes & (1 << i)) > 0) {
            continue;
        }
        let new_distance = nodes[present_node as usize][i as usize]
            + shortest_tour(nodes, i as i64, visited_nodes | (1 << i), lookupMatrix);

        if new_distance < lookupMatrix[present_node as usize][visited_nodes as usize] {
            lookupMatrix[present_node as usize][visited_nodes as usize] = new_distance;
        }
    }
    return lookupMatrix[present_node as usize][visited_nodes as usize];
}

fn main() {
    let mut distances: Vec<Vec<f64>> = Vec::new();
    let mut file = BufReader::with_capacity(
        8 * 1024,
        File::open(
            "/home/raja/Downloads/programming/python/AntColonyOptimization-master/att48_d.txt",
        )
        .unwrap(),
    );
    println!("read_file");
    let mut line = String::new();
    //let mut count = 0;
    while file
        .read_line({
            line.clear();
            &mut line
        })
        .unwrap()
        > 0
    {
        //println!("{:?}", line);
        //println!("{:?}", count);
        //count += 1;
        let words: Vec<f64> = line.split("      ").skip(1).map(|s| s.trim().parse().unwrap()).collect();
        //println!("{:?}", words);
        //let words: Vec<f64> = line.split("      ").skip(1).map(|s| s.parse().unwrap()).collect();
        distances.push(words);
    }
    let mut distances1: Vec<Vec<f64>> = Vec::new();
    let mut lookupMatrix: Vec<Vec<f64>> = Vec::new();
    for i in 0..25 {
        let mut temp: Vec<f64> = Vec::new();
        for j in 0..25 {
            temp.push(distances[i][j])
        }
        distances1.push(temp);
    }

    for i in 0..distances1.len() {
        let mut temp: Vec<f64> = Vec::new();
        for j in 0..((1 << distances1.len()) - 1) {
            temp.push(INFINITY)
        }
        lookupMatrix.push(temp);
    }

    println!("nodes size: {} {}", distances1.len(), distances.len());
    println!(
        "lookup size: {} {}",
        lookupMatrix.len(),
        lookupMatrix[0].len()
    );
    println!(
        "shortestTour {}",
        shortest_tour(&mut distances1, 0, 1, &mut lookupMatrix)
    );
}
