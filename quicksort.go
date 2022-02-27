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
			pivot := partition(array, low, high)

			quicksort(array, low, pivot)
			quicksort(array, pivot+1, high)
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
	// Hoare's partioning scheme
	pivot := array[low]
	left := low - 1
	right := high + 1

	for {
		for {
			left += 1
			if array[left] >= pivot {
				break
			}
		}

		for {
			right -= 1
			if array[right] <= pivot {
				break
			}
		}

		if left >= right {
			return right
		}

		swap(array, left, right)
	}
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
