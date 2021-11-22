package clubs

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func Clubs(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	scaner := bufio.NewScanner(reader)
	scaner.Scan()
	n, _ := strconv.Atoi(scaner.Text())
	writer := bufio.NewWriter(w)
	hash := map[string]bool{}
	for i := 0; i < n; i++ {
		scaner.Scan()
		if !hash[scaner.Text()] {
			hash[scaner.Text()] = true
			writer.WriteString(scaner.Text())
			writer.WriteByte('\n')
		}
	}
	writer.Flush()
}

func main() {
	Clubs(os.Stdin, os.Stdout)
}

func hasher(s string, mod int) int {
	sumOfAll := 0
	for _, b := range s {
		sumOfAll += int(b)
	}
	fmt.Println(sumOfAll, s)
	return sumOfAll % mod
}
