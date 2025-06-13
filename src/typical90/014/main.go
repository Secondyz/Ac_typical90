package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func readNInts(scanner *bufio.Scanner, N int) []int {
	res := make([]int, 0, N)
	for len(res) < N {
		scanner.Scan()
		line := scanner.Text()
		fields := splitFields(line)
		for _, f := range fields {
			n, err := strconv.Atoi(f)
			if err != nil {
				panic(fmt.Sprintf("invalid int: %s", f))
			}
			res = append(res, n)
		}
	}
	return res
}

func splitFields(s string) []string {
	result := []string{}
	start := -1
	for i, c := range s {
		if c != ' ' && start == -1 {
			start = i
		}
		if c == ' ' && start != -1 {
			result = append(result, s[start:i])
			start = -1
		}
	}
	if start != -1 {
		result = append(result, s[start:])
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const maxCapacity = 1024 * 1024 * 10 // 最大10MB
	buf := make([]byte, 1024*1024)       // 初期バッファサイズ1MB
	scanner.Buffer(buf, maxCapacity)

	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	A := readNInts(scanner, N)
	B := readNInts(scanner, N)

	sort.Ints(A)
	sort.Ints(B)

	total := 0
	for i := 0; i < N; i++ {
		diff := A[i] - B[i]
		if diff < 0 {
			diff = -diff
		}
		total += diff
	}

	fmt.Println(total)
}
