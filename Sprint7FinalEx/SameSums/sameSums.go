package sameSums

import (
	"bufio"
	"fmt"
	"io"
	"log"
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
	fmt.Println(sum)
	dp := make(Matrix, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, n+1)
		if i == 0 {
			//for index := 1; index < n+1; index++ {
			//	dp[i][index] = data[index-1]
			//
			//}
			continue
		}
		for g := 1; g < n+1; g++ {
			//if g == 0 {
			//	dp[i][g] = data[i-1]
			//	continue
			//}
			if i == g {
				dp[i][g] = dp[i-1][g]
				continue
			}
			dp[i][g] = data[i-1] + dp[i-1][g]
			log.Println(dp[i][g] != 1 && dp[i][g] != 0 && sum%dp[i][g] == 0)
		}
	}
	fmt.Println(dp)

	for i, v := range dp[n] {
		if i == 0 {
			continue
		}
		log.Println(sum-v, data[i-1])
		if sum2[sum-v] && sum-v != data[i-1] {

			writer.WriteString("True")
			writer.Flush()
			return
		}
	}
	writer.WriteString("False")
	writer.Flush()

}

func main() {
	SameSums(os.Stdin, os.Stdout)
}
