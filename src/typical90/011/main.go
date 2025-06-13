package main

import (
	"fmt"
	"sort"
)

type Job struct {
	D, C, S int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// solution
// 参考:ナップサック問題
// 締め切り順にソートする！ -> これ大事

func main() {
	var N int
	fmt.Scan(&N)

	jobs := make([]Job, N)
	maxDay := 0
	for i := 0; i < N; i++ {
		fmt.Scan(&jobs[i].D, &jobs[i].C, &jobs[i].S)
		if jobs[i].D > maxDay {
			maxDay = jobs[i].D
		}
	}

	// 締切で昇順ソート
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].D < jobs[j].D
	})

	dp := make([]int, maxDay+1)

	for _, job := range jobs {
		for day := job.D; day >= job.C; day-- {
			dp[day] = max(dp[day], dp[day-job.C]+job.S)
		}
	}

	ans := 0
	for _, v := range dp {
		ans = max(ans, v)
	}
	fmt.Println(ans)
}
