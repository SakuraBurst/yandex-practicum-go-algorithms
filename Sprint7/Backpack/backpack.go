package backpack

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type BackPackItem struct {
	value  int
	weight int
}

type BackPack []BackPackItem

type Matrix [][]int

func (m Matrix) String() string {
	s := ""
	for _, ints := range m {
		s += fmt.Sprint(ints)
		s += string('\n')
	}
	return s
}

func Backpack(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	initData := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(initData[0])
	M, _ := strconv.Atoi(initData[1])
	backPack := make(BackPack, 0, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		backPackItemData := strings.Fields(scanner.Text())
		weight, _ := strconv.Atoi(backPackItemData[0])
		value, _ := strconv.Atoi(backPackItemData[1])
		backPack = append(backPack, BackPackItem{
			value:  value,
			weight: weight,
		})
	}
	dp := make(Matrix, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, M+1)
		if i == 0 {
			continue
		}
		for g := 1; g < M+1; g++ {
			if g < backPack[i-1].weight {
				dp[i][g] = max(dp[i-1][g], dp[i][g-1])
			}
			if g >= backPack[i-1].weight {
				dp[i][g] = max(dp[i-1][g-backPack[i-1].weight]+backPack[i-1].value, dp[i-1][g])
			}
		}
	}
	currentPick := dp[n][M]
	i := n
	g := M
	result := make([]string, 0, 16)
	for currentPick != 0 {
		if dp[i-1][g] == currentPick {
			i--
			continue
		}
		if dp[i][g-1] == currentPick {
			g--
			continue
		}
		result = append(result, strconv.Itoa(i))
		g -= backPack[i-1].weight
		i--
		currentPick = dp[i][g]
	}
	scanner.Scan()
	writer := bufio.NewWriter(w)
	writer.WriteString(strconv.Itoa(len(result)))
	writer.WriteByte('\n')
	writer.WriteString(strings.Join(result, " "))
	writer.Flush()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	Backpack(os.Stdin, os.Stdout)
}
