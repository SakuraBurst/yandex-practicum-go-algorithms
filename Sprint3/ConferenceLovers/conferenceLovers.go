package conferenceLovers

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func ConferenceLovers(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	scaner := bufio.NewScanner(reader)
	scaner.Scan()
	scaner.Scan()
	data := strings.Fields(scaner.Text())
	scaner.Scan()
	k, _ := strconv.Atoi(scaner.Text())
	hash := map[int]int{}
	for _, v := range data {
		vInt, _ := strconv.Atoi(v)
		hash[vInt]++
	}
	count := make([][2]int, 0, len(hash))
	for k, v := range hash {
		count = append(count, [2]int{k, v})
	}
	sortedCount := merge(count)
	writer := bufio.NewWriter(w)

	for i, v := range sortedCount {
		if i == k-1 {
			writer.WriteString(strconv.Itoa(v[0]))
			writer.Flush()
			return
		}
		writer.WriteString(strconv.Itoa(v[0]) + " ")

	}
}

func merge(d [][2]int) [][2]int {
	if len(d) == 1 {
		return d
	} else {
		a := merge(d[len(d)/2:])
		b := merge(d[:len(d)/2])
		aIndex := 0
		bIndex := 0
		result := make([][2]int, 0, len(a)+len(b))
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
			if compare(a[aIndex], b[bIndex]) {
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

func compare(a, b [2]int) bool {
	if a[1] == b[1] {
		return a[0] < b[0]
	}
	return a[1] > b[1]
}
