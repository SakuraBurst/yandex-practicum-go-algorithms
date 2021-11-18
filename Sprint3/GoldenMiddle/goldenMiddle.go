package goldenMiddle

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func GoldenMiddle(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	nStr, _ := reader.ReadString('\n')
	mStr, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nStr))
	m, _ := strconv.Atoi(strings.TrimSpace(mStr))
	stringData, _ := reader.ReadString('\n')
	mData := strings.Fields(stringData)
	stringData, _ = reader.ReadString('\n')
	nData := strings.Fields(stringData)
	stringData = ""
	totalAmountOfElements := m + n
	mergeUntil := totalAmountOfElements / 2
	isEven := (totalAmountOfElements % 2) == 0
	middleData := make([]int, 0, 2)

	mIndex := 0
	nIndex := 0

	for mIndex < len(mData) || nIndex < len(nData) {
		overallIndex := mIndex + nIndex
		if overallIndex > mergeUntil {
			break
		}
		if mIndex == len(mData) {
			if overallIndex >= mergeUntil-1 {
				middleData = append(middleData, AtoiMust(nData[nIndex]))
			}
			nIndex++

			continue
		}
		if nIndex == len(nData) {
			if overallIndex >= mergeUntil-1 {
				middleData = append(middleData, AtoiMust(mData[mIndex]))
			}
			mIndex++
			continue
		}
		if AtoiMust(mData[mIndex]) <= AtoiMust(nData[nIndex]) {
			if overallIndex >= mergeUntil-1 {
				middleData = append(middleData, AtoiMust(mData[mIndex]))
			}
			mIndex++
		} else {
			if overallIndex >= mergeUntil-1 {
				middleData = append(middleData, AtoiMust(nData[nIndex]))
			}
			nIndex++
		}

	}
	var result string
	if isEven {
		sumOfMiddleValues := middleData[0] + middleData[1]
		if sumOfMiddleValues%2 == 0 {
			result = strconv.Itoa(sumOfMiddleValues / 2)
		} else {
			result = fmt.Sprintf("%.1f", float64(sumOfMiddleValues)/2)
		}
	} else {
		result = strconv.Itoa(middleData[1])
	}
	writer := bufio.NewWriter(w)
	writer.WriteString(result)
	writer.Flush()
}

func AtoiMust(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
