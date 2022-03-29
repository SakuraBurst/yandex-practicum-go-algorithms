package repeat

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func Repeat(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	s, _ := reader.ReadString('\n')
	s = s[:len(s)-1]
	prefix := string(s[0])
	prefixLength := len(prefix)
	i := 1

	for i+prefixLength <= len(s) {
		if prefix == s[i:i+prefixLength] {
			i += prefixLength
		} else {
			i++
			prefix = s[:i]
			prefixLength = len(prefix)
		}
	}
	if len(prefix) > len(s)/2 {
		prefix = s
		prefixLength = len(prefix)
	}
	io.WriteString(w, strconv.Itoa(len(s)/prefixLength))

}

func main() {
	Repeat(os.Stdin, os.Stdout)
}
