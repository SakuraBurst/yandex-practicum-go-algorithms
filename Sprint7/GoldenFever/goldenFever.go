package goldenFever

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type GoldenHeap struct {
	price          int
	numberOfChunks int
}

type GoldenMine []GoldenHeap

func GoldenFever(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	limit, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())
	goldenMine := make(GoldenMine, 0, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		heapData := strings.Fields(scanner.Text())
		price, _ := strconv.Atoi(heapData[0])
		numberOfChunks, _ := strconv.Atoi(heapData[1])
		goldenMine = append(goldenMine, GoldenHeap{
			price:          price,
			numberOfChunks: numberOfChunks,
		})
	}
	goldenMine = merge(goldenMine)
	result := 0
	for i := 0; i < N && limit > 0; i++ {
		for g := 0; g < goldenMine[i].numberOfChunks && limit > 0; g++ {
			result += goldenMine[i].price
			limit--
		}
	}
	writer := bufio.NewWriter(w)
	writer.WriteString(strconv.Itoa(result))
	writer.Flush()
}

func merge(mine GoldenMine) GoldenMine {
	mineLength := len(mine)
	if mineLength < 2 {
		return mine
	}
	mid := mineLength >> 1
	a := merge(mine[:mid])
	aLength := len(a)
	b := merge(mine[mid:])
	bLength := len(b)
	aIndex := 0
	bIndex := 0
	result := make(GoldenMine, 0, mineLength)
	for aIndex < aLength || bIndex < bLength {
		if aIndex == aLength {
			result = append(result, b[bIndex])
			bIndex++
			continue
		}
		if bIndex == bLength {
			result = append(result, a[aIndex])
			aIndex++
			continue
		}
		if a[aIndex].price >= b[bIndex].price {
			result = append(result, a[aIndex])
			aIndex++
		} else {
			result = append(result, b[bIndex])
			bIndex++
		}
	}
	return result

}

func main() {
	GoldenFever(os.Stdin, os.Stdout)
}
