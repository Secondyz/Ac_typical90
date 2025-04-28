package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func generate(n int, open int, close int, s string, result *[]string) {
	if len(s) == n {
		*result = append(*result, s)
		return
	}
	if open < n/2 {
		generate(n, open+1, close, s+"(", result)
	}
	if close < open {
		generate(n, open, close+1, s+")", result)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	if N%2 != 0 {
		fmt.Println()
		return
	}

	var result []string
	generate(N, 0, 0, "", &result)

	for _, s := range result {
		fmt.Println(s)
	}
}
