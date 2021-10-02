package sum2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Sum2Naive() {
	stdReader := bufio.NewReader(os.Stdin)
	stdWriter := bufio.NewWriter(os.Stdout)
	nLine, _ := stdReader.ReadString('\n')
	n, _ := strconv.ParseInt(strings.TrimSpace(nLine), 10, 64)
	arrayLine, _ := stdReader.ReadString('\n')
	splited := strings.Split(strings.TrimSpace(arrayLine), " ")
	sLine, _ := stdReader.ReadString('\n')
	s, _ := strconv.ParseInt(strings.TrimSpace(sLine), 10, 64)
	var i int64
Loop:
	for i = 0; i < n; i++ {
		firstItem, _ := strconv.ParseInt(strings.TrimSpace(splited[i]), 10, 64)
		for g := i + 1; g < n; g++ {
			secondItem, _ := strconv.ParseInt(strings.TrimSpace(splited[g]), 10, 64)
			if firstItem+secondItem == s {
				fmt.Fprintf(stdWriter, "%d %d", firstItem, secondItem)
				break Loop
			}
		}
	}
	if stdWriter.Buffered() == 0 {
		fmt.Fprint(stdWriter, "None")
	}
	stdWriter.Flush()
}
