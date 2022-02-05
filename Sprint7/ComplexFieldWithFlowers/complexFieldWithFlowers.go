package complexFieldWithFlowers

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

func ComplexFieldWithFlowers(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	matrixSizeData := strings.Fields(scanner.Text())
	matrixHeight, _ := strconv.Atoi(matrixSizeData[0])
	matrixWidth, _ := strconv.Atoi(matrixSizeData[1])
	matrix := make(Matrix, matrixHeight+1)
	dp := make(Matrix, matrixHeight+2)
	for i := 1; i < matrixHeight+1; i++ {
		scanner.Scan()
		matrix[i] = make([]int, matrixWidth+1)
		dp[i] = make([]int, matrixWidth+1)
		matrixRowData := strings.Split(strings.TrimSpace(scanner.Text()), "")
		for g, rowItemString := range matrixRowData {
			rowItem, _ := strconv.Atoi(rowItemString)
			matrix[i][g+1] = rowItem
		}
	}
	dp[matrixHeight+1] = make([]int, matrixWidth+1)
	for i := matrixHeight; i > 0; i-- {
		for g := 1; g < matrixWidth+1; g++ {
			dp[i][g] = max(dp[i+1][g], dp[i][g-1]) + matrix[i][g]
		}
	}
	height := 1
	width := matrixWidth
	sBuilder := strings.Builder{}
	for height != matrixHeight || width != 1 {
		if height == matrixHeight {
			width--
			sBuilder.WriteString("R")
			continue
		}
		if width == 1 {
			height++
			sBuilder.WriteString("U")
			continue
		}
		if dp[height+1][width] > dp[height][width-1] {
			height++
			sBuilder.WriteString("U")
		} else {
			width--
			sBuilder.WriteString("R")
		}
	}
	writer := bufio.NewWriter(w)
	writer.WriteString(strconv.Itoa(dp[1][matrixWidth]))
	writer.WriteByte('\n')
	writer.WriteString(reverse([]byte(sBuilder.String())))
	writer.Flush()
}

func reverse(s []byte) string {
	for i, g := 0, len(s)-1; i < g; i, g = i+1, g-1 {
		s[i], s[g] = s[g], s[i]
	}
	return string(s)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	ComplexFieldWithFlowers(os.Stdin, os.Stdout)
}
