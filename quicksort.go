// go build -o C:\Users\Marcus\Documents\Github\melg-quicksort\out
// out/quicksort.exe
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var count int
	var values []int

	// Input
	const maxCapacity = 2000000 // Max 500000 i32
	buffer := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buffer, maxCapacity)

	var temp string
	scanner.Scan()

	temp = scanner.Text()
	inputs := strings.Split(temp, " ")
	count, _ = strconv.Atoi(inputs[0])

	for i := 0; i < count; i++ {
		val, _ := strconv.Atoi(inputs[i+1])

		// Values
		values = append(values, val)
	}

	// Sort
	quicksort(values, 0, len(values)-1)

	// Print
	var output strings.Builder
	for _, j := range values {
		output.WriteString(strconv.Itoa(j))
		output.WriteString(" ")
	}

	fmt.Print(output.String())
}

// Based on https://www.geeksforgeeks.org/quick-sort/
func quicksort(array []int, low int, high int) {
	if low < high {
		// Use insertion sort for small arrays
		if high-low <= 10 {
			insertionsort(array, low, high)
		} else {
			// Use quick sort
			partionindex := partition(array, low, high)

			quicksort(array, low, partionindex-1)  // Left of index
			quicksort(array, partionindex+1, high) // Right of index
		}
	}
}

func swap(array []int, a int, b int) {
	tmp := array[a]
	array[a] = array[b]
	array[b] = tmp
}

// Place values that are less to the left and values that are higher to the right
func partition(array []int, low int, high int) int {
	pivot := array[high]
	index := low - 1

	for j := low; j < high; j++ {
		if array[j] < pivot {
			index += 1

			swap(array, index, j)
		}
	}

	swap(array, index+1, high)

	return index + 1
}

func insertionsort(array []int, low int, high int) {
	for i := 0; i < high+1; i++ {
		value := array[i]
		j := i

		// Move the column until it's in the right place
		for j > low && array[j-1] > value {
			array[j] = array[j-1]
			j -= 1
		}

		array[j] = value
	}
}
