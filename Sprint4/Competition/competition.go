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
	writer := bufio.NewWriter(w)
	firstLoopResult := 0
	m := map[int][]int{}
	if n == 0 {
		writer.WriteString("0")
	}
	for i := 0; i < n; i++ {
		if data[i] == "0" {
			firstLoopResult++
			m[firstLoopResult] = append(m[firstLoopResult], i+1)
		} else {
			firstLoopResult--
			m[firstLoopResult] = append(m[firstLoopResult], i+1)
		}
	}
	max := 0
	// fmt.Println(m)
	for k, v := range m {
		mr := 0
		if k == 0 {
			mr = v[len(v)-1]
		} else {
			mr = v[len(v)-1] - v[0]
		}

		if mr > max {
			max = mr
		}
	}
	// fmt.Println(max)

	writer.WriteString(strconv.Itoa(max))

	writer.Flush()
}

func main() {
	Competition(os.Stdin, os.Stdout)
}
