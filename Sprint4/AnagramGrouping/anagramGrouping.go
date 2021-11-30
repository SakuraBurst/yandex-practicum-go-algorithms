package anagramGrouping

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func AnagramGrouping(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	reader.ReadString('\n')
	s, _ := reader.ReadString('\n')
	sData := strings.Fields(s)
	hashMap := map[int][]int{}
	for i, v := range sData {
		hashMap[hash(v)] = append(hashMap[hash(v)], i)
	}
	resultSlice := make([][]int, 0, len(hashMap))
	for _, v := range hashMap {
		resultSlice = append(resultSlice, v)
	}
	resultSlice = simpleMerge(resultSlice)
	writer := bufio.NewWriter(w)
	for _, rv := range resultSlice {
		for i, v := range rv {
			writer.WriteString(strconv.Itoa(v))
			if len(rv)-1 > i {
				writer.WriteString(" ")
			}
		}
		writer.WriteByte('\n')

	}
	writer.Flush()
}

func simpleMerge(d [][]int) [][]int {
	if len(d) == 1 {
		return d
	}

	mid := len(d) >> 1

	a := simpleMerge(d[:mid])
	b := simpleMerge(d[mid:])
	aIndex := 0
	bIndex := 0
	result := make([][]int, 0, len(a)+len(b))
	for aIndex < len(a) || bIndex < len(b) {
		if aIndex == len(a) {
			result = append(result, b[bIndex])
			bIndex++
			continue
		}
		if bIndex == len(b) {
			result = append(result, a[aIndex])
			aIndex++
			continue
		}
		if a[aIndex][0] <= b[bIndex][0] {
			result = append(result, a[aIndex])
			aIndex++
		} else {
			result = append(result, b[bIndex])
			bIndex++
		}
	}
	return result
}

func hash(s string) int {
	hash := 1
	for i := 0; i < len(s); i++ {
		hash = hash * int(s[i]) % 1000007
	}
	return hash
}

func main() {
	AnagramGrouping(os.Stdin, os.Stdout)
}
