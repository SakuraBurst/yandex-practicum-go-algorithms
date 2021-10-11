package longestWord

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func LongestWord() {
	var n int
	fmt.Scan(&n)
	stdReader := bufio.NewReader(os.Stdin)
	stdReader.ReadByte()
	str, _ := stdReader.ReadString('\n')
	mainSliceOfStrings := strings.Split(strings.TrimSpace(str), " ")
	currentLongestWord := ""
	for _, v := range mainSliceOfStrings {
		if len(v) > len(currentLongestWord) {
			currentLongestWord = v
		}
	}
	stdWriter := bufio.NewWriter(os.Stdout)
	stdWriter.WriteString(currentLongestWord + "\n")
	stdWriter.Flush()
	fmt.Fprint(stdWriter, lengthOfWord(currentLongestWord))
	stdWriter.Flush()
}

func lengthOfWordByte(s []byte) int {
	i := 0
	counter := 0
	for i < len(s) {
		_, s := utf8.DecodeRune(s[i:])
		i += s
		counter++
	}
	return counter
}
func lengthOfWord(s string) int {
	counter := 0
	for range s {
		counter++
	}
	return counter
}
