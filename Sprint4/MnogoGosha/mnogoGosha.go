package mnogoGosha

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

const Q = 1000000007
const R = 1 << 32

func MnogoGosha(r io.Reader, w io.Writer) {

	reader := bufio.NewReader(r)
	nAndk, _ := reader.ReadString('\n')
	nAndkFields := strings.Fields(nAndk)
	n, _ := strconv.Atoi(nAndkFields[0])
	k, _ := strconv.Atoi(nAndkFields[1])
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)

	degrees := make([]int, len(s))
	degrees[0] = Q
	for i := 1; i < len(s); i++ {
		degrees[i] = degrees[i-1] * Q % R
	}
	hashMap := map[int][]int{}
	for i := 0; i < len(s)-n; i++ {
		h := hash(degrees, s[i:i+n])
		hashMap[h] = append(hashMap[h], i)
	}
	writer := bufio.NewWriter(w)
	for _, v := range hashMap {
		if len(v) >= k {
			writer.WriteString(strconv.Itoa(v[0]))
			writer.WriteByte(' ')
		}
	}
	writer.Flush()
}

func hash(degrees []int, s string) int {
	h := 0
	for i := 0; i < len(s); i++ {
		h = (h + int(s[i])*degrees[i]) % R
	}
	return h
}

func main() {
	MnogoGosha(os.Stdin, os.Stdout)
}
