package countingThePrefixFunction

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func CountingThePrefixFunction(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	s, _ := reader.ReadString('\n')
	s = s[:len(s)-1]
	dp := make([]int, len(s))
	writer := bufio.NewWriter(w)
	writer.WriteString("0 ")
	for i := 1; i < len(s); i++ {
		k := dp[i-1]
		for k > 0 && s[i] != s[k] {
			k = dp[k-1]
		}
		if s[i] == s[k] {
			dp[i] = k + 1
			writer.WriteString(strconv.Itoa(dp[i]) + " ")
		} else {
			dp[i] = 0
			writer.WriteString("0 ")
		}
	}
	writer.Flush()
}

func main() {
	CountingThePrefixFunction(os.Stdin, os.Stdout)
}
