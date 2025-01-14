package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Pair struct {
	Start int
	End   int
}

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)
	meetings := make([]Pair, n)

	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &meetings[i].Start, &meetings[i].End)
	}
	sort.Slice(meetings, func(i, j int) bool {
		if meetings[i].End == meetings[j].End {
			return meetings[i].Start < meetings[j].Start
		}
		return meetings[i].End < meetings[j].End
	})

	count := 0
	lastEnd := 0

	for _, meeting := range meetings {
		if meeting.Start >= lastEnd {
			count++
			lastEnd = meeting.End
		}
	}

	fmt.Fprintln(writer, count)
}