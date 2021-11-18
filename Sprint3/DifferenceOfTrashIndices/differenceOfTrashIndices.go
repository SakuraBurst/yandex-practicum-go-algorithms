package differenceOfTrashIndices

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func DifferenceOfTrashIndices(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	nString, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nString))
	stringData, _ := reader.ReadString('\n')
	data := strings.Fields(stringData)
	stringData = ""
	kString, _ := reader.ReadString('\n')
	k, _ := strconv.Atoi(strings.TrimSpace(kString))
	intData := merge(data)
	data = nil
	lowest := 0
	higest := intData[len(intData)-1] - intData[0]
	for lowest < higest {
		mid := (lowest + higest) >> 1
		if countPairs2(intData, n, mid) < k {
			lowest = mid + 1
		} else {
			higest = mid
		}
	}
	writer := bufio.NewWriter(w)
	writer.WriteString(strconv.Itoa(lowest))
	writer.Flush()

}

// returns number of pairs with absolute
// difference less than or equal to mid

func countPairs2(a []int, n, mid int) int {
	count, left := 0, 0
	for i := 0; i < n; i++ {
		for (a[i] - a[left]) > mid {
			left++
		}
		count += (i - left)
	}
	return count
}

func main() {
	DifferenceOfTrashIndices(os.Stdin, os.Stdout)
}
func merge(d []string) []int {
	if len(d) == 1 {
		i, _ := strconv.Atoi(d[0])
		r := make([]int, 1)
		r[0] = i
		return r
	}
	a := merge(d[len(d)/2:])
	b := merge(d[:len(d)/2])
	aIndex := 0
	bIndex := 0
	result := make([]int, 0, len(a)+len(b))
	for aIndex != len(a) || bIndex != len(b) {
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
