package moving_average

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MovingAverage() {
	stdReader := bufio.NewReader(os.Stdin)
	stdWriter := bufio.NewWriter(os.Stdout)
	nLine, _ := stdReader.ReadString('\n')
	n, _ := strconv.ParseInt(strings.TrimSpace(nLine), 10, 64)
	arrayLine, _ := stdReader.ReadString('\n')
	splited := strings.Split(strings.TrimSpace(arrayLine), " ")
	kLine, _ := stdReader.ReadString('\n')
	k, _ := strconv.ParseFloat(strings.TrimSpace(kLine), 64)
	result := make([]float64, n-(int64(k)-1))
	var accum float64

	for i, val := range splited {
		item, _ := strconv.ParseFloat(strings.TrimSpace(val), 64)
		accum += item
		startIndex := i - (int(k) - 1)
		if i >= (int(k) - 1) {
			result[startIndex] = accum / k
			prevVal, _ := strconv.ParseFloat(strings.TrimSpace(splited[startIndex]), 64)
			accum -= prevVal
		}

	}

	for _, elem := range result {
		fmt.Fprintf(stdWriter, "%f ", elem)
	}
	stdWriter.Flush()
}
