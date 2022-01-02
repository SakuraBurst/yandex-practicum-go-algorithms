package timeToGetOut

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	value    int
	pointsTo AdjacencyList
}

type AdjacencyList []*Node

func TimeToGetOut(r io.Reader, w io.Writer) {
	scaner := bufio.NewScanner(r)
	scaner.Scan()
	initData := strings.Fields(scaner.Text())
	n, _ := strconv.Atoi(initData[0])
	m, _ := strconv.Atoi(initData[1])
	al := make(AdjacencyList, n+1)
	for i := 0; i < m; i++ {
		scaner.Scan()
		edge := strings.Fields(scaner.Text())
		vertexAIndex, _ := strconv.Atoi(edge[0])
		vertexBIndex, _ := strconv.Atoi(edge[1])
		if al[vertexAIndex] == nil {
			al[vertexAIndex] = &Node{value: vertexAIndex}
		}
		if al[vertexBIndex] == nil {
			al[vertexBIndex] = &Node{value: vertexBIndex}
		}
		al[vertexAIndex].pointsTo = append(al[vertexAIndex].pointsTo, al[vertexBIndex])
	}

	var dfc func(n *Node) error
	colors := make([]string, n+1)
	timer := -1
	writer := bufio.NewWriter(w)
	timeIn := make([]string, n+1)
	timerOut := make([]string, n+1)
	dfc = func(n *Node) error {
		if n == nil {
			return errors.New("node doesen't exist")
		}
		timer++
		colors[n.value] = "black"
		timeIn[n.value] = strconv.Itoa(timer)
		n.pointsTo = merge(n.pointsTo)
		for _, v := range n.pointsTo {
			if colors[v.value] != "black" {
				err := dfc(v)
				if err != nil {
					return err
				}
			}
		}
		timer++
		timerOut[n.value] = strconv.Itoa(timer)
		return nil

	}
	err := dfc(al[1])
	if err != nil {
		_, err := writer.WriteString("0 1")
		if err != nil {
			return
		}
		err = writer.WriteByte('\n')
		if err != nil {
			return
		}
	} else {
		for i := 1; i < n+1; i++ {
			_, err := writer.WriteString(timeIn[i] + " " + timerOut[i])
			if err != nil {
				return
			}
			err = writer.WriteByte('\n')
			if err != nil {
				return
			}
		}
	}

	err = writer.Flush()
	if err != nil {
		return
	}

}

func merge(al AdjacencyList) AdjacencyList {
	if len(al) <= 1 {
		return al
	}
	mid := len(al) >> 1
	a := merge(al[:mid])
	b := merge(al[mid:])
	aIndex := 0
	bIndex := 0
	result := make(AdjacencyList, 0, len(a)+len(b))
	for aIndex < len(a) || bIndex < len(b) {
		if aIndex == len(a) {
			result = append(result, b[bIndex])
			bIndex++
			continue
		}
		if bIndex == len(b) {
			result = append(result, a[aIndex])
			aIndex++
			continue
		}
		if a[aIndex].value <= b[bIndex].value {
			result = append(result, a[aIndex])
			aIndex++
		} else {
			result = append(result, b[bIndex])
			bIndex++
		}
	}
	return result
}

func main() {
	TimeToGetOut(os.Stdin, os.Stdout)
}
