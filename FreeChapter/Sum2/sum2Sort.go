package sum2

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type stringsNumberSort []string

func (a stringsNumberSort) Len() int      { return len(a) }
func (a stringsNumberSort) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a stringsNumberSort) Less(i, j int) bool {
	firstItem, _ := strconv.ParseInt(strings.TrimSpace(a[i]), 10, 64)
	secondItem, _ := strconv.ParseInt(strings.TrimSpace(a[j]), 10, 64)
	return firstItem < secondItem
}

func Sum2Sort() {
	stdReader := bufio.NewReader(os.Stdin)
	stdWriter := bufio.NewWriter(os.Stdout)
	nLine, _ := stdReader.ReadString('\n')
	n, _ := strconv.ParseInt(strings.TrimSpace(nLine), 10, 64)
	arrayLine, _ := stdReader.ReadString('\n')
	splited := strings.Split(strings.TrimSpace(arrayLine), " ")
	sLine, _ := stdReader.ReadString('\n')
	s, _ := strconv.ParseInt(strings.TrimSpace(sLine), 10, 64)
	sort.Sort(stringsNumberSort(splited))
	var i int64
	var g int64
	for i, g = 0, n-1; i < g; {
		firstItem, _ := strconv.ParseInt(strings.TrimSpace(splited[i]), 10, 64)
		secondItem, _ := strconv.ParseInt(strings.TrimSpace(splited[g]), 10, 64)
		currSum := firstItem + secondItem
		if currSum > s {
			g--
			continue
		}
		if currSum < s {
			i++
			continue
		}
		fmt.Fprintf(stdWriter, "%d %d", firstItem, secondItem)
		break
	}
	if stdWriter.Buffered() == 0 {
		fmt.Fprint(stdWriter, "None")
	}
	stdWriter.Flush()

}
