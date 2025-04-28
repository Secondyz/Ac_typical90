package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func bfs(start int, graph [][]int) (int, int) {
	n := len(graph)
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = -1
	}
	dist[start] = 0

	queue := []int{start}

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		for _, to := range graph[v] {
			if dist[to] == -1 {
				dist[to] = dist[v] + 1
				queue = append(queue, to)
			}
		}
	}

	maxDist := 0
	node := start
	for i := 0; i < n; i++ {
		if dist[i] > maxDist {
			maxDist = dist[i]
			node = i
		}
	}

	return node, maxDist
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	nextInt := func() int {
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		return v
	}

	N := nextInt()
	graph := make([][]int, N)

	for i := 0; i < N-1; i++ {
		a := nextInt() - 1
		b := nextInt() - 1
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	u, _ := bfs(0, graph)
	_, diameter := bfs(u, graph)

	fmt.Println(diameter + 1)
}
