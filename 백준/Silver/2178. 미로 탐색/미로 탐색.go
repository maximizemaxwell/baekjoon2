// problem2137.go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func BFS(grid [][]int, start [2]int, end [2]int) int {
	n, m := len(grid), len(grid[0])
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	queue := [][3]int{{start[0], start[1], 1}} // [x, y, steps]
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}
	visited[start[0]][start[1]] = true

	for len(queue) > 0 {
		x, y, steps := queue[0][0], queue[0][1], queue[0][2]
		queue = queue[1:]

		if x == end[0] && y == end[1] {
			return steps
		}

		for _, dir := range directions {
			nx, ny := x+dir[0], y+dir[1]
			if nx >= 0 && nx < n && ny >= 0 && ny < m && grid[nx][ny] == 1 && !visited[nx][ny] {
				visited[nx][ny] = true
				queue = append(queue, [3]int{nx, ny, steps + 1})
			}
		}
	}

	return -1 // Should not reach here as per problem statement
}

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		var row string
		fmt.Fscanln(reader, &row)
		grid[i] = make([]int, m)
		for j := 0; j < m; j++ {
			grid[i][j] = int(row[j] - '0')
		}
	}

	start := [2]int{0, 0}
	end := [2]int{n - 1, m - 1}
	fmt.Fprintln(writer, BFS(grid, start, end))
}