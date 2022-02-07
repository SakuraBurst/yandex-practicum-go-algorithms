package horoscopes

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

func Horoscopes(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	dataString := strings.Fields(scanner.Text())
	nData := fromStringDataToData(n, dataString)
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	dataString = strings.Fields(scanner.Text())
	mData := fromStringDataToData(m, dataString)
	dp := make(Matrix, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
		if i == 0 {
			continue
		}
		lookingFor := nData[i-1]
		for g := 1; g < m+1; g++ {
			currentSequenceItem := mData[g-1]
			if lookingFor != currentSequenceItem {
				dp[i][g] = max(dp[i-1][g], dp[i][g-1])
			} else {
				dp[i][g] = dp[i-1][g-1] + 1
			}
		}
	}
	writer := bufio.NewWriter(w)
	if dp[n][m] == 0 {
		writer.WriteString("0")
		writer.Flush()
		return
	}
	var nRes []string = nil
	var mRes []string = nil
	i := n
	g := m
	for i != 0 || g != 0 {
		if dp[i][g] == 0 {
			break
		}
		if dp[i][g] == dp[i-1][g] {
			i--
			continue
		}
		if dp[i][g] == dp[i][g-1] {
			g--
			continue
		}
		nRes = append(nRes, strconv.Itoa(i))
		mRes = append(mRes, strconv.Itoa(g))
		i--
		g--

	}
	writer.WriteString(strconv.Itoa(dp[n][m]))
	writer.WriteByte('\n')
	writer.WriteString(reverseAndMakeString(nRes))
	writer.WriteByte('\n')
	writer.WriteString(reverseAndMakeString(mRes))
	writer.Flush()
}

func fromStringDataToData(n int, dataString []string) []int {
	nData := make([]int, n)
	for i := 0; i < n; i++ {
		nSequenceItem, _ := strconv.Atoi(dataString[i])
		nData[i] = nSequenceItem
	}
	return nData
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func reverseAndMakeString(s []string) string {
	for i, g := 0, len(s)-1; i < g; i, g = i+1, g-1 {
		s[i], s[g] = s[g], s[i]
	}
	return strings.Join(s, " ")
}

func main() {
	Horoscopes(os.Stdin, os.Stdout)
}
