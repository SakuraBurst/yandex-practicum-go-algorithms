package stackmax

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Stack []int

func (s *Stack) pop() error {
	if s.isEmpty() {
		return errors.New("error")
	}
	*s = (*s)[:len(*s)-1]
	return nil
}
func (s *Stack) push(i int) {
	*s = append(*s, i)

}
func (s Stack) getMax() (int, error) {
	if s.isEmpty() {
		return 0, errors.New("None")
	}
	max := s[0]
	if len(s) == 1 {
		return max, nil
	} else {
		for i := 1; i < len(s); i++ {
			if max < s[i] {
				max = s[i]
			}
		}
	}
	return max, nil
}

func (s Stack) isEmpty() bool {
	return len(s) == 0
}

var GET_MAX = regexp.MustCompile(`(?i)get_max`)
var POP = regexp.MustCompile(`(?i)pop`)
var PUSH = regexp.MustCompile(`(?i)^push`)
var INT = regexp.MustCompile(`-?\d+`)

func StackMax() {
	var n int
	fmt.Scan(&n)
	stdReader := bufio.NewReader(os.Stdin)
	stdReader.ReadByte()
	stdScanner := bufio.NewScanner(stdReader)
	stack := make(Stack, 0)
	stdWriter := bufio.NewWriter(os.Stdout)
	for i := 0; i < n; i++ {
		stdScanner.Scan()
		switch {
		case GET_MAX.MatchString(stdScanner.Text()):
			i, err := stack.getMax()
			if err != nil {
				stdWriter.WriteString(err.Error())
				stdWriter.WriteByte('\n')
			} else {
				stdWriter.WriteString(strconv.Itoa(i))
				stdWriter.WriteByte('\n')
			}
		case POP.MatchString(stdScanner.Text()):
			err := stack.pop()
			if err != nil {
				stdWriter.WriteString(err.Error())
				stdWriter.WriteByte('\n')
			}
		case PUSH.MatchString(stdScanner.Text()):
			mm := INT.FindString(stdScanner.Text())
			i, _ := strconv.Atoi(mm)
			stack.push(i)
		}
	}
	stdWriter.Flush()
}
