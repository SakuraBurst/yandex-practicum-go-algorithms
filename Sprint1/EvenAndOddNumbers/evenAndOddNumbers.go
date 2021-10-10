package evenAndOddNubers

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func EvenAndOddNubers() {
	stdReader := bufio.NewReader(os.Stdin)
	stdWriter := bufio.NewWriter(os.Stdout)
	allValuesString, _ := stdReader.ReadString('\n')
	allValues := strings.Split(allValuesString, " ")
	firstInt, _ := strconv.ParseInt(allValues[0], 10, 64)
	isSearchForEven := isEven(firstInt)
	for _, item := range allValues[1:] {
		currentInt, _ := strconv.ParseInt(strings.TrimSpace(item), 10, 64)
		if isSearchForEven != isEven(currentInt) {
			stdWriter.WriteString("FAIL")
			stdWriter.Flush()
			return
		}

	}
	stdWriter.WriteString("WIN")
	stdWriter.Flush()
}

func isEven(i int64) bool {
	return i%2 == 0
}
