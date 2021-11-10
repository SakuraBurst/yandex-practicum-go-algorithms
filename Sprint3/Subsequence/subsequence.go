package subsequence

import (
	"bufio"
	"io"
)

func Subsequence(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	writer := bufio.NewWriter(w)
	subS, _ := reader.ReadString('\n')
	s, _ := reader.ReadString('\n')
	if len(subS) == 0 || len(s) == 0 {
		writer.WriteString("False")
		writer.Flush()
		return
	}
	sIndex := 0
	for _, v := range subS {
		isItemSerched := false
		for i := sIndex; i < len(s); i++ {
			if rune(s[i]) == v {
				sIndex = i
				isItemSerched = true
				break
			}
		}
		if !isItemSerched {
			writer.WriteString("False")
			writer.Flush()
			return
		}
	}
	writer.WriteString("True")
	writer.Flush()
}
