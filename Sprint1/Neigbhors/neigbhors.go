package neigbhors

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Marix [][]string

func (m Marix) returnNeigbhors(x, y int) []int {
	r := make([]int, 0, 4*2)
	switch {
	case x == 0 && x == (len(m)-1):
		break

	case x < (len(m) - 1):
		if x == 0 {
			r = append(r, parseInt(m[x+1][y]))
		} else {
			r = append(r, parseInt(m[x+1][y]), parseInt(m[x-1][y]))
		}
	case x == (len(m) - 1):
		r = append(r, parseInt(m[x-1][y]))
	}
	switch {
	case y == 0 && y == (len(m[x])-1):
		break

	case y < (len(m[x]) - 1):
		if y == 0 {
			r = append(r, parseInt(m[x][y+1]))
		} else {
			r = append(r, parseInt(m[x][y+1]), parseInt(m[x][y-1]))
		}
	case y == (len(m[x]) - 1):
		r = append(r, parseInt(m[x][y-1]))
	}
	return r

}

func parseInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func Neigbhors() {
	stdReader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(stdReader)
	scanner.Split(bufio.ScanLines)
	howMuchLines, _ := matrixInfoFromScanner(scanner)
	matrix := make(Marix, howMuchLines)
	for i := 0; i < howMuchLines; i++ {
		scanner.Scan()
		matrix[i] = strings.Split(strings.TrimSpace(scanner.Text()), " ")
	}
	x, y := coordinates(scanner)

	stdWriter := bufio.NewWriter(os.Stdout)
	sorted := simpleSort(matrix.returnNeigbhors(x, y))
	for _, v := range sorted {
		fmt.Fprintf(stdWriter, "%d ", v)
	}
	stdWriter.Flush()
}

func simpleSort(m []int) []int {
	for i := 1; i < len(m); i++ {
		for g := i - 1; g >= 0 && m[g+1] < m[g]; g-- {
			m[g+1], m[g] = m[g], m[g+1]
		}
	}
	return m
}

func matrixInfoFromScanner(scanner *bufio.Scanner) (howMuchLines, linesLength int) {
	howMuchLines, linesLength = returnFirstTwoIntFromScanner(scanner)
	return
}

func coordinates(scanner *bufio.Scanner) (x, y int) {
	x, y = returnFirstTwoIntFromScanner(scanner)
	return
}

func returnFirstTwoIntFromScanner(scanner *bufio.Scanner) (x, y int) {
	scanner.Scan()
	x, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	y, _ = strconv.Atoi(scanner.Text())
	return
}
