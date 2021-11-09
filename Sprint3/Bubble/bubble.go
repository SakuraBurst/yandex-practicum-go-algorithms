package bubble

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func Bubble(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	writer := bufio.NewWriter(w)
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	stringData := strings.Fields(scanner.Text())
	data := make([]int, len(stringData))
	for i, v := range stringData {
		nData, _ := strconv.Atoi(v)
		data[i] = nData
	}
	for i := 0; i < n; i++ {
		c := 0
		isChanged := false
		for g := 0; g < (n - 1 - c); g++ {
			if data[g] > data[g+1] {
				isChanged = true
				data[g], data[g+1] = data[g+1], data[g]
			}
		}
		if !isChanged {
			if writer.Buffered() == 0 {
				for i, v := range data {
					if i == n-1 {
						writer.WriteString(strconv.Itoa(v))
						writer.WriteByte('\n')
					} else {
						writer.WriteString(strconv.Itoa(v))
						writer.WriteString(" ")
					}
				}
			}
			break
		} else {
			for i, v := range data {
				if i == n-1 {
					writer.WriteString(strconv.Itoa(v))
					writer.WriteByte('\n')
				} else {
					writer.WriteString(strconv.Itoa(v))
					writer.WriteString(" ")
				}
			}
		}
	}
	writer.Flush()
}
