package leprechaunGold

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type GoldHeap struct {
	weight  int
	canTake []int
}

func LeprechaunGold(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	initDataString, _ := reader.ReadString('\n')
	initData := strings.Fields(initDataString)
	n, _ := strconv.Atoi(initData[0])
	M, _ := strconv.Atoi(initData[1])
	mainDataString, _ := reader.ReadString('\n')
	mainData := strings.Fields(mainDataString)
	previousGoldHeap := GoldHeap{
		weight:  0,
		canTake: make([]int, M+1),
	}
	var goldHeap GoldHeap
	for i := 0; i < n; i++ {
		weight, _ := strconv.Atoi(mainData[i])
		goldHeap = GoldHeap{
			weight:  weight,
			canTake: make([]int, M+1),
		}
		for g := 0; g < M+1; g++ {
			if g-weight >= 0 {
				goldHeap.canTake[g] = max(previousGoldHeap.canTake[g], weight+previousGoldHeap.canTake[g-weight])
			} else {
				goldHeap.canTake[g] = previousGoldHeap.canTake[g]
			}

		}
		previousGoldHeap = goldHeap

	}
	writer := bufio.NewWriter(w)
	writer.WriteString(strconv.Itoa(goldHeap.canTake[M]))
	writer.Flush()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	LeprechaunGold(os.Stdin, os.Stdout)
}
