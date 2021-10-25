package fib

import "fmt"

func Fib() {
	var n int
	fmt.Scan(&n)
	fmt.Println(nFibFor(n))
}

func nFibFor(n int) uint64 {
	results := [...]uint64{1, 1}
	for i := 0; i < n-2; i++ {
		results[0], results[1] = results[1], (results[0] + results[1])
	}
	return results[1]
}

func nFib(n int) int {
	if n < 2 {
		return n
	}
	return nFib(n-1) + nFib(n-2)
}
