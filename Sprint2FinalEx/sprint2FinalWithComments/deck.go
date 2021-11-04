/*
-- ПРИНЦИП РАБОТЫ --
id в контесте 55521327
Привет :) Ну, сразу к задаче, решил опять по тупому, потому что контест не ругается на такие решения)
Задается дека с размером слайса от n+1, начальные значения (для фронтовой части 1, для бековой 0, разные значения, чтобы они себя не перезаписывали в самом начале). Дальше просто
регулярками проверяем методы и вызываем их (наверняка есть метод покрасивее, но я не придумал такой). В самих методах новый индекс высчитывается банальным
инкрементом/декрементом(increaceToTheLimit, decreaceToTheZero) где при достижении лимита (0 или n) возвращаются противоположые значения (n или 0).



-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
Благодря анолаогичным друг для друга функциям инкременитрования/декременитрования ,
pop и push корректно работают, а выйти за пределы им не позволяет isFull и isEmpty

(не думаю, что я правильно доказал, поэтому для наглядности описал внизу)

Пример корректной работы на почти заполненной деке
...34___89... front == 5, back == 7, maxLength - currentLenght == 3
pushFront(5), pushBack(7)
...345_789... front == 6, back == 6, maxLength - currentLenght == 1
pushBack(6)
...3456789... front == 6, back == 5, isFull == true
popFront()
front == 5, back == 5, maxLength - currentLenght == 1 ...34_6789...
(6 там конечно отсается, но про она забывается, и потом перезаписывается как-будто ее и не было)



-- ВРЕМЕННАЯ СЛОЖНОСТЬ --

И добавление и снятие с обеих сторон O(1)

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --

При инициализации деки создается массив размера n, и дополнительные перменные для подсчета индексов и т.д. , если их округлить, то я считаю,
что пространтсвенная сложность равна O(n)

*/

package main

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strconv"
)

type LimitDeck struct {
	deck          []string
	maxLength     int
	front         int
	back          int
	currentLength int
}

func (l LimitDeck) isFull() bool {
	return l.maxLength == l.currentLength
}
func (l LimitDeck) isEmpty() bool {
	return l.currentLength == 0
}

func (l *LimitDeck) PushFront(val string) error {
	if l.isFull() {
		return errors.New("error")
	}
	l.deck[l.front] = val

	l.front = increaceToTheLimit(l.front, l.maxLength)
	l.currentLength++
	return nil
}

func (l *LimitDeck) PushBack(val string) error {
	if l.isFull() {
		return errors.New("error")
	}
	l.deck[l.back] = val
	l.back = decreaceToTheZero(l.back, l.maxLength)
	l.currentLength++
	return nil
}
func (l *LimitDeck) PopFront() (string, error) {
	if l.isEmpty() {
		return "", errors.New("error")
	}

	l.front = decreaceToTheZero(l.front, l.maxLength)
	l.currentLength--
	return l.deck[l.front], nil
}

func (l *LimitDeck) PopBack() (string, error) {
	if l.isEmpty() {
		return "", errors.New("error")
	}
	l.back = increaceToTheLimit(l.back, l.maxLength)
	l.currentLength--
	return l.deck[l.back], nil
}

// нейминг мех, я знаю :(
func increaceToTheLimit(val, limit int) int {
	val++
	if val > limit {
		val = 0
	}
	return val
}
func decreaceToTheZero(val, goToAfterZero int) int {
	val--
	if val < 0 {
		val = goToAfterZero
	}
	return val
}

var PUSH_BACK = regexp.MustCompile(`^push_back`)
var PUSH_FRONT = regexp.MustCompile(`^push_front`)
var POP_BACK = regexp.MustCompile(`^pop_back`)
var POP_FRONT = regexp.MustCompile(`^pop_front`)
var INT = regexp.MustCompile(`-?\d+`)

func main22() {
	stdReader := bufio.NewReader(os.Stdin)
	stdScanner := bufio.NewScanner(stdReader)
	stdScanner.Scan()
	nStr := stdScanner.Text()
	n, _ := strconv.Atoi(nStr)
	stdScanner.Scan()
	lengthStr := stdScanner.Text()
	length, _ := strconv.Atoi(lengthStr)
	deck := LimitDeck{deck: make([]string, (length + 1)), maxLength: length, back: 0, front: 1, currentLength: 0}
	stdWriter := bufio.NewWriter(os.Stdout)
	for i := 0; i < n; i++ {
		stdScanner.Scan()
		switch {
		case PUSH_BACK.MatchString(stdScanner.Text()):
			value := INT.FindString(stdScanner.Text())
			if err := deck.PushBack(value); err != nil {
				writeWithNewLine(stdWriter, err.Error())
			}
		case PUSH_FRONT.MatchString(stdScanner.Text()):
			value := INT.FindString(stdScanner.Text())
			if err := deck.PushFront(value); err != nil {
				writeWithNewLine(stdWriter, err.Error())
			}
		case POP_BACK.MatchString(stdScanner.Text()):
			val, err := deck.PopBack()
			if err != nil {
				writeWithNewLine(stdWriter, err.Error())
			} else {
				writeWithNewLine(stdWriter, val)
			}

		case POP_FRONT.MatchString(stdScanner.Text()):
			val, err := deck.PopFront()
			if err != nil {
				writeWithNewLine(stdWriter, err.Error())
			} else {
				writeWithNewLine(stdWriter, val)
			}
		}
	}
	stdWriter.Flush()
}

func writeWithNewLine(w *bufio.Writer, s string) {
	w.WriteString(s)
	w.WriteByte('\n')
}

// если нужно использовать весь дек
// if (l.end == l.maxLength) {
// 	l.end = 0
// } else {
// 	l.end++
// }
