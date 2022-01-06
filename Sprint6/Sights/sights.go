package sights

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

type Link struct {
	weight int
	point  *Node
}

type PointsTo []Link

type Node struct {
	value    int
	pointsTo PointsTo
}

type AdjacencyList []*Node

const INFINT = 1 << 62

func Sights(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	initData := strings.Fields(scanner.Text())
	peaks, _ := strconv.Atoi(initData[0])
	edges, _ := strconv.Atoi(initData[1])
	AL := make(AdjacencyList, peaks+1)
	for i := 0; i < peaks; i++ {
		alIndex := i + 1
		AL[alIndex] = &Node{
			value:    alIndex,
			pointsTo: nil,
		}
	}
	for i := 0; i < edges; i++ {
		scanner.Scan()
		edgeData := strings.Fields(scanner.Text())
		edgeA, _ := strconv.Atoi(edgeData[0])
		edgeB, _ := strconv.Atoi(edgeData[1])
		weight, _ := strconv.Atoi(edgeData[2])
		AL[edgeA].pointsTo = append(AL[edgeA].pointsTo, Link{
			weight: weight,
			point:  AL[edgeB],
		})
		AL[edgeB].pointsTo = append(AL[edgeB].pointsTo, Link{
			weight: weight,
			point:  AL[edgeA],
		})
	}
	writer := bufio.NewWriter(w)
	for _, node := range AL {
		if node == nil {
			continue
		}
		visited := make([]bool, peaks+1)
		minWeights := make([]int, peaks+1)

		// первая версия
		// visited[node.value] = true
		//for _, link := range node.pointsTo {
		//	relax(minWeights, link, node)
		//}
		//Dijkstra(AL, visited, minWeights)
		//for i := 0; i < peaks; i++ {
		//	correctIndex := i + 1
		//	if minWeights[correctIndex] == 0 && correctIndex != node.value {
		//		writer.WriteString("-1 ")
		//	} else {
		//		writer.WriteString(strconv.Itoa(minWeights[correctIndex]) + " ")
		//	}
		//}

		// ближе к теории
		for i := 0; i < peaks+1; i++ {
			minWeights[i] = INFINT
		}
		minWeights[node.value] = 0
		Dijkstra2(AL, visited, minWeights)
		for i := 0; i < peaks; i++ {
			correctIndex := i + 1
			if minWeights[correctIndex] == INFINT && correctIndex != node.value {
				writer.WriteString("-1 ")
			} else {
				writer.WriteString(strconv.Itoa(minWeights[correctIndex]) + " ")
			}
		}
		writer.WriteByte('\n')
	}
	writer.Flush()
}

func Dijkstra(AL AdjacencyList, visited []bool, minWeights []int) {
	for node, err := getMinValue(AL, visited, minWeights); err == nil; node, err = getMinValue(AL, visited, minWeights) {
		visited[node.value] = true
		for _, link := range node.pointsTo {
			if !visited[link.point.value] {
				relax(minWeights, link, node)
			}

		}
	}
}

func getMinValue(AL AdjacencyList, visited []bool, minWeights []int) (*Node, error) {
	min := INFINT
	minIndex := -1
	for i, currentMin := range minWeights {
		if !visited[i] && currentMin != 0 && currentMin < min {
			min = currentMin
			minIndex = i
		}
	}
	if minIndex == -1 {
		return nil, errors.New("complete")
	}
	return AL[minIndex], nil
}

func relax(minWeights []int, link Link, node *Node) {
	if minWeights[link.point.value] == 0 || minWeights[link.point.value] > minWeights[node.value]+link.weight {
		minWeights[link.point.value] = minWeights[node.value] + link.weight
	}
}

func Dijkstra2(AL AdjacencyList, visited []bool, minWeights []int) {
	for node, err := getMinValue2(AL, visited, minWeights); err == nil; node, err = getMinValue(AL, visited, minWeights) {
		visited[node.value] = true
		for _, link := range node.pointsTo {
			if !visited[link.point.value] {
				relax2(minWeights, link, node)
			}

		}
	}
}

func getMinValue2(AL AdjacencyList, visited []bool, minWeights []int) (*Node, error) {
	min := INFINT
	minIndex := -1
	for i, currentMin := range minWeights {
		if !visited[i] && currentMin < min {
			min = currentMin
			minIndex = i
		}
	}
	if minIndex == -1 {
		return nil, errors.New("complete")
	}
	return AL[minIndex], nil
}

func relax2(minWeights []int, link Link, node *Node) {
	if minWeights[link.point.value] > minWeights[node.value]+link.weight {
		minWeights[link.point.value] = minWeights[node.value] + link.weight
	}
}
func main() {
	Sights(os.Stdin, os.Stdout)
}
