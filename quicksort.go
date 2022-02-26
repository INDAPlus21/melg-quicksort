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
	const maxCapacity = 512 * 1024 * 8
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
		partionindex := partition(array, low, high)

		quicksort(array, low, partionindex-1)  // Left of index
		quicksort(array, partionindex+1, high) // Right of index
	}
}

// Place values that are less to the left and values that are higher to the right
func partition(array []int, low int, high int) int {
	pivot := array[high]
	index := low - 1

	for j := low; j < high; j++ {
		if array[j] < pivot {
			index += 1

			// Swap
			temp := array[index]
			array[index] = array[j]
			array[j] = temp
		}
	}

	// Swap
	temp := array[index+1]
	array[index+1] = array[high]
	array[high] = temp

	return index + 1
}
