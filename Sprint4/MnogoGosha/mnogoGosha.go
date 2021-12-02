package mnogoGosha

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

const Q = 1000000009
const R = 1 << 32

func MnogoGosha(r io.Reader, w io.Writer) {

	reader := bufio.NewReader(r)
	nAndk, _ := reader.ReadString('\n')
	nAndkFields := strings.Fields(nAndk)
	n, _ := strconv.Atoi(nAndkFields[0])
	k, _ := strconv.Atoi(nAndkFields[1])
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	windowSlideDegree := slideDegreeOfQFModR(n)
	hashes := hashAllSubstrings(s)
	hashMap := map[int][]int{}
	for i := 0; i < len(s)+1-n; i++ {
		start := i
		end := i + n
		h := (hashes[end] - hashes[start]*windowSlideDegree) % R
		if h < 0 {
			h += R
			h %= R
		}
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

func slideDegreeOfQFModR(sw int) int {
	degrees := make([]int, sw+1)
	degrees[0] = 0
	degrees[1] = Q
	for i := 2; i < sw+1; i++ {
		degrees[i] = degrees[i-1] * Q % R
	}
	return degrees[sw]
}

func hashAllSubstrings(s string) []int {
	hashes := make([]int, 0, len(s)+1)
	hashes = append(hashes, 0)
	for i := 0; i < len(s); i++ {
		hashes = append(hashes, (hashes[i]*Q+int(s[i]))%R)
	}
	return hashes
}

func main() {
	MnogoGosha(os.Stdin, os.Stdout)
}
