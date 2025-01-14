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

	var n int
	fmt.Fscanln(reader, &n)

	ropes := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &ropes[i])
	}

	sort.Slice(ropes, func(i, j int) bool {
		return ropes[i] > ropes[j]
	})

	maxWeight := 0
	for i := 0; i < n; i++ {
		weight := ropes[i] * (i + 1)
		if weight > maxWeight {
			maxWeight = weight
		}
	}

	fmt.Fprintln(writer, maxWeight)
}