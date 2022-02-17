package levenshteinDistance

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Matrix [][]int

func (m Matrix) String(si, mi string) string {
	s := ""
	s += "     " + fmt.Sprint(strings.Join(strings.Split(mi, ""), " "))
	s += string('\n')
	for i, ints := range m {
		if i > 0 {
			s += string(si[i-1]) + " "
		} else {
			s += "  "
		}
		s += fmt.Sprint(ints)
		s += string('\n')
	}
	return s
}

func LevenshteinDistance(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	s := scanner.Text()
	scanner.Scan()
	m := scanner.Text()
	//var maxStringLength int
	//var minStringLength int
	if len(s) > len(m) {
		//maxStringLength = len(s)
		//minStringLength = len(m)
	} else {
		//maxStringLength = len(m)
		s, m = m, s
		//minStringLength = len(s)
	}
	sChars := strings.Split(s, "")
	mChars := strings.Split(m, "")
	matrix := make(Matrix, len(sChars)+1)
	for i := 0; i < len(sChars)+1; i++ {
		matrix[i] = make([]int, len(mChars)+1)
		if i == 0 {
			continue
		}
		for g := 1; g < len(mChars)+1; g++ {
			if sChars[i-1] == mChars[g-1] {
				if matrix[i][g-1] == 0 {
					matrix[i][g] = 0
				} else {
					matrix[i][g] = matrix[i][g-1] - 1
				}

			} else {
				matrix[i][g] = matrix[i][g-1] + 1
			}
		}
	}
	fmt.Println(matrix.String(s, m))
	//fmt.Println(maxStringLength, minStringLength, matrix[len(sChars)][len(mChars)])
	writer := bufio.NewWriter(w)
	writer.WriteString(strconv.Itoa(matrix[len(sChars)][len(mChars)]))
	writer.Flush()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	LevenshteinDistance(os.Stdin, os.Stdout)
}
