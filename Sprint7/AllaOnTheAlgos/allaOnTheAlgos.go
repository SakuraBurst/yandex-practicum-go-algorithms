package allaOnTheAlgos

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type ATM []int

func AllaOnTheAlgos(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	ATMData := strings.Fields(scanner.Text())
	atm := make(ATM, 0, n)
	for _, datum := range ATMData {
		intData, _ := strconv.Atoi(datum)
		atm = append(atm, intData)
	}
	prevResult := make([]int, x+1)
	result := make([]int, x+1)
	for i := 0; i < n; i++ {

		for g := 0; g < x+1; g++ {
			result[g] = min(result[g], prevResult[g])
			if g-atm[i] >= 0 && (result[g-atm[i]] > 0 || g%atm[i] == 0) {
				result[g] = min(result[g], result[g-atm[i]]+1)
			}
		}
		log.Println(atm[i])
		log.Println(prevResult)
		log.Println(result)
		prevResult = result
		result = make([]int, x+1)
	}
	answer := -1
	if prevResult[x] > 0 {
		answer = prevResult[x]
	}
	writer := bufio.NewWriter(w)
	writer.WriteString(strconv.Itoa(answer))
	writer.Flush()

}

func min(a, b int) int {
	if a == 0 || b == 0 {
		if a > b {
			return a
		}
		return b
	}
	if a < b {
		return a
	}
	return b
}

func main() {
	AllaOnTheAlgos(os.Stdin, os.Stdout)
}
