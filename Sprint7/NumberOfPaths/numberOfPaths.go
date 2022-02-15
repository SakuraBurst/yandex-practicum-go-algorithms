package numberOfPaths

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

const Mod = 1000000007

type Vertex struct {
	value    int
	pointsTo []*Vertex
}

type AdjacencyList []Vertex

func NumberOfPaths(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	initData := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(initData[0])
	m, _ := strconv.Atoi(initData[1])
	AL := make(AdjacencyList, n+1)
	for i := 1; i < n+1; i++ {
		AL[i].value = i
	}
	for i := 0; i < m; i++ {
		scanner.Scan()
		edgeData := strings.Fields(scanner.Text())
		edgeA, _ := strconv.Atoi(edgeData[0])
		edgeB, _ := strconv.Atoi(edgeData[1])
		AL[edgeA].pointsTo = append(AL[edgeA].pointsTo, &AL[edgeB])
	}
	scanner.Scan()
	searchData := strings.Fields(scanner.Text())
	searchFrom, _ := strconv.Atoi(searchData[0])
	searchTo, _ := strconv.Atoi(searchData[1])
	paths := make([]int, n+1)
	visited := make([]bool, n+1)
	paths[searchTo] = 1
	dfs(&AL[searchFrom], visited, paths)
	writer := bufio.NewWriter(w)
	writer.WriteString(strconv.Itoa(paths[searchFrom]))
	writer.Flush()
}

func dfs(vertex *Vertex, visited []bool, paths []int) {
	visited[vertex.value] = true
	for _, v := range vertex.pointsTo {
		if !visited[v.value] {
			dfs(v, visited, paths)
		}
		paths[vertex.value] = (paths[vertex.value] + paths[v.value]) % Mod
	}
}

func main() {
	NumberOfPaths(os.Stdin, os.Stdout)
}
