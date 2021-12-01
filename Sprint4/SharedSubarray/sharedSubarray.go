package sharedSubarray

import (
	"bufio"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func SharedSubarray(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	reader.ReadString('\n')
	nDataString, _ := reader.ReadString('\n')
	nData := strings.Fields(nDataString)
	reader.ReadString('\n')
	mDataString, _ := reader.ReadString('\n')
	mData := strings.Fields(mDataString)
	hashMap := map[string][]int{}
	for i, v := range nData {
		hashMap[v] = append(hashMap[v], i)
	}
	var potencialIndexes []int = nil
	currentMax := 0
	searchForIndex := -1
	max := 0
	result := func() {
		potencialIndexes = nil
		if currentMax > max {
			max = currentMax
		}
		currentMax = 0
		searchForIndex = -1
	}
	for _, v := range mData {
		if hashMap[v] != nil {
			if searchForIndex == -1 && potencialIndexes != nil {
				if si, ok := some(hashMap[v], potencialIndexes); ok {
					searchForIndex = si
					currentMax++
				} else {
					result()
					if len(hashMap[v]) == 1 {
						searchForIndex = hashMap[v][0]
						currentMax++
					} else {
						potencialIndexes = hashMap[v]
						currentMax++
					}
				}
			} else if searchForIndex == -1 {
				if len(hashMap[v]) == 1 {
					searchForIndex = hashMap[v][0]
					currentMax++
				} else {
					potencialIndexes = hashMap[v]
					currentMax++
				}
			} else {
				if r := binary(hashMap[v], searchForIndex+1); r != -1 {
					searchForIndex++
					currentMax++
				} else {
					result()
				}
			}
		} else {
			result()
		}
	}
	if currentMax > max {
		max = currentMax
	}
	writer := bufio.NewWriter(w)
	writer.WriteString(strconv.Itoa(max))
	writer.Flush()
}

func some(a, b []int) (int, bool) {
	for i, v := range a {
		if binary(b, v-1) != -1 {
			return i, true
		}
	}
	return -1, false
}

func binary(a []int, s int) int {
	logA := int(math.Log2(float64(len(a)))) + 2
	low := 0
	high := len(a)
	for i := 0; i < logA; i++ {
		middle := (low + high) >> 1
		if a[middle] == s {
			return middle
		}
		if a[middle] > s {
			high = middle
		} else {
			low = middle
		}
	}
	return -1
}

func main() {
	SharedSubarray(os.Stdin, os.Stdout)
}
