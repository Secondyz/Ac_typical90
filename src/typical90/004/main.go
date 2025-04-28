package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calcRowSum(A [][]int) []int {
	H := len(A)
	W := len(A[0])
	rowSum := make([]int, H)
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			rowSum[i] += A[i][j]
		}
	}
	return rowSum
}

func calcColSum(A [][]int) []int {
	H := len(A)
	W := len(A[0])
	colSum := make([]int, W)
	for j := 0; j < W; j++ {
		for i := 0; i < H; i++ {
			colSum[j] += A[i][j]
		}
	}
	return colSum
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer func(writer *bufio.Writer) {
		err := writer.Flush()
		if err != nil {

		}
	}(writer)

	readInts := func() []int {
		line, _ := reader.ReadString('\n')
		fields := strings.Fields(line)
		res := make([]int, len(fields))
		for i, f := range fields {
			res[i], _ = strconv.Atoi(f)
		}
		return res
	}

	HW := readInts()
	H, W := HW[0], HW[1]

	A := make([][]int, H)
	for i := 0; i < H; i++ {
		A[i] = readInts()
	}

	row := calcRowSum(A)
	col := calcColSum(A)

	B := make([][]int, H)
	for i := 0; i < H; i++ {
		B[i] = make([]int, W)
		for j := 0; j < W; j++ {
			B[i][j] = row[i] + col[j] - A[i][j]
		}
	}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if j > 0 {
				err := writer.WriteByte(' ')
				if err != nil {
					return
				}
			}
			_, err := fmt.Fprint(writer, B[i][j])
			if err != nil {
				return
			}
		}
		err := writer.WriteByte('\n')
		if err != nil {
			return
		}
	}
}
