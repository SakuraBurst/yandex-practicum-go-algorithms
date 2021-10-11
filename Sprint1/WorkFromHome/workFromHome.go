package workFromeHome

import (
	"fmt"
	"strconv"
)

func WorkFromeHome() {
	var n int
	fmt.Scan(&n)
	result := ""
	for n > 1 {
		result += strconv.Itoa(n % 2)
		n /= 2
	}
	result += strconv.Itoa(n)
	fmt.Println(string(byteStringReverse([]byte(result))))
}

func byteStringReverse(s []byte) []byte {
	for i, g := 0, len(s)-1; i < g; i, g = i+1, g-1 {
		s[i], s[g] = s[g], s[i]
	}
	return s
}
