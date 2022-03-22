package globalReplacement

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func GlobalReplacement(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	s, _ := reader.ReadString('\n')
	s = s[:len(s)-1]
	pattern, _ := reader.ReadString('\n')
	pattern = pattern[:len(pattern)-1]
	replacer, _ := reader.ReadString('\n')
	replacer = replacer[:len(replacer)-1]
	sentinel := pattern + "#" + s
	patterLen := len(pattern)
	prefixLen := patterLen + 1
	dp := make([]int, len(sentinel))
	result := make([]bool, len(s))
	for i := prefixLen; i < len(sentinel); i++ {
		k := dp[i-1]
		for sentinel[i] != sentinel[k] && k > 0 {
			k = dp[k-1]
		}
		if sentinel[k] != sentinel[i] {
			dp[i] = k
		} else {
			dp[i] = k + 1
			if dp[i] == patterLen {
				result[i-patterLen*2] = true
			}
		}

	}
	dp = nil
	builder := strings.Builder{}
	for i := 0; i < len(s); i++ {
		if result[i] {
			builder.WriteString(replacer)
			i += patterLen - 1
		} else {
			builder.WriteByte(s[i])

		}
	}
	io.WriteString(w, builder.String())
}

func main() {
	GlobalReplacement(os.Stdin, os.Stdout)
}
