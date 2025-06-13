package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 1_000_000_007
const MAX = 200000

var fact [MAX + 1]int64
var invFact [MAX + 1]int64

func modpow(a, b, m int64) int64 {
	res := int64(1)
	x := a % m
	e := b
	for e > 0 {
		if e&1 == 1 {
			res = (res * x) % m
		}
		x = (x * x) % m
		e >>= 1
	}
	return res
}

func modinv(a, m int64) int64 {
	return modpow(a, m-2, m)
}

func initFact() {
	fact[0] = 1
	for i := 1; i <= MAX; i++ {
		fact[i] = (fact[i-1] * int64(i)) % MOD
	}
	invFact[MAX] = modinv(fact[MAX], MOD)
	for i := MAX - 1; i >= 0; i-- {
		invFact[i] = (invFact[i+1] * int64(i+1)) % MOD
	}
}

func nCr(n, r int) int64 {
	if n < r || r < 0 {
		return 0
	}
	return (fact[n] * ((invFact[r] * invFact[n-r]) % MOD)) % MOD
}

func query(N, K int) int64 {
	var ret int64 = 0
	for i := 1; i <= N/K+1; i++ {
		s1 := N - (K-1)*(i-1)
		s2 := i
		ret += nCr(s1, s2)
		ret %= MOD
	}
	return ret
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var N int
	fmt.Sscanf(scanner.Text(), "%d", &N)

	initFact()

	for k := 1; k <= N; k++ {
		ans := query(N, k)
		fmt.Println(ans)
	}
}
