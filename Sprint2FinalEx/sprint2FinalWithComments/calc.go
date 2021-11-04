/*
-- ПРИНЦИП РАБОТЫ --
id в контесте 55657416

Привет) Ну собственно сделал как и было написано в задаче :D Когда встречаю число - помещаю его в стек, когда встречаю математический оператор -
беру первое число из стека как второй элемент выражения, второе число из стекак как первый элемент выражения, и произвожу операцию. Для деления с отрицательным делимыым
перевожу в флоат, и округляю в меньшую сторону



-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --

Так как, насколько я понял, при парсинге польской нотации нужно при получении знака операции, применить операцию к двум последним числам, стек с этим прекрасно справляется,
вытаскивая эти числа, и помещаая результат вместо них



-- ВРЕМЕННАЯ СЛОЖНОСТЬ --

Сама операция высчитывания константная и требует тольк снятия последних двух элементов стека
O(1)

высчитывание всего выражения требует, я думаю, О(n), где n == количество математических операций

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --

Так как числа и знаки могут идти в произвольном порядке, и может быть так, что сначала идут числа, а только потом знаки,
то в худшем случае пространственная слжоность равна O(n) где n количество чисел, в среднем, я думаю, пространственная сложность равна O(n/2)
*/

package main

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

func main21() {
	stdReader := bufio.NewReader(os.Stdin)
	// пишу это, потому что в тот раз вы меня упрекнули за то, что не было сканера) Мне нравится сканер, но я тесирую через терминал где не приходит eof, и собсна, если
	// не прописано сколько строк придет, то это не мой вариант
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
