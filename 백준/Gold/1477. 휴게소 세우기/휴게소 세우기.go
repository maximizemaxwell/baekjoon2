package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m, l int
	fmt.Fscanln(reader, &n, &m, &l)

	restAreas := make([]int, n)
	if n > 0 {
		for i := 0; i < n; i++ {
			fmt.Fscan(reader, &restAreas[i])
		}
	}

	sort.Ints(restAreas)

	restAreas = append([]int{0}, restAreas...)
	restAreas = append(restAreas, l)

	left, right := 1, l
	result := right

	for left <= right {
		mid := (left + right) / 2
		count := 0

		for i := 1; i < len(restAreas); i++ {
			section := restAreas[i] - restAreas[i-1]
			if section > mid {
				count += (section - 1) / mid
			}
		}

		if count <= m {
			result = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	fmt.Fprintln(writer, result)
}