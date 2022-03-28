package frequentWord

import (
	"bufio"
	"io"
	"os"
	"sort"
	"strconv"
	"unicode"
)

func FrequentWord(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	nS, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(nS[:len(nS)-1])
	countMap := map[string]int{}
	for i := 0; i < n; i++ {
		str, _ := reader.ReadString('\n')
		if unicode.IsSpace(rune(str[len(str)-1])) {
			str = str[:len(str)-1]
		}
		countMap[str]++
	}
	max := -1
	var maxWords []string = nil
	for k, v := range countMap {
		if max == -1 || v > max {
			max = v
			maxWords = []string{k}
			continue
		} else if v == max {
			maxWords = append(maxWords, k)
		}
	}
	sort.Strings(maxWords)
	io.WriteString(w, maxWords[0])
}

func main() {
	FrequentWord(os.Stdin, os.Stdout)
}
