package binaryEx

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func BinaryEx() {
	stdReader := bufio.NewReader(os.Stdin)
	line, _ := stdReader.ReadString('\n')
	mainString := strings.TrimSpace(line)
	line, _ = stdReader.ReadString('\n')
	secondMainString := strings.TrimSpace(line)
	if len(mainString) < len(secondMainString) {
		mainString, secondMainString = secondMainString, mainString
	}
	result := ""
	reminder := false
	for i := len(mainString) - 1; i >= 0; i-- {
		secondStringIndex := len(secondMainString) - len(mainString) + i
		if reminder && secondStringIndex >= 0 {
			switch {
			case mainString[i] == '1' && secondMainString[secondStringIndex] == '1':
				result += "1"
			case mainString[i] == '1' || secondMainString[secondStringIndex] == '1':
				result += "0"
			default:
				result += "1"
				reminder = false
			}
		} else if reminder && secondStringIndex < 0 {
			switch {
			case mainString[i] == '1':
				result += "0"
			default:
				result += "1"
				reminder = false
			}
		} else if !reminder && secondStringIndex >= 0 {
			switch {
			case mainString[i] == '1' && secondMainString[secondStringIndex] == '1':
				result += "0"
				reminder = true
			case mainString[i] == '1' || secondMainString[secondStringIndex] == '1':
				result += "1"
			default:
				result += "0"
			}
		} else {
			result += string(mainString[i])
		}

	}
	if reminder {
		result += "1"
	}
	fmt.Println(string(byteStringReverse([]byte(result))))
}

func byteStringReverse(s []byte) []byte {
	for i, g := 0, len(s)-1; i < g; i, g = i+1, g-1 {
		s[i], s[g] = s[g], s[i]
	}
	return s
}
