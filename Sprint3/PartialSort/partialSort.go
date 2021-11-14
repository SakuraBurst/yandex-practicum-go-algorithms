package partialSort

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func PartialSort(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	scaner := bufio.NewScanner(reader)
	scaner.Scan()
	n, _ := strconv.Atoi(scaner.Text())
	scaner.Scan()
	stringsData := strings.Fields(scaner.Text())
	intData := make([]int, len(stringsData))
	for i, v := range stringsData {
		iV, _ := strconv.Atoi(v)
		intData[i] = iV
	}
	lookingFor := 0
	currentMax := 0
	blockCounter := 0
	betwinLookingForAndCurentMax := 0
	for i := 0; i < n; i++ {
		if lookingFor == intData[i] && currentMax == 0 {
			lookingFor = i + 1
			blockCounter++
		} else {
			if currentMax < intData[i] {
				currentMax = intData[i]
				betwinLookingForAndCurentMax++
			}
			if intData[i] < currentMax {
				betwinLookingForAndCurentMax++
			}
			if betwinLookingForAndCurentMax == currentMax-lookingFor+1 {
				blockCounter++
				currentMax = 0
				lookingFor = i + 1
			}
		}

	}
	if blockCounter == 0 {
		blockCounter = 1
	}

	writer := bufio.NewWriter(w)
	writer.WriteString(strconv.Itoa(blockCounter))
	writer.Flush()
}
