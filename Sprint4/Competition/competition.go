package competition

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func Competition(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	nString, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nString))
	dataString, _ := reader.ReadString('\n')
	data := strings.Fields(dataString)
	max := 0
	for i := 0; i < n; i++ {
		count := 0
		currentMax := 0
		for g, c := i, 0; g < n; g, c = g+1, c+1 {
			if n-count < 0 {
				break
			}
			if data[g] == "0" {
				count--
			} else {
				count++
			}
			if count == 0 {
				currentMax = c + 1
			}
		}
		if currentMax > max {
			max = currentMax
		}
	}
	writer := bufio.NewWriter(w)

	writer.WriteString(strconv.Itoa(max))

	writer.Flush()
}

func main() {
	Competition(os.Stdin, os.Stdout)
}
