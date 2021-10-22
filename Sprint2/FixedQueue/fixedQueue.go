package fixedqueue

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strconv"
)

type SizedQueue struct {
	max   int
	queue []string
	start int
	end   int
	size  int
}

func (q *SizedQueue) Push(i string) error {
	if q.size == q.max {
		return errors.New("error")
	}
	q.queue[q.end] = i
	q.end = (q.end + 1) % q.max
	q.size++
	return nil
}

func (q *SizedQueue) Pop() (string, error) {
	if q.size == 0 {
		return "", errors.New("None")
	}
	res := q.queue[q.start]
	q.queue[q.start] = ""
	q.start = (q.start + 1) % q.max
	q.size--
	return res, nil
}

func (q SizedQueue) Peek() (string, error) {
	if q.size == 0 {
		return "", errors.New("None")
	}
	return q.queue[q.start], nil
}

func (q SizedQueue) Size() string {
	return strconv.Itoa(q.size)
}

var SIZE = regexp.MustCompile(`(?i)size`)
var PEEK = regexp.MustCompile(`(?i)peek`)
var POP = regexp.MustCompile(`(?i)pop`)
var PUSH = regexp.MustCompile(`(?i)^push`)
var INT = regexp.MustCompile(`-?\d+`)

func FixedQueue() {
	stdReader := bufio.NewReader(os.Stdin)
	stdScanner := bufio.NewScanner(stdReader)
	stdScanner.Scan()
	n, _ := strconv.Atoi(stdScanner.Text())
	stdScanner.Scan()
	length, _ := strconv.Atoi(stdScanner.Text())
	sizedQueue := SizedQueue{length, make([]string, length), 0, 0, 0}
	stdWriter := bufio.NewWriter(os.Stdout)
	for i := 0; i < n; i++ {
		stdScanner.Scan()
		switch {
		case SIZE.MatchString(stdScanner.Text()):
			stdWriter.WriteString(sizedQueue.Size())
			stdWriter.WriteByte('\n')
		case POP.MatchString(stdScanner.Text()):
			result, err := sizedQueue.Pop()
			if err != nil {
				stdWriter.WriteString(err.Error())
				stdWriter.WriteByte('\n')
				break
			}
			stdWriter.WriteString(result)
			stdWriter.WriteByte('\n')
		case PUSH.MatchString(stdScanner.Text()):
			pushItem := INT.FindString(stdScanner.Text())
			if err := sizedQueue.Push(pushItem); err != nil {
				stdWriter.WriteString(err.Error())
				stdWriter.WriteByte('\n')
			}
		case PEEK.MatchString(stdScanner.Text()):
			result, err := sizedQueue.Peek()
			if err != nil {
				stdWriter.WriteString(err.Error())
				stdWriter.WriteByte('\n')
				break
			}
			stdWriter.WriteString(result)
			stdWriter.WriteByte('\n')
		}
	}
	stdWriter.Flush()
}
