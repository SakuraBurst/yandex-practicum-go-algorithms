/*
-- ПРИНЦИП РАБОТЫ --
id в контесте 59829334

Привет! На самом деле мне не очень нравится это решение, но я уже на неделю опаздываю, а оно рабочее и судя по тестам даже неплохое
Плохо оно справляется только с последним тестом, где, я так понял, просто миллион документов с одним и тем же словом, и миллион запросов
с такими же строками, но во время я там укладываюсь)

Суть до смешного проста, поэтому мне и не нравится:
Есть мапа с ключом - словом, которое приходит в документах, и значением - тоже мапа с ключом - индекс документа, и значением - сколько раз в этом документе
появлялось это слово.

Дальше я прохожусь по строке запроса и по сути перекладываю результаты подсчетов в другую мапу где ключ - это индекс документа,
значение сколько раз в этом документе попадались слова которые есть в поисковом запросе

Потом я прохожусь по этой мапе результатов, складываю результат в структуру по которой в последствии буду сортировать

Дальше происходит сортировка, я выбрал сортировку слиянием потому-что я поглядел на то, сколько мне дается памяти
и подумал, почему бы ее не использовать на максимум)

Дальше вывожу первые 5 или меньше правильных ответов

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --

Подсчитывая сумму вхождений слов в документы, а потом используя это на этапе получения поисковой строки мы получаем корректное решение

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
N - количество документов
Ni - количество слов в документе
M - количество запросов
Mi - количество слов в запросе

Мы записываем все слова в документе за O(N * Ni)

Мы проверяем все слова в каждом запросе за O(M * Mi)

Мы получаем подсчет вхождения слова во все документы за O(N) (в среднем за O(N/2) потому-что маловероятно, что слово могло попасть во все документы
но во первых такое может быть, во вторых O(N/2) = O(N))

Получается чтобы получить все вхождения слов поискового запроса из всех документов нам понадобится O(Mi * N)

Результат подсчетов я помечу как R

Дальше нам нужно положить все результаты в массив за O(R)

И отсортировать за (logR * R)

Вывести результат за O(1)

В результате получается что-то типа O(N * Ni + M * Mi * N + R + logR * R)

Я не умею сокращать и даже пробовать не буду, но если я все правильно подсчитал, то это кокой-то ужас. Либо я неправильно подсчитал
либо мое решение не имеет права на жизнь


-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
N - количество документов
Ni - количество слов в документе
R - Результат подсчетов
Все почти тоже самое, что и с временной сложностью

Мы храним подсчеты всех слов в документе O(N * Ni)

Мы храним мапу результатов O(R)

И такой же массив O(R)

Используем сортировку слиянием которая каждый шаг кроме базового случая хранит в себе в пике O(R + R)

Получаем в пике O(N * Ni + R + R + R + R)

Ну тут, думаю, все же можно сократить до O(N * Ni + 4R)

Вообще я, конечно, понимаю, что хранить одновременно мапу и массив с одними и теми же результатами не эффективно, но опять же, посмотрев на то,
сколько мне дается памяти, я решил не мелочиться)

*/

package searchSystem

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type SearchResult struct {
	index      int
	wordsFound int
}

func SearchSystem(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	scaner := bufio.NewScanner(reader)
	scaner.Scan()
	nDocks, _ := strconv.Atoi(scaner.Text())
	searchIndexesMap := map[string]map[int]int{}
	for d := 0; d < nDocks; d++ {
		scaner.Scan()
		for _, v := range strings.Fields(scaner.Text()) {
			if searchIndexesMap[v] == nil {
				searchIndexesMap[v] = map[int]int{}
			}
			searchIndexesMap[v][d]++
		}
	}
	scaner.Scan()
	kSearchStrings, _ := strconv.Atoi(scaner.Text())
	writer := bufio.NewWriter(w)
	for s := 0; s < kSearchStrings; s++ {
		scaner.Scan()
		result := make(map[int]int, len(scaner.Text()))
		// для подсчета только уникальных слов
		alreadyCounted := make(map[string]bool)
		for _, w := range strings.Fields(scaner.Text()) {
			if !alreadyCounted[w] {
				for k, wl := range searchIndexesMap[w] {
					result[k] += wl
				}
				alreadyCounted[w] = true
			}

		}
		counterArr := make([]SearchResult, 0, len(result))
		for k, v := range result {
			counterArr = append(counterArr, SearchResult{index: k + 1, wordsFound: v})
		}
		counterArr = merge(counterArr)
		for i, v := range counterArr {
			if i == 5 {
				break
			}
			writer.WriteString(strconv.Itoa(v.index))
			writer.WriteByte(' ')
		}
		writer.WriteByte('\n')

	}
	writer.Flush()
	//
}

func merge(d []SearchResult) []SearchResult {
	if len(d) <= 1 {
		return d
	}
	mid := len(d) >> 1
	firstHalf := merge(d[:mid])
	secondHalf := merge(d[mid:])
	firstHalfIndex := 0
	secondHalfIndex := 0
	result := make([]SearchResult, 0, len(firstHalf)+len(secondHalf))
	for firstHalfIndex < len(firstHalf) || secondHalfIndex < len(secondHalf) {
		// если закончилась первая плоловина
		if firstHalfIndex == len(firstHalf) {
			result = append(result, secondHalf[secondHalfIndex])
			secondHalfIndex++
			continue
		}
		// если закончилась вторая половина
		if secondHalfIndex == len(secondHalf) {
			result = append(result, firstHalf[firstHalfIndex])
			firstHalfIndex++
			continue
		}
		// если обе половины что-то содержат
		if compareResults(firstHalf[firstHalfIndex], secondHalf[secondHalfIndex]) {
			result = append(result, firstHalf[firstHalfIndex])
			firstHalfIndex++
		} else {
			result = append(result, secondHalf[secondHalfIndex])
			secondHalfIndex++
		}
	}
	return result
}

func compareResults(a, b SearchResult) bool {
	if a.wordsFound == b.wordsFound {
		return a.index < b.index
	}
	return a.wordsFound > b.wordsFound
}

func main() {
	SearchSystem(os.Stdin, os.Stdout)
}
