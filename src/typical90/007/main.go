package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	inputBytes, _ := io.ReadAll(os.Stdin)
	tokens := strings.Fields(string(inputBytes))

	idx := 0
	readInt := func() int {
		val, _ := strconv.Atoi(tokens[idx])
		idx++
		return val
	}

	N := readInt()
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = readInt()
	}
	sort.Ints(A)

	Q := readInt()
	output := make([]string, Q)
	for i := 0; i < Q; i++ {
		b := readInt()

		// 二分探索
		pos := sort.SearchInts(A, b)

		//最小値判定
		minDiff := 1 << 60
		if pos < len(A) {
			minDiff = abs(A[pos] - b)
		}
		if pos > 0 {
			diff := abs(A[pos-1] - b)
			if diff < minDiff {
				minDiff = diff
			}
		}
		output[i] = strconv.Itoa(minDiff)
	}

	fmt.Println(strings.Join(output, "\n"))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
