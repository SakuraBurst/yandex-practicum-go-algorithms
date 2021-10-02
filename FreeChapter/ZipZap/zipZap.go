package zipzap

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ZipZap() {
	stdReader := bufio.NewReader(os.Stdin)
	stdWriter := bufio.NewWriter(os.Stdout)
	nLine, _ := stdReader.ReadString('\n')
	n, _ := strconv.ParseInt(strings.TrimSpace(nLine), 10, 64)
	firstArrayLine, _ := stdReader.ReadString('\n')
	firstSplited := strings.Split(strings.TrimSpace(firstArrayLine), " ")
	secondArrayLine, _ := stdReader.ReadString('\n')
	secondSplited := strings.Split(strings.TrimSpace(secondArrayLine), " ")
	result := make([]int64, n*2)
	var i int64
	for i = 0; i < n; i++ {
		writeIndex := i * 2
		firstItem, _ := strconv.ParseInt(strings.TrimSpace(firstSplited[i]), 10, 64)
		secondItem, _ := strconv.ParseInt(strings.TrimSpace(secondSplited[i]), 10, 64)
		result[writeIndex], result[writeIndex+1] = firstItem, secondItem
	}

	for _, elem := range result {
		fmt.Fprintf(stdWriter, "%d", elem)
	}
	stdWriter.Flush()
}
