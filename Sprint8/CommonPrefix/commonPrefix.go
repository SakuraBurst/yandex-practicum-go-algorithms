package commonPrefix

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func CommonPrefix(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	nS, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(nS[:len(nS)-1])
	lastPrefix := ""
	for i := 0; i < n; i++ {
		str, _ := reader.ReadString('\n')
		str = str[:len(str)-1]
		if i == 0 {
			lastPrefix = str
			continue
		}

		i := 0
		g := 0
		for i < len(lastPrefix) && g < len(str) && lastPrefix[i] == str[g] {
			i++
			g++
		}
		lastPrefix = str[:g]
	}
	io.WriteString(w, strconv.Itoa(len(lastPrefix)))
}

func main() {
	CommonPrefix(os.Stdin, os.Stdout)
}
