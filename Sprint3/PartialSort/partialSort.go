package partialSort

import (
	"bufio"
	"io"
	"os"
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
	blockCounter := 0
	currentBlockLenght := 0
	for i := 0; i < n; i++ {
		if lookingFor == intData[i] && currentBlockLenght == 0 {
			lookingFor = i + 1
			blockCounter++
		} else {
			if intData[i] > lookingFor {
				currentBlockLenght++
			} else {
				currentBlockLenght--
			}
			if currentBlockLenght == 0 {
				blockCounter++
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

func main() {
	PartialSort(os.Stdin, os.Stdout)
}
