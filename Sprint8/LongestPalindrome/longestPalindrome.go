package longestPalindrome

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

func LongestPalindrome(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	s, _ := reader.ReadString('\n')
	s = s[:len(s)-1]
	m := make([]int, 26)
	for _, r := range s {
		m[r-97]++
	}
	palindromeHalf := bytes.Buffer{}
	for i, v := range m {
		if v%2 == 0 {
			for g := 0; g < v/2; g++ {
				palindromeHalf.WriteByte(byte(i + 97))
			}
			m[i] = 0
		} else if v > 1 {
			for g := 0; g < v/2; g++ {
				palindromeHalf.WriteByte(byte(i + 97))
			}
			m[i] = 1
		}
	}
	var lastByte byte
	for i, v := range m {
		if v == 1 {
			lastByte = byte(i + 97)
			break
		}
	}
	io.WriteString(w, createPalindrome(palindromeHalf.Bytes(), lastByte))
}

func createPalindrome(b []byte, lastByte byte) string {
	firstHalf := string(b)

	for i, g := 0, len(b)-1; i < g; i, g = i+1, g-1 {
		b[i], b[g] = b[g], b[i]
	}
	if lastByte != 0 {
		return firstHalf + string(lastByte) + string(b)
	}
	return firstHalf + string(b)
}

func main() {
	LongestPalindrome(os.Stdin, os.Stdout)
}
