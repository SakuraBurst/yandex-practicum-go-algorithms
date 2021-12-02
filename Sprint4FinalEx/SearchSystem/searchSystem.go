package searchSystem

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type WordLocation struct {
	docIndex  int
	wordIndex int
}

func SearchSystem(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	scaner := bufio.NewScanner(reader)
	scaner.Scan()
	nDocks, _ := strconv.Atoi(scaner.Text())
	searchIndexesMap := map[string][]WordLocation{}
	for d := 0; d < nDocks; d++ {
		scaner.Scan()
		for w, v := range strings.Fields(scaner.Text()) {
			searchIndexesMap[v] = append(searchIndexesMap[v], WordLocation{docIndex: d, wordIndex: w})
		}
	}
	scaner.Scan()
	kSearchStrings, _ := strconv.Atoi(scaner.Text())
	writer := bufio.NewWriter(w)
	// mapOfResults = map[string]map[int]int{}
	for s := 0; s < kSearchStrings; s++ {
		scaner.Scan()
		result := make(map[int]int, len(scaner.Text()))
		alreadyCounted := make(map[string]bool)
		for _, w := range strings.Fields(scaner.Text()) {
			if !alreadyCounted[w] {
				for _, wl := range searchIndexesMap[w] {
					result[wl.docIndex]++
				}
				alreadyCounted[w] = true
			}

		}
		counter := 0
		for k := range result {
			if counter == 5 {
				break
			}
			counter++
			writer.WriteString(strconv.Itoa(k + 1))
			writer.WriteByte(' ')
		}
		writer.WriteByte('\n')

	}
	writer.Flush()
	//
}

func main() {
	SearchSystem(os.Stdin, os.Stdout)
}
