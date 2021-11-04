package calc

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type SumStack []int

// я бы сюда эррор засунул, но по сути гаранитруется же, что не будет математических операций без данных
func (s *SumStack) Pop() int {
	item := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return item
}

func (s *SumStack) Push(i int) {
	*s = append(*s, i)
}

func (s *SumStack) MathOperation(operator string) {
	secondNum := s.Pop()
	firstNum := s.Pop()
	switch operator {
	case "-":
		s.Push(firstNum - secondNum)
	case "+":
		s.Push(firstNum + secondNum)
	case "/":
		if firstNum < 0 {
			// просто точный ответ округилил вниз
			s.Push(int(math.Floor(float64(firstNum) / float64(secondNum))))
		} else {

			s.Push(firstNum / secondNum)
		}

	case "*":
		s.Push(firstNum * secondNum)
	}
}

func Calc() {
	stdReader := bufio.NewReader(os.Stdin)
	intRegExp := regexp.MustCompile(`^[\D]$`)
	mainString, _ := stdReader.ReadString('\n')
	mainSlice := strings.Fields(mainString)
	stack := make(SumStack, 0, len(mainSlice))
	for _, v := range mainSlice {
		if intRegExp.MatchString(v) {
			stack.MathOperation(v)
		} else {
			n, _ := strconv.Atoi(v)
			stack.Push(n)
		}
	}
	fmt.Println(stack.Pop())
}
