package listqueue

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strconv"
)

type Node struct {
	data string
	next *Node
}

type LinkedListQueue struct {
	size  int
	first *Node
	last  *Node
}

func (lq *LinkedListQueue) Get() (string, error) {
	if lq.size == 0 {
		return "", errors.New("error")
	}
	defer func() {
		if lq.size == 0 {
			lq.first = nil
			lq.last = nil
		} else {
			lq.first = lq.first.next
		}

	}()
	lq.size--
	return lq.first.data, nil

}

func (lq *LinkedListQueue) Put(s string) {
	node := &Node{data: s, next: nil}
	if lq.size == 0 {
		lq.first = node
		lq.last = node
	} else {
		lq.last.next = node
		lq.last = lq.last.next
	}
	lq.size++
}

func (lq LinkedListQueue) Size() string {
	return strconv.Itoa(lq.size)
}

var GET = regexp.MustCompile(`(?i)^(get)`)
var PUT = regexp.MustCompile(`(?i)^(put)`)
var SIZE = regexp.MustCompile(`(?i)^(size)`)
var INT = regexp.MustCompile(`-?\d+`)

func ListQueue() {
	stdReader := bufio.NewReader(os.Stdin)
	stdScanner := bufio.NewScanner(stdReader)
	stdScanner.Scan()
	n, _ := strconv.Atoi(stdScanner.Text())
	lLQ := LinkedListQueue{0, nil, nil}
	stdWriter := bufio.NewWriter(os.Stdout)
	for i := 0; i < n; i++ {
		stdScanner.Scan()
		switch {
		case GET.MatchString(stdScanner.Text()):
			item, err := lLQ.Get()
			if err != nil {
				stdWriter.WriteString(err.Error())
			} else {
				stdWriter.WriteString(item)
			}
			stdWriter.WriteByte('\n')
		case PUT.MatchString(stdScanner.Text()):
			i := INT.FindString(stdScanner.Text())
			lLQ.Put(i)
		case SIZE.MatchString(stdScanner.Text()):
			stdWriter.WriteString(lLQ.Size())
			stdWriter.WriteByte('\n')
		}
	}
	stdWriter.Flush()
}
