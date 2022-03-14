package shiftSearch

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"unicode"
)

func ShiftSearch(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	data := stringToIntSlice(n, scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	searchFor := stringToIntSlice(m, scanner.Text())
	offset := 0
	ol := n - m + 1
Outer:
	for i := 0; i < ol; i++ {

		offset = searchFor[0] - data[i]
		for g := 1; g < m; g++ {
			if data[i+g]+offset != searchFor[g] {
				continue Outer
			}
		}
		io.WriteString(w, strconv.Itoa(i+1)+" ")

	}
}

func stringToIntSlice(l int, s string) []int {
	data := make([]int, 0, l)
	start := 0
	for i, r := range s {
		if !unicode.IsSpace(r) {
			continue
		}
		d, _ := strconv.Atoi(s[start:i])
		data = append(data, d)
		start = i + 1
	}
	d, _ := strconv.Atoi(s[start:])
	data = append(data, d)
	return data
}

func main() {
	ShiftSearch(os.Stdin, os.Stdout)
}
