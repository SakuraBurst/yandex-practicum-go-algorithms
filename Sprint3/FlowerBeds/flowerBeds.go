package flowerBeds

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func FlowerBeds(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	scaner := bufio.NewScanner(reader)
	scaner.Scan()
	data := [][2]int{}
	for scaner.Scan() {
		v := [2]int{}
		for i, s := range strings.Fields(scaner.Text()) {
			sInt, _ := strconv.Atoi(s)
			v[i] = sInt
		}
		data = append(data, v)
	}
	data = mergeSort(data)
	result := [][2]int{data[0]}
	for i := 1; i < len(data); i++ {
		lastResultIndex := len(result) - 1
		if data[i][0] >= result[lastResultIndex][0] && data[i][1] <= result[lastResultIndex][1] {
			continue
		}
		if data[i][0] <= result[lastResultIndex][1] && data[i][1] >= result[lastResultIndex][1] {
			newResult := [2]int{}
			if data[i][0] <= result[lastResultIndex][0] {
				newResult[0] = data[i][0]
			} else {
				newResult[0] = result[lastResultIndex][0]
			}
			newResult[1] = data[i][1]
			result[lastResultIndex] = newResult
		} else {
			result = append(result, data[i])
		}
	}
	writer := bufio.NewWriter(w)
	for _, v := range result {
		writer.WriteString(strconv.Itoa(v[0]) + " " + strconv.Itoa(v[1]))
		writer.WriteByte('\n')
	}
	writer.Flush()
}

func mergeSort(bracket [][2]int) [][2]int {
	bracketLenght := len(bracket)
	if bracketLenght == 1 {
		return bracket
	} else {
		half := bracketLenght / 2
		first := mergeSort(bracket[half:])
		second := mergeSort(bracket[:half])
		firstIndex := 0
		secondIndex := 0
		newVal := [][2]int{}
		for i := 0; i < bracketLenght; i++ {
			if firstIndex == len(first) {
				newVal = append(newVal, second[secondIndex])
				secondIndex++
				continue
			}
			if secondIndex == len(second) {
				newVal = append(newVal, first[firstIndex])
				firstIndex++
				continue
			}
			if first[firstIndex][0] <= second[secondIndex][0] {
				newVal = append(newVal, first[firstIndex])
				firstIndex++
			} else {
				newVal = append(newVal, second[secondIndex])
				secondIndex++
			}
		}
		return newVal
	}
}
