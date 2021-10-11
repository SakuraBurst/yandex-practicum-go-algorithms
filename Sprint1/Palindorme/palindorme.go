package palindorme

import (
	"bufio"
	"os"
	"strings"
)

func Palindorme() {
	stdReader := bufio.NewReader(os.Stdin)
	stdWriter := bufio.NewWriter(os.Stdout)
	line, _ := stdReader.ReadString('\n')
	mainString := strings.ToLower(strings.TrimSpace(line))
	for i, g := 0, len(mainString)-1; i < g; i, g = i+1, g-1 {
		for !isLetter(mainString[i]) {
			i++
		}
		for !isLetter(mainString[g]) {
			g--
		}
		if mainString[i] != mainString[g] {
			stdWriter.WriteString("False")
			stdWriter.Flush()
			return
		}
	}
	stdWriter.WriteString("True")
	stdWriter.Flush()
}

func isLetter(r byte) bool {
	if r > 'z' || r < 'A' {
		return false
	}
	return true
}
