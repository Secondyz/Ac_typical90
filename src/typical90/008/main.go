package main

import (
	"fmt"
	"strconv"
)

func main() {
	const MOD = 1000000007

	var N string
	var S string
	_, err := fmt.Scan(&N)
	if err != nil {
		return
	}
	_, err = fmt.Scan(&S)
	if err != nil {
		return
	}

	val, _ := strconv.Atoi(N)
	T := "atcoder"

	dp := make([][]int, val+1)
	for i := range dp {
		dp[i] = make([]int, 8)
	}
	dp[0][0] = 1

	for i := 0; i < val; i++ {
		for j := 0; j <= 7; j++ {
			dp[i+1][j] = (dp[i+1][j] + dp[i][j]) % MOD
			if j < 7 && S[i] == T[j] {
				dp[i+1][j+1] = (dp[i+1][j+1] + dp[i][j]) % MOD
			}
		}
	}

	fmt.Println(dp[val][7])
}
