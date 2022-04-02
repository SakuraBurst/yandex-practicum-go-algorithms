package packedPrefix

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func PackedPrefix(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	prefix := ""
	currentString := ""
	for i := 0; i < n; i++ {
		scanner.Scan()
		if i == 0 {
			prefix = unpack(scanner.Text())
			continue
		}
		currentString = fixedLengthUnpack(scanner.Text(), len(prefix), 0)
		prefix = currentString[:firstNotEqualsElementIndex(prefix, currentString)]

	}
	io.WriteString(w, prefix)
}

func firstNotEqualsElementIndex(a, b string) int {
	i := 0
	max := 0
	if len(a) > len(b) {
		max = len(a)
	} else {
		max = len(b)
	}
	for ; i < max; i++ {
		if i >= len(a) || i >= len(b) || a[i] != b[i] {
			break
		}
	}
	return i
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

func fixedLengthUnpack(s string, neededLength int, currentLength int) string {
	builder := strings.Builder{}
	for i := 0; i < len(s); {
		length := builder.Len() + currentLength
		if length >= neededLength {
			return cutString(builder.String(), length-neededLength)
		}
		if unicode.IsDigit(rune(s[i])) {
			howMuch, _ := strconv.Atoi(string(s[i]))
			i += 2
			end := searchForEndOfPackedSequense(s, i)
			unpackedS := fixedLengthUnpack(s[i:end], neededLength, builder.Len()+currentLength)
			for g := 0; g < howMuch; g++ {
				length := builder.Len() + currentLength
				if length >= neededLength {
					return cutString(builder.String(), length-neededLength)
				}
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

func cutString(s string, cutTo int) string {
	return s[:len(s)-cutTo]
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

func main() {
	PackedPrefix(os.Stdin, os.Stdout)
}
