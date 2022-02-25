/**
-- ПРИНЦИП РАБОТЫ --
1. Что будет храниться в матрице dp?
	В матрице dp будет храниться то, сколькими способами можно набрать сумму g используя число data[i]
2. Каким будет базовый случай для задачи?
	Если data[i] == g, то можно набрать сумму g как минимум одним способом, это и будет базовый случай.
3. Каким будет переход динамики?
	Ищем g - data[i]
4. Каким будет порядок вычисления данных в массиве dp?
	Используя Динамическое программирование назад, смотря предыдущие значения
5. Где будет располагаться ответ на исходный вопрос?
	В матрице dp[len(data)][sum]
-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
*/

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
	for i := 0; i < n; i++ {
		data[i], _ = strconv.Atoi(inputData[i])
		sum += data[i]
	}
	writer := bufio.NewWriter(w)
	if sum%2 != 0 {
		writer.WriteString("False")
		writer.Flush()
		return
	}
	dpPrev := make([]int, sum/2+1)
	dp := make([]int, sum/2+1)
	for i := 1; i < n+1; i++ {
		for g := 1; g < sum/2+1; g++ {
			dp[g] = dpPrev[g]
			if g == data[i-1] {
				dp[g]++
			}
			if g > data[i-1] && dpPrev[g-data[i-1]] > 0 {
				dp[g]++
			}
		}
		dpPrev = dp
		dp = make([]int, sum/2+1)
	}
	//dp := make(Matrix, n+1)
	//for i := 0; i < n+1; i++ {
	//	dp[i] = make([]int, sum/2+1)
	//	if i == 0 {
	//		continue
	//	}
	//	for g := 1; g < sum/2+1; g++ {
	//		dp[i][g] = dp[i-1][g]
	//		if g == data[i-1] {
	//			dp[i][g]++
	//		}
	//		if g > data[i-1] && dp[i-1][g-data[i-1]] > 0 {
	//			dp[i][g]++
	//		}
	//
	//	}
	//}
	if dpPrev[sum/2] > 1 {
		writer.WriteString("True")
	} else {
		writer.WriteString("False")
	}

	writer.Flush()

}

func main() {
	SameSums(os.Stdin, os.Stdout)
}
