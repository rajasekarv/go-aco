use std::f32::INFINITY;
use std::io::BufRead;
use std::io::BufReader;
use std::io::{BufWriter, Write};

fn shortest_tour(
    nodes: Vec<Vec<f64>>,
    present_node: i64,
    visited_nodes: i64,
    lookupMatrix: Vec<Vec<f64>>,
) {
    if visited_nodes == (1 << nodes.len()) - 1 {
        return nodes[present_node][0];
    }
    if lookupMatrix[present_node][visited_nodes] != INFINITY {
        return lookupMatrix[present_node][visited_nodes];
    }

    for i in 0..nodes.len() {
        if (i == present_node) || ((visited_nodes & (1 << i)) > 0) {
            continue;
        }
        let new_distance = nodes[present_node][i]
            + shortest_tour(nodes, i, visited_nodes | (1 << i), lookupMatrix);

        if new_distance < lookupMatrix[present_node][visited_nodes] {
            lookupMatrix[present_node][visited_nodes] = new_distance;
        }
    }
    return lookupMatrix[present_node][visited_nodes];
}

fn main() {
    let distances: Vec<Vec<f64>> = Vec::new();
    let mut file = BufReader::with_capacity(
        1000 * 1024,
        File::open(
            "/home/raja/Downloads/programming/python/AntColonyOptimization-master/att48_d.txt",
        )
        .unwrap(),
    );
    let mut line = String::new();
    while file.read_line(&mut line).unwrap() > 0 {
        let words: Vec<f64> = line.split("      ").parse().collect();
        distances.push(words);
    }
    let distances1: Vec<Vec<f64>> = Vec::new();
    let lookupMatrix: Vec<Vec<f64>> = Vec::new();
    for i in 0..24 {
        let temp: Vec<f64> = Vec::new();
        for j in 0..24 {
            temp.push(distances[i][j])
        }
        distances1.push(temp);
    }

    for i in 0..distances1.len() {
        let temp: Vec<f64> = Vec::new();
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
        shortest_tour(distances1, 0, 1, lookupMatrix)
    );
}
