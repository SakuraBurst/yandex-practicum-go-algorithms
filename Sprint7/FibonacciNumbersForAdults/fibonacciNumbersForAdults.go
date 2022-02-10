package fibonacciNumbersForAdults

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

const Mod = 1000000007

func FibonacciNumbersForAdults(r io.Reader, w io.Writer) {

	reader := bufio.NewReader(r)
	nString, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nString))
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	writer := bufio.NewWriter(w)
	for i := 2; i < n+1; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % Mod
	}

	writer.WriteString(strconv.Itoa(dp[n]))
	writer.Flush()
}

func main() {
	FibonacciNumbersForAdults(os.Stdin, os.Stdout)
}
