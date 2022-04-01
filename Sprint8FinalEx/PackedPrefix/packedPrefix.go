package packedPrefix

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func PackedPrefix(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < n; i++ {
		scanner.Scan()
		log.Println(unpack(scanner.Text()))
	}
}

func unpack(s string) string {
	builder := strings.Builder{}
	for i := 0; i < len(s); {
		if unicode.IsDigit(rune(s[i])) {
			howMuch, _ := strconv.Atoi(string(s[i]))
			i += 2
			end := searchForEndOfPackedSequense(s, i)
			unpackedS := unpack(s[i:end])
			for g := 0; g < howMuch; g++ {
				builder.WriteString(unpackedS)
			}
			i = end + 1
		} else {
			builder.WriteByte(s[i])
			i++
		}

	}
	return builder.String()
}

func searchForEndOfPackedSequense(s string, start int) int {
	bracketCount := 1
	g := start
	for ; g < len(s); g++ {
		if s[g] == '[' {
			bracketCount++
			continue
		}
		if s[g] == ']' {
			bracketCount--
			if bracketCount == 0 {
				break
			}
		}
	}
	return g
}
