package Array_Manipulation

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

/*
 * Complete the 'arrayManipulation' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. 2D_INTEGER_ARRAY queries
 */

type a struct {
	to      int32
	howMuch int32
	result  int64
	q       *PriorityQueue
}

type PriorityQueue struct {
	queue []a
	index int
}

func (p *PriorityQueue) push(v a) {
	p.queue[p.index] = v
	p.rebalanceTreeUp(p.index)
	p.index++
}

func (p *PriorityQueue) pop() a {
	item := p.queue[1]
	p.queue[1], p.queue[p.index-1] = p.queue[p.index-1], p.queue[1]
	p.index--
	p.rebalanceTreeDown(1)
	return item
}

func (p *PriorityQueue) rebalanceTreeUp(index int) {
	if index == 1 {
		return
	}
	parrentIndex := index / 2
	if p.queue[index].to < p.queue[parrentIndex].to {
		p.queue[index], p.queue[parrentIndex] = p.queue[parrentIndex], p.queue[index]
		p.rebalanceTreeUp(parrentIndex)
	}

}

func (p *PriorityQueue) rebalanceTreeDown(index int) {
	childLeft := index * 2
	childRight := childLeft + 1
	if childLeft >= p.index {
		return
	}

	best := 0

	if childRight < p.index && p.queue[childRight].to < p.queue[childLeft].to {
		best = childRight
	} else {
		best = childLeft
	}
	if p.queue[best].to < p.queue[index].to {
		p.queue[best], p.queue[index] = p.queue[index], p.queue[best]
		p.rebalanceTreeDown(best)
	}
}

func (p *PriorityQueue) isEnd(index int) bool {
	if p.index == 1 {
		return false
	}
	return index == int(p.queue[1].to)
}

func arrayManipulation(n int32, queries [][]int32) int64 {
	// Write your code here
	piramide := PriorityQueue{
		queue: make([]a, len(queries)+1),
		index: 1,
	}
	result := make([][]a, n)
	for _, v := range queries {
		result[v[0]-1] = append(result[v[0]-1], a{
			to:      v[1],
			howMuch: v[2],
		})
	}
	var r int64 = 0
	var accum int64 = 0
	r2 := make([]int64, n)
	for i := 0; i < len(result); i++ {
		for piramide.isEnd(i) {
			v := piramide.pop()
			accum -= int64(v.howMuch)
		}
		for _, v := range result[i] {
			piramide.push(v)
			accum += int64(v.howMuch)
		}
		r2[i] = accum
		if accum > r {
			r = accum
		}
	}
	return r
}

func PrepareData(r io.Reader, w io.Writer) {
	reader := bufio.NewReaderSize(r, 16*1024*1024)

	writer := bufio.NewWriterSize(w, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	var queries [][]int32
	for i := 0; i < int(m); i++ {
		queriesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != 3 {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	result := arrayManipulation(n, queries)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
