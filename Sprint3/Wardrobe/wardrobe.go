package wardrobe

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func Wardrobe(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	nStr, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nStr))
	dataStr, _ := reader.ReadString('\n')
	data := strings.Fields(dataStr)
	counter := [3]int{}
	for _, v := range data {
		vInt, _ := strconv.Atoi(v)
		counter[vInt]++
	}
	result := make([]string, n)
	i := 0
	for g, v := range counter {
		for v > 0 {
			result[i] = strconv.Itoa(g)
			v--
			i++
		}
	}
	writer := bufio.NewWriter(w)
	writer.WriteString(strings.Join(result, " "))
	writer.Flush()
}
