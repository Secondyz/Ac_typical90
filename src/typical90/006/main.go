package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	buf := make([]byte, 4096)
	scanner.Buffer(buf, 2000000)

	nextInt := func() int {
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		return v
	}

	nextString := func() string {
		scanner.Scan()
		return scanner.Text()
	}

	N := nextInt()
	K := nextInt()
	S := nextString()

	var result strings.Builder
	pos := 0

	for i := 0; i < K; i++ {
		minChar := byte('z' + 1)
		minIdx := -1

		// S[j]の境界地付近でエラーが出る？？？？
		for j := pos; j <= N-(K-i); j++ {
			if S[j] < minChar {
				minChar = S[j]
				minIdx = j
			}
		}

		result.WriteByte(minChar)
		pos = minIdx + 1

	}
	fmt.Println(result.String())
}
