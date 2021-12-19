/*
-- ПРИНЦИП РАБОТЫ --
id в контесте 61709867

Привет! Ну, тут было совсем легко благодаря двум факторам. Абсолютно такая же сортировка требовалась в финальной задаче
третьего спринта по реализации эффективной быстрой сортировки, так что я без зазрений совести многое взял оттуда, включая тип
и реализацию интерфейса sort.Interface. Второй фактор это то, что самый сложный элемент в этой сортировке, просеивание вверх
и вниз, были разжёваны в теории + к тому же еще и требовались на практике в одноименных задачах)

Суть решения: У меня есть структура PriorityQueue (далее pq) у которой есть поля queue - []CompetitionResult и lastItemIndex.

На самом деле сразу хочу отметить что поле lastItemIndex в дальнейшем описании будет совершенно необязательно, его можно будет
очень легко заменить на pq.Len(), а если поле необязательно, то и сама структура не нужна и можно было бы сделать просто именованный тип слайса,
но почему-то я решил сделать так, и переделывать не хочу)

У структуры есть много методов, опишу каждый
(возможно термин ребалансирование неверный, если так, то прошу прощения))

pq.Append(elem): добавляет элемент в очередь, после чего ребалансирует кучу с помощью просеивания вверх этого элемента (pq.SiftUpRebalance(pq.lastItemIndex))

pq.Next()elem: возвращает самый первый (приоритетный) элемент из кучи, после чего (с помощью defer) ставит последний элемент на место первого
уменьшает слайс на 1 элемент чтобы удалить перемещенный элемент из конца, и делает ребалансирование с помощью просеивания вниз
(pq.SiftDownRebalance(1)) чтобы первый элемент стал самым приоритетным, а элемент который был забран с конца занял правильное место

pq.SiftUpRebalance(index): работает очень просто - если родитель текущего индекса менее приоритетный, то текущий элемент
меняется местами с родителем и запускается новый цикл рекурсии, в который передается уже новый индекс нашего элемента (то есть старый индекс родителя)
Базовый случай тут если родитель более приоритетный, либо если текущий индекс == 1, и итем является самым приоритетным.

pq.SiftDownRebalance(index): работает уже посложнее. Если есть оба предка, то тогда они сравниваются на то, какой приоритетнее,
и самый приоритетный из пары сравнивается с текущим элементом, если текущий элемент менее приоритетный, то он меняется местами
с самым приоритетным элементом. Если есть только левый предок, то, по очевидным причинам, сравнения пары предков не происходит,
в остальном все тоже самое.
Далее запускается новый цикл рекурсии куда передается новый индекс элемента (то есть старый индекс самого приоритетного элемента)
Базовый случай тут если левого предка нет, либо элемент на своем текущем месте самый приоритетный

pq.Less(a,b): проверяет приоритетнее ли a чем b

pq.Swap(a,b): меняет местами a и b

pq.Len(): возвращает размер очереди

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
Вообщем-то все просто, pq.Append(elem) добавляет итем в очередь, поддерживая то, чтобы в самом начале был самый приоритетный элемент
с помощью pq.SiftUpRebalance()
После того, как все итемы добавлены, начинаем постепенно забирать самые приоритетные элементы с помощью pq.Next(), который
поддерживает то, чтобы самый второй по приоритетности элемент оказался на вершине кучи с помощью pq.SiftDownRebalance()

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
В целом временная сложность равна O(log n)

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
Создается только слайс размера n и пару доп переменных, так что Пространственная сложность равна O(n)
*/

package piramideSort

import (
	"bufio"
	"io"
	"os"
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
		pq.queue = pq.queue[:pq.lastItemIndex]
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
	if left >= pq.Len() {
		return
	}
	best := 0

	if right < pq.Len() && pq.Less(right, left) {
		best = right
	} else {
		best = left
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

func (pq PriorityQueue) Len() int {
	return len(pq.queue)
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
	s1Lenght := len(s1)
	s2Lenght := len(s2)
	minLenght := 0
	if s1Lenght > s2Lenght {
		minLenght = s2Lenght
	} else {
		minLenght = s1Lenght
	}
	i, sum1, sum2 := 0, 0, 0
	for i < minLenght && sum1 == sum2 {
		sum1 += int(s1[i])
		sum2 += int(s2[i])
		i++
	}
	if sum1 == sum2 {
		return s1Lenght < s2Lenght
	}
	return sum1 < sum2
}

func main() {
	PiramideSort(os.Stdin, os.Stdout)
}
