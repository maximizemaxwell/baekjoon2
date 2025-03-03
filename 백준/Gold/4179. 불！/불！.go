package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

const MAX = 1002

var (
	dx         = []int{1, 0, -1, 0}
	dy         = []int{0, 1, 0, -1}
	maze       [MAX][MAX]rune
	fireTime   [MAX][MAX]int
	jihoonTime [MAX][MAX]int
	r, c       int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &r, &c)

	fireQueue := []Point{}
	jihoonQueue := []Point{}

	for i := 0; i < r; i++ {
		var line string
		fmt.Fscan(reader, &line)
		for j := 0; j < c; j++ {
			maze[i][j] = rune(line[j])
			fireTime[i][j] = -1
			jihoonTime[i][j] = -1
			if maze[i][j] == 'J' {
				jihoonQueue = append(jihoonQueue, Point{i, j})
				jihoonTime[i][j] = 0
			} else if maze[i][j] == 'F' {
				fireQueue = append(fireQueue, Point{i, j})
				fireTime[i][j] = 0
			}
		}
	}

	// 불
	for len(fireQueue) > 0 {
		curr := fireQueue[0]
		fireQueue = fireQueue[1:]

		for d := 0; d < 4; d++ {
			nx, ny := curr.x+dx[d], curr.y+dy[d]

			if nx < 0 || nx >= r || ny < 0 || ny >= c || maze[nx][ny] == '#' || fireTime[nx][ny] >= 0 {
				continue
			}
			fireTime[nx][ny] = fireTime[curr.x][curr.y] + 1
			fireQueue = append(fireQueue, Point{nx, ny})
		}
	}

	// 지훈
	for len(jihoonQueue) > 0 {
		curr := jihoonQueue[0]
		jihoonQueue = jihoonQueue[1:]

		if curr.x == 0 || curr.x == r-1 || curr.y == 0 || curr.y == c-1 {
			fmt.Println(jihoonTime[curr.x][curr.y] + 1)
			return
		}

		for d := 0; d < 4; d++ {
			nx, ny := curr.x+dx[d], curr.y+dy[d]

			if nx < 0 || nx >= r || ny < 0 || ny >= c || maze[nx][ny] == '#' || jihoonTime[nx][ny] >= 0 {
				continue
			}
			if fireTime[nx][ny] != -1 && fireTime[nx][ny] <= jihoonTime[curr.x][curr.y]+1 {
				continue
			}
			jihoonTime[nx][ny] = jihoonTime[curr.x][curr.y] + 1
			jihoonQueue = append(jihoonQueue, Point{nx, ny})
		}
	}

	fmt.Println("IMPOSSIBLE")
}