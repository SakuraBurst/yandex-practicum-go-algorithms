package differenceOfTrashIndices

import (
	"bufio"
	"io"
	"math"
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
	kString, _ := reader.ReadString('\n')
	k, _ := strconv.Atoi(strings.TrimSpace(kString))
	intData := make([]int, n)
	for i, v := range data {
		iV, _ := strconv.Atoi(v)
		intData[i] = iV
	}
	minDiffs := make([]int, 0, k)
	for i := 0; i < n; i++ {
		for g := i + 1; g < n; g++ {
			r := int(math.Abs(float64(intData[i] - intData[g])))
			if len(minDiffs) < k {
				minDiffs = append(minDiffs, r)
				tupichSort(minDiffs)
			} else {

				if r < minDiffs[k-1] {
					insertLow(minDiffs, r)

				} else {
					continue
				}

			}
		}
	}
	writer := bufio.NewWriter(w)
	writer.WriteString(strconv.Itoa(minDiffs[k-1]))
	writer.Flush()
}

func insertLow(d []int, insert int) {
	for i := len(d) - 1; i >= 0; i-- {
		if insert < d[i] {
			if i == len(d)-1 {
				continue
			} else {
				d[i+1] = d[i]
			}
		} else {
			d[i+1] = insert
			return
		}
	}
	if insert < d[0] {
		d[0] = insert
		return
	}

}

func tupichSort(d []int) {
	for i := 1; i < len(d); i++ {
		for g := i; g > 0 && d[g] < d[g-1]; g-- {
			d[g], d[g-1] = d[g-1], d[g]
		}
	}
}

func main() {
	DifferenceOfTrashIndices(os.Stdin, os.Stdout)
}
