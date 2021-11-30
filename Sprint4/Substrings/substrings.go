package substrings

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func Substrings(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	helperMap := map[string]int{}
	max := 0
	for i := 0; i < len(s); {
		stringRune := string(s[i])
		if _, ok := helperMap[stringRune]; !ok {
			helperMap[stringRune] = i
			i++
		} else {
			minKey, minValue, maxValue := currentMin(helperMap)
			currentMax := maxValue + 1 - minValue
			if currentMax > max {
				max = currentMax
			}
			delete(helperMap, minKey)
		}
	}
	_, minValue, maxValue := currentMin(helperMap)
	currentMax := maxValue + 1 - minValue
	if currentMax > max {
		max = currentMax
	}
	writer := bufio.NewWriter(w)
	writer.WriteString(strconv.Itoa(max))
	writer.Flush()
}

func currentMin(m map[string]int) (minKey string, min, max int) {
	min = -1
	for k, v := range m {
		if min == -1 {
			min = v
		}
		if v <= min {
			min = v
			minKey = k
		}
		if v >= max {
			max = v
		}

	}
	return
}

func main() {
	Substrings(os.Stdin, os.Stdout)
}
