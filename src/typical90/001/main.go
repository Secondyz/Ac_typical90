package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	nextInt := func() int {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		return n
	}

	N := nextInt()
	L := nextInt()
	K := nextInt()

	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = nextInt()
	}

	low := 1
	high := L

	isPossible := func(x int) bool {
		count := 0
		prev := 0
		for _, a := range A {
			if a-prev >= x {
				count++
				prev = a
			}
		}
		if L-prev >= x {
			count++
		}
		return count >= K+1
	}

	for high-low > 1 {
		mid := (low + high) / 2
		if isPossible(mid) {
			low = mid
		} else {
			high = mid
		}
	}

	fmt.Println(low)
}
