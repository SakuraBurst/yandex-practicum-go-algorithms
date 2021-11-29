package sumOfFour

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type SemiResult struct {
	firstValue       int
	secondValue      int
	firstValueIndex  int
	secondValueIndex int
}

func SumOfFour(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	scaner := bufio.NewScanner(reader)
	scaner.Scan()
	n, _ := strconv.Atoi(scaner.Text())
	writer := bufio.NewWriter(w)
	if n == 0 {
		writer.WriteString("0")
		writer.WriteByte('\n')
		writer.Flush()
		return
	}
	scaner.Scan()
	s, _ := strconv.Atoi(scaner.Text())
	scaner.Scan()
	stringData := strings.Fields(scaner.Text())
	data := merge(stringData)
	stringData = nil
	semiResultsHash := map[int][]SemiResult{}
	for i := 0; i < n; i++ {
		for g := i + 1; g < n; g++ {
			semiResultsHash[data[g]+data[i]] = append(semiResultsHash[data[g]+data[i]], SemiResult{firstValue: data[i], secondValue: data[g], firstValueIndex: i, secondValueIndex: g})
		}
	}
	// fmt.Println(semiResultsHash)
	resultsHash := map[string]bool{}
	resultsArr := make([][]int, 0)
	for valueKey, firstPartPairs := range semiResultsHash {
		if secondPartPairs, ok := semiResultsHash[s-valueKey]; ok {
			for _, firstPair := range firstPartPairs {
				for _, secondPair := range secondPartPairs {
					if semiResultsHasOnlyUniqueIndexes(firstPair, secondPair) {
						sortedResult := merge([]string{strconv.Itoa(firstPair.firstValue), strconv.Itoa(firstPair.secondValue), strconv.Itoa(secondPair.firstValue), strconv.Itoa(secondPair.secondValue)})
						stringResult := fmt.Sprintf("%d %d %d %d", sortedResult[0], sortedResult[1], sortedResult[2], sortedResult[3])
						if _, ok := resultsHash[stringResult]; !ok {
							resultsArr = append(resultsArr, sortedResult)
							resultsHash[stringResult] = true
						}

					}
				}
			}
			// fmt.Println(v, semiData)

		}
	}
	// fmt.Println("res", resultsHash)
	writer.WriteString(strconv.Itoa(len(resultsArr)))
	writer.WriteByte('\n')
	resultsArr = merge2(resultsArr)
	// fmt.Println(resultsArr)
	for _, v := range resultsArr {
		stringResult := fmt.Sprintf("%d %d %d %d", v[0], v[1], v[2], v[3])
		writer.WriteString(stringResult)
		writer.WriteByte('\n')
	}
	writer.Flush()
	// fmt.Println(n, s, data)
}

func semiResultsHasOnlyUniqueIndexes(a, b SemiResult) bool {
	if a.firstValueIndex == b.firstValueIndex || a.firstValueIndex == b.secondValueIndex {
		return false
	}
	if a.secondValueIndex == b.firstValueIndex || a.secondValueIndex == b.secondValueIndex {
		return false
	}
	return true
}

func merge(d []string) []int {
	if len(d) == 1 {
		i, _ := strconv.Atoi(d[0])
		return []int{i}
	}
	mid := len(d) >> 1
	a := merge(d[:mid])
	b := merge(d[mid:])
	result := make([]int, 0, len(a)+len(b))
	aIndex := 0
	bIndex := 0
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

func merge2(d [][]int) [][]int {
	if len(d) == 1 {
		return d
	}
	mid := len(d) >> 1
	a := merge2(d[:mid])
	b := merge2(d[mid:])
	result := make([][]int, 0, len(a)+len(b))
	aIndex := 0
	bIndex := 0
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
		if lessArraysOfFour(a[aIndex], b[bIndex]) {
			result = append(result, a[aIndex])
			aIndex++
		} else {
			result = append(result, b[bIndex])
			bIndex++
		}
	}
	return result
}

func lessArraysOfFour(a, b []int) bool {
	switch {
	case a[0] != b[0]:
		return a[0] < b[0]
	case a[1] != b[1]:
		return a[1] < b[1]
	case a[2] != b[2]:
		return a[2] < b[2]
	default:
		return a[3] <= b[3]
	}
}

func main() {
	SumOfFour(os.Stdin, os.Stdout)
}
