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

func dfsOnlyCheck(AL AdjacencyList, vertex Vertex, citiesWithThatWayType []uint16, roads Matrix, currentWay WayType) {
	citiesWithThatWayType = append(citiesWithThatWayType, vertex.value)
Upper:
	for _, edge := range vertex.pointsTo {
		for _, i := range citiesWithThatWayType {
			if roads[edge] == nil {
				roads[edge] = make([]WayType, cap(roads))
			}
			roadValue := roads[edge][i]
			if roadValue == currentWay {
				continue Upper
			}

			roads[edge][i] = currentWay
		}
		dfsOnlyCheck(AL, AL[edge], citiesWithThatWayType, roads, currentWay)
	}
}

func dfs(AL AdjacencyList, vertex Vertex, citiesWithThatWayType []uint16, roads Matrix, currentWay WayType) bool {
	citiesWithThatWayType = append(citiesWithThatWayType, vertex.value)
Upper:
	for _, edge := range vertex.pointsTo {

		for _, i := range citiesWithThatWayType {
			if roads[edge] == nil {
				roads[edge] = make([]WayType, cap(roads))
			} else {
				roadValue := roads[edge][i]
				if roadValue != UnknownWay && roadValue != currentWay {
					return false
				} else if roadValue == currentWay {
					continue Upper
				}
			}

			roads[edge][i] = currentWay
		}

		ok := dfs(AL, AL[edge], citiesWithThatWayType, roads, currentWay)
		if !ok {
			return false
		}
	}
	return true
}
func main() {
	Railroads(os.Stdin, os.Stdout)
}
