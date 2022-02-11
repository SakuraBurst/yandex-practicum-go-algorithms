package atm

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

func Atm(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	atmData := strings.Fields(scanner.Text())
	atm := make([]int, 0, n)
	for _, bankNoteString := range atmData {
		bankNote, _ := strconv.Atoi(bankNoteString)
		atm = append(atm, bankNote)
	}
	matrix := make(Matrix, n+1)
	for i := 0; i < n+1; i++ {
		matrix[i] = make([]int, m+1)
		if i == 0 {
			continue
		}
		if atm[i-1] < m+1 {
			matrix[i][atm[i-1]] = 1
		}
		for g := 0; g < m+1; g++ {

			matrix[i][g] += matrix[i-1][g]

			if g-atm[i-1] >= 0 && matrix[i][g-atm[i-1]] > 0 {
				matrix[i][g] += matrix[i][g-atm[i-1]]
			}
		}
	}
	writer := bufio.NewWriter(w)
	writer.WriteString(strconv.Itoa(matrix[n][m]))
	writer.Flush()
}

func main() {
	Atm(os.Stdin, os.Stdout)
}
