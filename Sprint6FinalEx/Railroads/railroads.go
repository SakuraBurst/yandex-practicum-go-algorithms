/*
-- ПРИНЦИП РАБОТЫ --
id в контесте 64325471

Привет! А вот эта задача была уже очень сложная) Впервые я потратил на задачу финальную задачу больше, чем 4 часа, решал я ее 5 дней. Как оказалось,
корректное решение мне пришло в голову еще на второй день, но я очень много времени потратил на оптимизацию, сначала памяти, потом времени. В целом, задача
сама по себе не простая, но когда я начинал финальные задачи меня впервые за все это смутное время поразил ковид, и теперь я понимаю о каком уменьшении iq
говорят люди) Первая задача была легкой, поэтому я даже в этом состоянии смог ее решить, а эта... Ну да ладно) Самой действенной оптимизацией памяти оказалась
сменить все int на uint с определенным количеством бит

В общем решил я ее через 2 списка смежности для каждого типа дорог, и матрицу достижимости (изначально это был []map[int]WayType, но попросив подсказку я
узнал про матрицу, это было эффективнее по времени + даже есть крутое название). Через каждый список прохожусь по dfs и обновляю, из каких вершин можно попасть
в другие вершину. По первому списку я прохожусь чтобы просто расставить значения первого типа в матрице, по второму я прохожусь и смотрю, есть ли другое значение
в ячейке матрицы, куда я хочу вставить значение. Если есть, и оно не равно значению текущего листа, то возвращаю false и прерываю рекурсию (так себе
описал, надеюсь в доказательстве корректности будет более яснее)

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --

Я подготовил 2 скрина, чтобы было понятнее)

https://puu.sh/IEkRM/7ed7ec1eed.png
https://puu.sh/IEkU7/3b0b3e9e98.png

я так себе художник)

Наверху показано состояние матрицы достижимости на каждой итерации цикла который на 91 строке.

Например, в верхнем правом углу первого скрина показано состояние матрицы достижимости после того, как dfs отработал - в 5 по пути B можно прийти через 1, 2 и 4

В каждой из dfs функций у нас есть массив маршрута, который хранит весь текущий путь. В самом начале функции в этот массив добавляется текущая обрабатываемая
вершина. Когда мы проходимся циклом по ребрам текущей вершины, мы смотрим, есть ли в смежной вершине хотя бы один элемент из пути, и, если есть, то переходим
к следующей вершине. Если нет, то переходим к смежной вершине. Благодаря тому, что мы в любом случае посетим все вершины хотя бы один раз, мы гарантированно имеем
корректную матрицу достижимости для первого типа дорог. Во втором цикле делаем все то же самое, что и в первом, за исключением того, что мы не только проверяем
на существование в смежной вершине хотя бы одного элемента текущего пути, но и проверяем правильного ли типа этот путь. Если неправильного - возвращаем false
и прекращаем рекурсию. В примере на скринах, это происходит на втором скрине, на самой первой итерации: когда мы приходим в вершину 3 текущий путь равен [2,3]
мы начинаем проверять вершину 5, есть ли к ней пути из таких же вершин, и сразу видим, что есть путь из 2 другого типа.


-- ВРЕМЕННАЯ СЛОЖНОСТЬ --

При переборе каждого ребра каждой вершины мы проходимся по массиву маршрута, который в пике может занимать O(|V|) памяти. Получается,
что временная сложность равна O(|V| + |V|^|E| + E) (я не совсем в этом уверен, какие-то у меня проблемы с асимптотикой графов

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --

2 списка смежности O(|V| + |E|), матрица достижимости O(|V|^2) ну и массив маршрута, который тоже может принимать до O(|V|) элементов.
Итого получаем O(|V| * 2 + |E| + |V|^2) = O(|V| + |E| + |V|^2)

*/

package railroads

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type WayType uint8

const (
	UnknownWay WayType = iota
	RWay
	BWay
)

type Matrix [][]WayType

type Roads map[int]WayType

type Edge uint16

type Vertex struct {
	value    uint16
	pointsTo []Edge
}

type AdjacencyList []Vertex

func Railroads(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	N, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	ALR := make(AdjacencyList, N+1)
	ALB := make(AdjacencyList, N+1)
	extractData(N, scanner, ALR, ALB)
	writer := bufio.NewWriter(w)
	roads := make(Matrix, N+1)
	for _, vertex := range ALR {
		if vertex.pointsTo == nil {
			continue
		}
		dfsOnlyCheck(ALR, vertex, nil, roads, RWay)
	}
	for _, vertex := range ALB {
		if vertex.pointsTo == nil {
			continue
		}
		ok := dfs(ALB, vertex, nil, roads, BWay)
		if !ok {
			writer.WriteString("NO")
			writer.Flush()
			return
		}
	}

	writer.WriteString("YES")

	writer.Flush()
}

func extractData(N int, scanner *bufio.Scanner, ALR AdjacencyList, ALB AdjacencyList) {
	for i := 1; i < N; i++ {
		scanner.Scan()
		railData := scanner.Text()
		if ALR[i].pointsTo == nil {
			ALR[i].pointsTo = make([]Edge, 0, N+1-i)
			ALR[i].value = uint16(i)
		}
		if ALB[i].pointsTo == nil {
			ALB[i].pointsTo = make([]Edge, 0, N+1-i)
			ALB[i].value = uint16(i)
		}
		for g, datum := range railData {
			if datum == 82 {
				ALR[i].pointsTo = append(ALR[i].pointsTo, Edge(i+g+1))
			} else {
				ALB[i].pointsTo = append(ALB[i].pointsTo, Edge(i+g+1))
			}
		}

	}
}

func dfsOnlyCheck(AL AdjacencyList, vertex Vertex, roadsWithCurrentWayType []uint16, roads Matrix, currentWay WayType) {
	roadsWithCurrentWayType = append(roadsWithCurrentWayType, vertex.value)
Upper:
	for _, edge := range vertex.pointsTo {
		// этот цикл у меня был в самом начале функции
		for _, i := range roadsWithCurrentWayType {
			if roads[edge] == nil {
				roads[edge] = make([]WayType, cap(roads))
			}
			roadWayType := roads[edge][i]
			if roadWayType == currentWay {
				continue Upper
			}

			roads[edge][i] = currentWay
		}
		dfsOnlyCheck(AL, AL[edge], roadsWithCurrentWayType, roads, currentWay)
	}
}

func dfs(AL AdjacencyList, vertex Vertex, roadsWithCurrentWayType []uint16, roads Matrix, currentWay WayType) bool {
	roadsWithCurrentWayType = append(roadsWithCurrentWayType, vertex.value)
Upper:
	for _, edge := range vertex.pointsTo {
		// этот цикл у меня был в самом начале функции
		for _, i := range roadsWithCurrentWayType {
			if roads[edge] == nil {
				roads[edge] = make([]WayType, cap(roads))
			} else {
				roadWayType := roads[edge][i]
				if roadWayType != UnknownWay && roadWayType != currentWay {
					return false
				} else if roadWayType == currentWay {
					continue Upper
				}
			}

			roads[edge][i] = currentWay
		}

		ok := dfs(AL, AL[edge], roadsWithCurrentWayType, roads, currentWay)
		if !ok {
			return false
		}
	}
	return true
}
func main() {
	Railroads(os.Stdin, os.Stdout)
}
