package fib

import "fmt"

func Fib() {
	var n int
	fmt.Scan(&n)
	fmt.Println(nFib(n + 1))
}

func nFib(n int) int {
	if n < 2 {
		return n
	}
	return nFib(n-1) + nFib(n-2)
}
