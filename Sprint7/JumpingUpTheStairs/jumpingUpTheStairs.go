package jumpingUpTheStairs

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

const Mod = 1000000007

func JumpingUpTheStairs(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	initDataString, _ := reader.ReadString('\n')
	initData := strings.Fields(initDataString)
	N, _ := strconv.Atoi(initData[0])
	K, _ := strconv.Atoi(initData[1])
	dp := make([]int, N+1)
	dp[1] = 0
	for i := 2; i < K+2 && i < N+1; i++ {
		dp[i] = 1
	}
	for i := 2; i < N+1; i++ {
		nextIndex := i + 1
		for g := nextIndex; g < nextIndex+K && g < N+1; g++ {

			dp[g] = (dp[g] + dp[i]) % Mod
		}
	}

	writer := bufio.NewWriter(w)
	writer.WriteString(strconv.Itoa(dp[N]))
	writer.Flush()
}

func main() {
	JumpingUpTheStairs(os.Stdin, os.Stdout)
}
