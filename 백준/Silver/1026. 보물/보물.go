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

	// 배열 A와 B 입력받기
	arr1 := make([]int, n)
	arr2 := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr1[i])
	}

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr2[i])
	}

	sort.Sort(sort.IntSlice(arr1))
	sort.Sort(sort.Reverse(sort.IntSlice(arr2)))

	ans := 0
	for i := 0; i < n; i++ {
		ans += arr1[i] * arr2[i]
	}

	fmt.Fprintln(writer, ans)
}