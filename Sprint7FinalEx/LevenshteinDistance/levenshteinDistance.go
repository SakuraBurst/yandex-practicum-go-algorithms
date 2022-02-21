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
	sChars := strings.Split(s, "")
	mChars := strings.Split(m, "")
	matrix := make(Matrix, len(sChars)+1)
	for i := 0; i < len(sChars)+1; i++ {
		matrix[i] = make([]int, len(mChars)+1)
		if i == 0 {
			for g := 0; g < len(mChars)+1; g++ {
				matrix[i][g] = g
			}
			continue
		}

		for g := 0; g < len(mChars)+1; g++ {
			if g == 0 {
				matrix[i][g] = i
				continue
			}
			matrix[i][g] = min(matrix[i-1][g]+1, matrix[i][g-1]+1, matrix[i-1][g-1]+diff(sChars[i-1], mChars[g-1]))
		}
	}
	//fmt.Println(matrix.String(s, m))
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

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

func diff(a, b string) int {
	if a == b {
		return 0
	}
	return 1
}

func main() {
	LevenshteinDistance(os.Stdin, os.Stdout)
}

//subStirngModifedIndexes := make(map[int]bool, len(mChars))
//for i := 0; i < len(sChars)+1; i++ {
//matrix[i] = make([]int, len(mChars)+1)
//if i == 0 {
//for g := 0; g < len(mChars)+1; g++ {
//matrix[i][g] = g
//}
//continue
//}
//
//isModifed := false
//for g := 0; g < len(mChars)+1; g++ {
//if g == 0 {
//matrix[i][g] = i
//continue
//}
//if sChars[i-1] == mChars[g-1] && !isModifed && !subStirngModifedIndexes[g] {
//matrix[i][g] = max(matrix[i][g-1], matrix[i-1][g]) - 1
//isModifed = true
//subStirngModifedIndexes[g] = true
//} else {
//if isModifed || subStirngModifedIndexes[g] {
//matrix[i][g] = min(matrix[i][g-1]+1, matrix[i-1][g]+1)
//} else {
//matrix[i][g] = max(matrix[i][g-1], matrix[i-1][g])
//}
//
//}
//}
//}
