package fieldWithFlowers

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

func FieldWithFlowers(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	matrixSizeData := strings.Fields(scanner.Text())
	matrixHeight, _ := strconv.Atoi(matrixSizeData[0])
	matrixWidth, _ := strconv.Atoi(matrixSizeData[1])
	dataMatrix := make(Matrix, matrixHeight+1)
	dp := make(Matrix, matrixHeight+2)
	for i := 0; i < matrixHeight+1; i++ {
		dataMatrix[i] = make([]int, matrixWidth+1)
		dp[i] = make([]int, matrixWidth+1)
	}
	dp[matrixHeight+1] = make([]int, matrixWidth+1)
	for i := 1; i < matrixHeight+1; i++ {
		scanner.Scan()
		rowData := strings.Split(strings.TrimSpace(scanner.Text()), "")
		for g, datum := range rowData {
			if datum == "0" {
				continue
			}
			numberOfFlowers, _ := strconv.Atoi(datum)
			dataMatrix[i][g+1] = numberOfFlowers
		}
	}
	for i := matrixHeight; i > 0; i-- {
		for g := 1; g < matrixWidth+1; g++ {
			dp[i][g] = max(dp[i+1][g], dp[i][g-1]) + dataMatrix[i][g]
		}
	}
	writer := bufio.NewWriter(w)
	writer.WriteString(strconv.Itoa(dp[1][matrixWidth]))
	writer.Flush()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	FieldWithFlowers(os.Stdin, os.Stdout)
}
