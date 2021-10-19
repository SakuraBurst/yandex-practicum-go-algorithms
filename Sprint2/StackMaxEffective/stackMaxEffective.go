package stackmaxeffective

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type StackMax struct {
	stack      []int
	max        []int
	currentMax int
}

func (s *StackMax) pop() error {
	if s.isEmpty() {
		return errors.New("error")
	}
	deletedInt := s.stack[len(s.stack)-1]
	if deletedInt == s.currentMax {
		s.max = s.max[:len(s.max)-1]
		if len(s.max) > 0 {
			s.currentMax = s.max[len(s.max)-1]
		}
	}
	s.stack = s.stack[:len(s.stack)-1]
	return nil
}
func (s *StackMax) push(i int) {
	if s.isEmpty() {
		s.currentMax = i
		s.max = append(s.max, i)
	} else {
		if s.currentMax <= i {
			s.currentMax = i
			s.max = append(s.max, i)
		}
	}
	s.stack = append(s.stack, i)

}
func (s StackMax) getMax() (int, error) {
	if s.isEmpty() {
		return 0, errors.New("None")
	}
	return s.currentMax, nil
}

func (s StackMax) isEmpty() bool {
	return len(s.stack) == 0
}

var GET_MAX = regexp.MustCompile(`(?i)get_max`)
var POP = regexp.MustCompile(`(?i)pop`)
var PUSH = regexp.MustCompile(`(?i)^push`)
var INT = regexp.MustCompile(`-?\d+`)

func StackMaxEffective() {
	var n int
	fmt.Scan(&n)
	stdReader := bufio.NewReader(os.Stdin)
	stdReader.ReadByte()
	stdScanner := bufio.NewScanner(stdReader)
	stack := StackMax{make([]int, 0, n), make([]int, 0, n), 0}
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
