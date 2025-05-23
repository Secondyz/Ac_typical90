package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Edge struct {
	to, cost int
}
type Item struct {
	node, dist int
}
type PriorityQueue []Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Item))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

func dijkstra(N, start int, graph [][]Edge) []int {
	dist := make([]int, N+1)
	for i := range dist {
		dist[i] = math.MaxInt64
	}
	dist[start] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Item{start, 0})

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(Item)
		if cur.dist > dist[cur.node] {
			continue
		}
		for _, e := range graph[cur.node] {
			if dist[e.to] > dist[cur.node]+e.cost {
				dist[e.to] = dist[cur.node] + e.cost
				heap.Push(pq, Item{e.to, dist[e.to]})
			}
		}
	}
	return dist
}

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	graph := make([][]Edge, N+1)
	for i := 0; i < M; i++ {
		var a, b, c int
		fmt.Scan(&a, &b, &c)
		graph[a] = append(graph[a], Edge{b, c})
		graph[b] = append(graph[b], Edge{a, c}) // 双方向
	}

	distFrom1 := dijkstra(N, 1, graph)
	distFromN := dijkstra(N, N, graph)

	for k := 1; k <= N; k++ {
		fmt.Println(distFrom1[k] + distFromN[k])
	}
}
