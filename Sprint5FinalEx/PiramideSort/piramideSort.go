package piramideSort

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type CompetitionResult struct {
	name       string
	completeEx int
	penalty    int
}

type PriorityQueue struct {
	queue         []CompetitionResult
	lastItemIndex int
}

func (pq *PriorityQueue) Append(cr CompetitionResult) {
	pq.lastItemIndex++
	pq.queue[pq.lastItemIndex] = cr
	pq.SiftUpRebalance(pq.lastItemIndex)
}

func (pq *PriorityQueue) Next() CompetitionResult {

	defer func() {
		pq.queue[1] = pq.queue[pq.lastItemIndex]
		pq.queue = pq.queue[:(len(pq.queue) - 1)]
		pq.SiftDownRebalance(1)
		pq.lastItemIndex--
	}()
	return pq.queue[1]
}

func (pq PriorityQueue) SiftUpRebalance(index int) {
	if index == 1 {
		return
	}
	parrentIndex := index >> 1
	if pq.Less(index, parrentIndex) {
		pq.Swap(index, parrentIndex)
		pq.SiftUpRebalance(parrentIndex)
	}

}

func (pq PriorityQueue) SiftDownRebalance(index int) {
	left := index * 2
	right := index*2 + 1
	if left >= len(pq.queue) {
		return
	}
	best := 0

	if right < len(pq.queue) && pq.Less(right, left) {
		best = right
	} else {
		best = left
	}

	if pq.Equals(index, best) {
		return
	}

	if pq.Less(best, index) {
		pq.Swap(best, index)
		pq.SiftDownRebalance(best)
	}
}

func (pq PriorityQueue) Less(a, b int) bool {
	if pq.queue[a].completeEx == pq.queue[b].completeEx {
		if pq.queue[a].penalty == pq.queue[b].penalty {
			return compareStings(pq.queue[a].name, pq.queue[b].name)
		} else {
			return pq.queue[a].penalty <= pq.queue[b].penalty
		}
	} else {
		return pq.queue[a].completeEx >= pq.queue[b].completeEx
	}
}

func (pq PriorityQueue) Swap(a, b int) {
	pq.queue[a], pq.queue[b] = pq.queue[b], pq.queue[a]
}

func (pq PriorityQueue) Equals(a, b int) bool {
	return pq.queue[a] == pq.queue[b]
}

func PiramideSort(r io.Reader, w io.Writer) {
	scaner := bufio.NewScanner(r)
	scaner.Scan()
	n, _ := strconv.Atoi(scaner.Text())
	pq := PriorityQueue{queue: make([]CompetitionResult, (n + 1)), lastItemIndex: 0}

	for i := 0; i < n; i++ {
		scaner.Scan()
		fields := strings.Fields(scaner.Text())
		name := fields[0]
		completeEx, _ := strconv.Atoi(fields[1])
		penalty, _ := strconv.Atoi(fields[2])
		pq.Append(CompetitionResult{name, completeEx, penalty})
	}
	writer := bufio.NewWriter(w)
	for i := 0; i < n; i++ {
		writer.WriteString(pq.Next().name)
		writer.WriteByte('\n')
	}
	writer.Flush()
	//
	//
}

func compareStings(s1, s2 string) bool {
	minLenght := 0
	if len(s1) > len(s2) {
		minLenght = len(s2)
	} else {
		minLenght = len(s1)
	}
	i, sum1, sum2 := 0, 0, 0
	for i < minLenght && sum1 == sum2 {
		sum1 += int(s1[i])
		sum2 += int(s2[i])
		i++
	}
	return sum1 <= sum2
}
