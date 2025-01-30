// problem7576.go
package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	X, Y int
}

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	var board [1002][1002]int
	var dist [1002][1002]int

	dx := [4]int{1, 0, -1, 0}
	dy := [4]int{0, 1, 0, -1}
	var queue []Point

	fmt.Fscanln(reader, &m, &n)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &board[i][j])
			if board[i][j] == 1 {
				queue = append(queue, Point{i, j})
			}
			if board[i][j] == 0 {
				dist[i][j] = -1
			}
		}
	}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for dir := 0; dir < 4; dir++ {
			nx, ny := cur.X+dx[dir], cur.Y+dy[dir]

			if nx < 0 || ny < 0 || nx >= n || ny >= m {
				continue
			}
			if dist[nx][ny] != -1 {
				continue
			}
			dist[nx][ny] = dist[cur.X][cur.Y] + 1
			queue = append(queue, Point{nx, ny})
		}
	}

	maxdays := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if dist[i][j] == -1 {
				fmt.Fprintln(writer, -1)
				return
			}
			if dist[i][j] > maxdays {
				maxdays = dist[i][j]
			}
		}
	}
	fmt.Fprintln(writer, maxdays)
}