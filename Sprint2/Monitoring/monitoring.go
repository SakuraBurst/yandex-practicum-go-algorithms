package monitoring

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Matrix [][]string

func Monitoring() {
	stdReader := bufio.NewReader(os.Stdin)
	stdScaner := bufio.NewScanner(stdReader)
	stdScaner.Scan()
	linesCountString := stdScaner.Text()
	stdScaner.Scan()
	lineLengthString := stdScaner.Text()
	linesCount, _ := strconv.Atoi(linesCountString)
	lineLength, _ := strconv.Atoi(lineLengthString)
	m := make(Matrix, lineLength)
	for i := 0; i < linesCount; i++ {
		stdScaner.Scan()
		for g, v := range strings.Fields(stdScaner.Text()) {
			if m[g] == nil {
				m[g] = make([]string, 0, linesCount)
			}
			m[g] = append(m[g], v)
		}
	}
	stdWriter := bufio.NewWriter(os.Stdout)
	for _, l := range m {
		stdWriter.WriteString(strings.Join(l, " "))
		stdWriter.WriteByte('\n')
	}
	stdWriter.Flush()
}
