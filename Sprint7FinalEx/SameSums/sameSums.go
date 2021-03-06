/**
id в контесте 65440529
-- ПРИНЦИП РАБОТЫ --
Привет! Эта задача была легкая, я почти сразу вспомнил задачу про рюкзак которая разбиралась в теории и применил похожий трюк тут)
В начале мы берем общую сумму всех очков - sum, если sum % 2 != 0, то ответ сразу false, если все ок, то приступаем к динамике через 2 массива
dpPrev и dp, которые будут обновляться каждую итерацию. Сразу оговорюсь, что g будет считаться только до sum / 2 + 1, потому что мой ответ
на исходный вопрос будет находиться по индексу sum / 2
1. Что будет храниться в массиве dp?
	В массиве dp будет храниться то, сколькими способами можно набрать сумму g используя число data[i]
2. Каким будет базовый случай для задачи?
	Если очков вообще нет, то любую сумму можно набрать только 0 способов
3. Каким будет переход динамики?
	Набрать сумму g можно как минимум dpPrev[g], поэтому dp[g] всегда будет как минимум dpPrev[g]
	Если data[i] == g, то мы увеличиваем dp[g] на 1
	Когда g > data[i] мы проверяем можно ли каким-нибудь способом набрать сумму g - data[i] (проще говоря if dpPrev[g - data[i]) > 0),
 	и, если можно, то мы тоже увеличиваем dp[g] на 1
4. Каким будет порядок вычисления данных в массиве dp?
	Используя Динамическое программирование назад, смотря предыдущие значения
5. Где будет располагаться ответ на исходный вопрос?
	В массиве dpPrev[sum / 2], если dpPrev[sum / 2] > 1, то сумму точно можно разбить на 2 части (или больше), если нет то нет)
-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
Тут опять же корректность алгоритма состоит в базовом случае и переходе динамики:
Предположим у нас есть массив [2,2,2,2], сумма очков равна 8, мы ищем, то, сколькими способами можно набрать 8 / 2

Первая итерация - сумму 2 можно набрать 0 + 1 способами, больше, используя одно число набрать нельзя:
     1 2 3 4
  [0 0 0 0 0]
2 [0 0 1 0 0]
2 [0 0 0 0 0]
2 [0 0 0 0 0]
2 [0 0 0 0 0]

Вторая итерация - сумму 2 можно набрать 1 + 1 способами, также можно получить сумму 4, прибавив текущую 2 к той 2, которую мы получили в прошлой
итерации и + 0 (количество способов получить 4 из предыдущей итерации):
     1 2 3 4
  [0 0 0 0 0]
2 [0 0 1 0 0]
2 [0 0 2 0 1]
2 [0 0 0 0 0]
2 [0 0 0 0 0]

Третья итерация - сумму 2 можно набрать 2 + 1 способами, и сумму 4, так же как и в прошлой итерации + 1 (количество способов получить 4 из предыдущей итерации):
     1 2 3 4
  [0 0 0 0 0]
2 [0 0 1 0 0]
2 [0 0 2 0 1]
2 [0 0 3 0 2]
2 [0 0 0 0 0]

Ну и последняя:
     1 2 3 4
  [0 0 0 0 0]
2 [0 0 1 0 0]
2 [0 0 2 0 1]
2 [0 0 3 0 2]
2 [0 0 4 0 3]

В итоге получаем 3, это >= 2, поэтому ответ True
-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
Обычная двумерная динамика O(n * (sum / 2)) = O(n * sum)
-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
O(sum * 2 + n) = O(sum + n)

P.S.
Благополучно опоздал к окончанию курса, прям по гиту видно, куда я закидываю свои задачи, что после новогодних праздников у меня совсем не было времени,
очередной дедлайн был, все бегали туда-сюда) Но самым важным фактором, по моему, была корона, теперь я в полной мере почувствовал уменьшение продуктивности,
про которое говорят люди. Заставлять себя работать еще норм, потому что использую те знания, которые у меня есть, но вот учиться очень сложно :(

И насчет тестов в отдельном пакете, тут все просто) Еще на втором или третьем спринте я сделал в main файле пару функций для создания файлов задач, которые
будут в этом спринте https://puu.sh/ILEAA/fbe4a65d4f.png , и там я сделал тесты в отдельном пакете, а менять лень) Поэтому у меня тут пустой ридми файл, например))
*/

package sameSums

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Matrix [][]int

func (m Matrix) String() string {
	s := ""
	for _, ints := range m {
		s += fmt.Sprint(ints)
		s += string('\n')
	}
	return s
}

func SameSums(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	inputData := strings.Fields(scanner.Text())
	data := make([]int, n)
	sum := 0
	for i := 0; i < n; i++ {
		data[i], _ = strconv.Atoi(inputData[i])
		sum += data[i]
	}
	writer := bufio.NewWriter(w)
	if sum%2 != 0 {
		writer.WriteString("False")
		writer.Flush()
		return
	}
	dpPrev := make([]int, sum/2+1)
	dp := make([]int, sum/2+1)
	for i := 1; i < n+1; i++ {
		for g := 1; g < sum/2+1; g++ {
			dp[g] = dpPrev[g]
			if g == data[i-1] {
				dp[g]++
			}
			if g > data[i-1] && dpPrev[g-data[i-1]] > 0 {
				dp[g]++
			}
		}
		dpPrev = dp
		dp = make([]int, sum/2+1)
	}
	if dpPrev[sum/2] > 1 {
		writer.WriteString("True")
	} else {
		writer.WriteString("False")
	}

	writer.Flush()

}

func main() {
	SameSums(os.Stdin, os.Stdout)
}
