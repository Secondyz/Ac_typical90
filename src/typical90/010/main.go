package main

import (
	"fmt"
	"io"
	"os"
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
	class1Sum := make([]int, N+1)
	class2Sum := make([]int, N+1)

	for i := 0; i < N; i++ {
		class := readInt()
		score := readInt()

		// 累積和を前から順に作る
		class1Sum[i+1] = class1Sum[i]
		class2Sum[i+1] = class2Sum[i]

		if class == 1 {
			class1Sum[i+1] += score
		} else if class == 2 {
			class2Sum[i+1] += score
		}
	}

	Q := readInt()
	var sb strings.Builder
	for i := 0; i < Q; i++ {
		L := readInt()
		R := readInt()

		c1 := class1Sum[R] - class1Sum[L-1]
		c2 := class2Sum[R] - class2Sum[L-1]

		sb.WriteString(fmt.Sprintf("%d %d\n", c1, c2))
	}

	fmt.Print(sb.String())
}
