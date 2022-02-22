package sameSums

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Matrix [][]int

func (m Matrix) String() string {
	s := ""
	for _, ints := range m {
		s += fmt.Sprint(ints)
		s += string('\n')
	}
	return s
}

func SameSums(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	inputData := strings.Fields(scanner.Text())
	data := make([]int, n)
	sum := 0
	sum2 := make(map[int]bool)
	for i := 0; i < n; i++ {
		data[i], _ = strconv.Atoi(inputData[i])
		sum += data[i]
		sum2[data[i]] = true
	}
	writer := bufio.NewWriter(w)
	if sum%2 != 0 {
		writer.WriteString("False")
		writer.Flush()
		return
	}
	dp := make(Matrix, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, sum/2+1)
		if i == 0 {
			continue
		}
		for g := 1; g < sum/2+1; g++ {
			dp[i][g] = dp[i-1][g]
			if g == data[i-1] {
				dp[i][g]++
			}
			if g > data[i-1] && dp[i-1][g-data[i-1]] > 0 {
				dp[i][g]++
			}

		}
	}
	//fmt.Println(dp)

	//for i, v := range dp[n] {
	//	if i == 0 {
	//		continue
	//	}
	//	log.Println(sum-v, data[i-1])
	//	if sum2[sum-v] && sum-v != data[i-1] {
	//
	//		writer.WriteString("True")
	//		writer.Flush()
	//		return
	//	}
	//}
	if dp[n][sum/2] > 1 {
		writer.WriteString("True")
	} else {
		writer.WriteString("False")
	}

	writer.Flush()

}

func main() {
	SameSums(os.Stdin, os.Stdout)
}
