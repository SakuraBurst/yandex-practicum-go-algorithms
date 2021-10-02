package sum2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Sum2Hash() {
	stdReader := bufio.NewReader(os.Stdin)
	stdWriter := bufio.NewWriter(os.Stdout)
	nLine, _ := stdReader.ReadString('\n')
	n, _ := strconv.ParseInt(strings.TrimSpace(nLine), 10, 64)
	arrayLine, _ := stdReader.ReadString('\n')
	splited := strings.Split(strings.TrimSpace(arrayLine), " ")
	sLine, _ := stdReader.ReadString('\n')
	s, _ := strconv.ParseInt(strings.TrimSpace(sLine), 10, 64)
	hash := map[int64]int64{}
	var i int64
	for i = 0; i < n; i++ {
		item, _ := strconv.ParseInt(strings.TrimSpace(splited[i]), 10, 64)
		_, result := hash[s-item]
		if result {
			fmt.Fprintf(stdWriter, "%d %d", hash[s-item], item)
			break
		}

		hash[item] = item

	}
	if stdWriter.Buffered() == 0 {
		fmt.Fprint(stdWriter, "None")
	}
	stdWriter.Flush()
}
