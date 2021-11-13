package housebuying

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func HouseBuying(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	scaner := bufio.NewScanner(reader)
	scaner.Scan()
	initData := strings.Fields(scaner.Text())
	money, _ := strconv.Atoi(initData[1])
	scaner.Scan()
	housePrices := strings.Fields(scaner.Text())
	housePricesSorted := merge(housePrices)
	index := 0
	for money >= 0 {
		if len(housePricesSorted) == index {
			index++
			break
		}
		money -= housePricesSorted[index]
		index++
	}
	writer := bufio.NewWriter(w)
	writer.WriteString(strconv.Itoa(index - 1))
	writer.Flush()
}

func merge(s []string) []int {
	if len(s) == 1 {
		intSlice := make([]int, 1)
		i, _ := strconv.Atoi(s[0])
		intSlice[0] = i
		return intSlice
	} else {
		mid := len(s) / 2
		a := merge(s[:mid])
		b := merge(s[mid:])
		aIndex := 0
		bIndex := 0
		result := make([]int, 0, len(a)+len(b))
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
			if a[aIndex] <= b[bIndex] {
				result = append(result, a[aIndex])
				aIndex++
			} else {
				result = append(result, b[bIndex])
				bIndex++
			}
		}
		return result
	}
}
