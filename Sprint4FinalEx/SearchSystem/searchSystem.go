package searchSystem

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type SearchResult struct {
	index      int
	wordsFound int
}

func SearchSystem(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	scaner := bufio.NewScanner(reader)
	scaner.Scan()
	nDocks, _ := strconv.Atoi(scaner.Text())
	searchIndexesMap := map[string]map[int]int{}
	for d := 0; d < nDocks; d++ {
		scaner.Scan()
		for _, v := range strings.Fields(scaner.Text()) {
			if searchIndexesMap[v] == nil {
				searchIndexesMap[v] = map[int]int{}
			}
			searchIndexesMap[v][d]++
		}
	}
	scaner.Scan()
	kSearchStrings, _ := strconv.Atoi(scaner.Text())
	writer := bufio.NewWriter(w)
	for s := 0; s < kSearchStrings; s++ {
		scaner.Scan()
		result := make(map[int]int, len(scaner.Text()))
		alreadyCounted := make(map[string]bool)
		for _, w := range strings.Fields(scaner.Text()) {
			if !alreadyCounted[w] {
				for k, wl := range searchIndexesMap[w] {
					result[k] += wl
				}
				alreadyCounted[w] = true
			}

		}
		counterArr := make([]SearchResult, 0, len(result))
		for k, v := range result {
			counterArr = append(counterArr, SearchResult{index: k + 1, wordsFound: v})
		}
		counterArr = merge(counterArr)
		for i, v := range counterArr {
			if i == 5 {
				break
			}
			writer.WriteString(strconv.Itoa(v.index))
			writer.WriteByte(' ')
		}
		writer.WriteByte('\n')

	}
	writer.Flush()
	//
}

func merge(d []SearchResult) []SearchResult {
	if len(d) <= 1 {
		return d
	}
	mid := len(d) >> 1
	firstHalf := merge(d[:mid])
	secondHalf := merge(d[mid:])
	firstHalfIndex := 0
	secondHalfIndex := 0
	result := make([]SearchResult, 0, len(firstHalf)+len(secondHalf))
	for firstHalfIndex < len(firstHalf) || secondHalfIndex < len(secondHalf) {
		// если закончилась первая плоловина
		if firstHalfIndex == len(firstHalf) {
			result = append(result, secondHalf[secondHalfIndex])
			secondHalfIndex++
			continue
		}
		// если закончилась вторая половина
		if secondHalfIndex == len(secondHalf) {
			result = append(result, firstHalf[firstHalfIndex])
			firstHalfIndex++
			continue
		}
		// если обе половины что-то содержат
		if compareResults(firstHalf[firstHalfIndex], secondHalf[secondHalfIndex]) {
			result = append(result, firstHalf[firstHalfIndex])
			firstHalfIndex++
		} else {
			result = append(result, secondHalf[secondHalfIndex])
			secondHalfIndex++
		}
	}
	return result
}

func compareResults(a, b SearchResult) bool {
	if a.wordsFound == b.wordsFound {
		return a.index < b.index
	}
	return a.wordsFound > b.wordsFound
}

func main() {
	SearchSystem(os.Stdin, os.Stdout)
}
