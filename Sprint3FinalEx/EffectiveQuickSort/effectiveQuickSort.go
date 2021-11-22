/*
-- ПРИНЦИП РАБОТЫ --
id в контесте 58675957

Привет) Сразу напишу, что это не я придумал) Где-то с пол года назад я читал и реализовывал чтобы понять как это работает
вот эту статью https://yourbasic.org/golang/quicksort-optimizations/ собственно весь принцип работы можно узнать отсюда, но я перескажу:
Берется массив, если для его длинны не имеет особенного смысла делить его с помощью partition, применяется сортировка вставками,
я взял длину меньше 6, не знаю почему, просто так) Если массив больше, то сначала вычисляется pivot с помощью 3 элементов массива выбранных рандомно -
в статье говорится (ну и в этом есть смысл, по-моему), что выбирать рандомный элемент рискованно, также как и выбирать 3 элемента из начала, середины
и конца, и лучше просто соединить 2 эти идеи. Потом начинается в partiton передается сам массив, и элемент под индексом pivot'a. Суть в том, чтобы
получать элемент, и, сравнивая его с pivot'ом, класть его либо влево, либо вправо, переменная mid нужна для сохранения pivot между меньшими и большими значениями.
Потом функция отдает 2 переменные low и high. low - это тот индекс до которого значения меньше pivot, high - это тот индекс от которого начинаются значения выше.
Разрезав слайс с помощью этих индексов, мы начинаем новый цикл рекурсии со значениями выше и ниже, значения равные (у нас оно одно, потому что все элементы уникальны) остаются на месте
и в дальнейшей сортировке не участвуют. Для работы я создал структуру CompetitionResult которая содержит данные о каждом участнике соревнования и CompetitionData - слайс, который
частично удовлетворяет интерфейсу sort.Interface, у него есть Less и Swap, которые много где будут использоваться. Остальные уточнения далее по коду)



-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --

Всегда видимо будут проблемы с этим разделом, я все еще плохо понимаю что тут надо писать)
Из задания следует, что нужно реализовать быструю сортировку без использования доп памяти, все это тут есть, перекладывание элементов в partiton происходит по индексу, и дополнительного
массива не требует, а сам алгоритм это и есть быстрая сортировка, просто немного улучшенная)



-- ВРЕМЕННАЯ СЛОЖНОСТЬ --

в худшем случае работает за O(n^2) если pivot будет выбран плохо в каждой итерации рекурсии, в среднем работает за O(log n * n) где log n это количество рекурсий, а n это производимые
действия над массивом (надеюсь правильно описал)

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --

дополнительной памяти, кроме переменных, не выделяется, значит O(1)


== зы ==
из-за разных приколов в жизни и сдачи проекта на работе я опаздываю к жесткому дедлайну, я отсылаю это 20 числа, а финальные задания 4го спринта мне нужно сдать уже 28, есть шанс того
что я останусь на "второй год" :D я просто не знаю меняются ли ревьюеры, но если меняются, то спасибо вам большое! Как минимуму благодаря вам я узнал что регулярки это медленно,
и начал делать тесты, которые, кстати, идут в комплекте с этим файлом 😎
*/

package effectivequicksort

import (
	"bufio"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type CompetitionResult struct {
	name       string
	completeEx int
	penalty    int
}

// испльзовал подход как в интерфейсе sort.Interface

type CompetitionData []CompetitionResult

func (cd CompetitionData) Less(a, b int) bool {
	if cd[a].completeEx == cd[b].completeEx {
		if cd[a].penalty == cd[b].penalty {
			return compareStings(cd[a].name, cd[b].name)
		} else {
			return cd[a].penalty < cd[b].penalty
		}
	} else {
		return cd[a].completeEx > cd[b].completeEx
	}
}

func (cd CompetitionData) Swap(a, b int) {
	cd[a], cd[b] = cd[b], cd[a]
}

func EffectiveQUickSort(r io.Reader, w io.Writer) {
	rand.Seed(time.Now().Unix())
	reader := bufio.NewReader(r)
	scaner := bufio.NewScanner(reader)
	scaner.Scan()
	// n, кстати, был бы не нужен, если бы приходили данные без каких-то комментарииев, которые мне испортили первую попытку)
	// например https://puu.sh/Iqc6I/973aeb904a.png
	n, _ := strconv.Atoi(scaner.Text())
	competitionData := make(CompetitionData, 0)
	for i := 0; i < n; i++ {
		scaner.Scan()
		stringData := strings.Fields(scaner.Text())
		completeEx, _ := strconv.Atoi(stringData[1])
		penalty, _ := strconv.Atoi(stringData[2])
		competitionResult := CompetitionResult{
			stringData[0], completeEx, penalty,
		}
		competitionData = append(competitionData, competitionResult)
	}
	quickSort(competitionData)
	writer := bufio.NewWriter(w)
	for _, v := range competitionData {
		writer.WriteString(v.name)
		writer.WriteByte('\n')
	}
	writer.Flush()
}

func partition(d CompetitionData, pivot CompetitionResult) (low, high int) {
	low, mid, high := 0, 0, len(d)
	for mid < high {
		if d[mid] == pivot {
			mid++
			continue
		}
		if compareResultsHelper(d[mid], pivot) {
			d.Swap(mid, low)
			low++
			mid++
		} else {
			d.Swap(mid, high-1)
			high--

		}
	}
	return
}

func quickSort(d CompetitionData) {
	if len(d) < 6 {
		insertionSort(d)
		return
	}
	p := pivot(d)
	low, high := partition(d, d[p])
	quickSort(d[:low])
	quickSort(d[high:])

}

func pivot(d CompetitionData) int {
	return median(d, rand.Intn(len(d)), rand.Intn(len(d)), rand.Intn(len(d)))
}

func median(d CompetitionData, a, b, c int) int {
	if d.Less(a, b) {
		if d.Less(b, c) {
			return b
		}
		if d.Less(c, a) {
			return a
		}
		return c
	}
	if d.Less(c, b) {
		return b
	}
	if d.Less(c, a) {
		return c
	}
	return a
}

func insertionSort(d CompetitionData) {
	for i := 1; i < len(d); i++ {
		for g := i; g > 0 && d.Less(g, g-1); g-- {
			d.Swap(g, g-1)
		}
	}
}

// так как pivot отдельная переменная, для сравнения его с другим элементом массива, я сделал вот такую функцию
func compareResultsHelper(a, b CompetitionResult) bool {
	customCompetitonData := CompetitionData{a, b}
	return customCompetitonData.Less(0, 1)
}

//изначально я сравнивал первые биты никнеймов, но потом понял, что они различия могут быть и глубже, поэтому сделал вот это
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
	return sum1 < sum2
}

func main() {
	EffectiveQUickSort(os.Stdin, os.Stdout)
}
