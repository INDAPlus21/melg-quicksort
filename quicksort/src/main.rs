use std::io;
use std::io::BufRead;

fn main() {
    // Input
    let mut line = String::with_capacity(2_000_000); // 500000 chars
    io::stdin().lock().read_line(&mut line);

    let mut values: Vec<isize> = Vec::with_capacity(line.len());
    values = line
        .split_whitespace()
        .skip(1)
        .map(|_val| _val.parse::<isize>().unwrap())
        .collect();
    let length = values.len();

    // Sort
    quicksort(&mut values, 0, length - 1);

    // Print
    let mut output = String::with_capacity(line.len());
    for val in values {
        output.push_str(&val.to_string());
        output.push_str(" ");
    }

    print!("{}", output.to_string())
}

fn quicksort(values: &mut Vec<isize>, low: usize, high: usize) {
    // Tail call optimization
    let mut new_low = low;
    while new_low < high {
        // Use insertion sort for small arrays
        if high - new_low <= 10 {
            insertionsort(values, new_low, high);
            break;
        } else {
            // Use quick sort
            let pivot = partition(values, new_low, high);
            quicksort(values, new_low, pivot);
            new_low = pivot + 1;
        }
    }
}

// Hoare's partioning scheme
fn partition(values: &mut Vec<isize>, low: usize, high: usize) -> usize {
    let pivot = values[low];
    let mut left = low - 1;
    let mut right = high + 1;

    loop {
        loop {
            left += 1;
            if values[left] >= pivot {
                break;
            }
        }

        loop {
            right -= 1;
            if values[right] <= pivot {
                break;
            }
        }

        if left >= right {
            return right;
        }

        values.swap(left, right);
    }
}

fn insertionsort(values: &mut Vec<isize>, low: usize, high: usize) {
    for i in 0..(high + 1) {
        let value = values[i];
        let mut j = i;

        while j > low && values[j - 1] > value {
            values[j] = values[j - 1];
            j -= 1;
        }

        values[j] = value;
    }
}
