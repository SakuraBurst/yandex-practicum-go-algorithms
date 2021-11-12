package cookies

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func Cookies(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	scaner := bufio.NewScanner(reader)
	scaner.Scan()
	// nChildren, _ := strconv.Atoi(scaner.Text())
	scaner.Scan()
	childrenGreed := strings.Fields(scaner.Text())
	scaner.Scan()
	// nCookies, _ := strconv.Atoi(scaner.Text())
	scaner.Scan()
	cookies := strings.Fields(scaner.Text())
	childrenGreedSorted := merge(childrenGreed)
	cookiesSorted := merge(cookies)
	cokkiesIndex := 0
	happyChildrenCounter := 0
	for _, v := range childrenGreedSorted {
		if cokkiesIndex >= len(cookiesSorted) {
			break
		}
		for ; cokkiesIndex < len(cookiesSorted); cokkiesIndex++ {
			if cookiesSorted[cokkiesIndex] >= v {
				happyChildrenCounter++
				cokkiesIndex++
				break
			}
		}
	}
	writer := bufio.NewWriter(w)
	writer.WriteString(strconv.Itoa(happyChildrenCounter))
	writer.Flush()
}

func merge(s []string) []int {

	if len(s) == 1 {
		i, _ := strconv.Atoi(s[0])
		result := []int{i}
		return result
	}
	middle := len(s) / 2
	a := merge(s[:middle])
	b := merge(s[middle:])

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
