package compareTwoStrings

import (
	"bufio"
	"io"
	"os"
	"strings"
	"unicode"
)

func CompareTwoStrings(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	a, _ := reader.ReadString('\n')
	aBulder := strings.Builder{}
	for _, rune := range a {
		if unicode.Is(unicode.White_Space, rune) {
			break
		}
		if rune%2 == 0 {
			aBulder.WriteRune(rune)
		}
	}
	b, _ := reader.ReadString('\n')
	bBuilder := strings.Builder{}
	for _, rune := range b {
		if unicode.Is(unicode.White_Space, rune) {
			break
		}
		if rune%2 == 0 {
			bBuilder.WriteRune(rune)
		}
	}
	io.WriteString(w, compare(aBulder.String(), bBuilder.String()))
}

func compare(a, b string) string {
	if a == b {
		return "0"
	}
	for i := 0; i < len(a); i++ {
		if i == len(b) {
			return "1"
		}
		if a[i] == b[i] {
			continue
		}
		if a[i] < b[i] {
			return "-1"
		} else {
			return "1"
		}
	}
	return "-1"
}

func main() {
	CompareTwoStrings(os.Stdin, os.Stdout)
}
