package wierdComparison

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func WierdComparison(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	writer := bufio.NewWriter(w)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)
	t, _ := reader.ReadString('\n')
	t = strings.TrimSpace(t)
	if len(s) != len(t) {
		writer.WriteString("NO")
		writer.Flush()
		return
	}
	s_to_t_characterMap := map[string]string{}
	t_to_s_characterMap := map[string]string{}
	for i := 0; i < len(s); i++ {
		sLetter := string(s[i])
		tLetter := string(t[i])
		if s_to_t_characterMap[sLetter] == "" && t_to_s_characterMap[tLetter] == "" {
			s_to_t_characterMap[sLetter] = tLetter
			t_to_s_characterMap[tLetter] = sLetter
		}
		if s_to_t_characterMap[sLetter] != tLetter || t_to_s_characterMap[tLetter] != sLetter {
			writer.WriteString("NO")
			writer.Flush()
			return
		}
	}
	writer.WriteString("YES")
	writer.Flush()
}

func main() {
	WierdComparison(os.Stdin, os.Stdout)
}
