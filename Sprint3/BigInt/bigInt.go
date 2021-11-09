package bigInt

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func BigInt(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	scaner := bufio.NewScanner(reader)
	scaner.Scan()
	n, _ := strconv.Atoi(scaner.Text())
	scaner.Scan()
	numbers := strings.Fields(scaner.Text())
	for i := 1; i < n; i++ {
		for g := i; g > 0 && Comapre(numbers[g], numbers[g-1]); g-- {
			numbers[g], numbers[g-1] = numbers[g-1], numbers[g]
		}
	}
	writer := bufio.NewWriter(w)
	writer.Write([]byte(strings.Join(numbers, "")))
	writer.Flush()

}

func Comapre(a, b string) bool {
	aInt, _ := strconv.Atoi(a + b)
	bInt, _ := strconv.Atoi(b + a)
	return aInt > bInt
}
