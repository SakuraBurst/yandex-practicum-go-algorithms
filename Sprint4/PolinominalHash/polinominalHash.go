package polinominalHash

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func PolinominalHash(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	aString, _ := reader.ReadString('\n')
	a, _ := strconv.Atoi(strings.TrimSpace(aString))
	mString, _ := reader.ReadString('\n')
	m, _ := strconv.Atoi(strings.TrimSpace(mString))
	s, _ := reader.ReadString('\n')
	writer := bufio.NewWriter(w)
	writer.WriteString(hash(a, m, strings.TrimSpace(s)))
	writer.Flush()
}

func hash(a int, m int, s string) string {

	len := len(s)
	if len == 0 {
		return "0"
	}
	if len == 1 {
		return strconv.Itoa(int(s[0]) % m)
	}
	result := int((int(s[0])*a)+int(s[1])) % m
	for i := 2; i < len; i++ {
		result = ((result * a) + int(s[i])) % m

	}
	return strconv.Itoa(result)
}

func main() {
	PolinominalHash(os.Stdin, os.Stdout)
}
