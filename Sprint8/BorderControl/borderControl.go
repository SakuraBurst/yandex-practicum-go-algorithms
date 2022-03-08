package borderControl

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func BorderControl(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	max, _ := reader.ReadString('\n')
	max = strings.TrimSpace(max)
	min, _ := reader.ReadString('\n')
	min = strings.TrimSpace(min)
	if len(max) < len(min) {
		max, min = min, max
	}
	if len(max)-len(min) > 2 {
		io.WriteString(w, "FAIL")
		return
	}
	i := 0
	g := 0
	diff := 0
	for i < len(max) && g < len(min) {
		if g == len(min) {
			diff++
			i++
			if diff > 1 {
				io.WriteString(w, "FAIL")
				return
			}
			continue
		}
		if max[i] != min[g] {
			diff++
			if diff > 1 {
				io.WriteString(w, "FAIL")
				return
			}
			if max[i+1] == min[g] {
				i++
			} else if max[i+1] == min[g+1] {
				i++
				g++
			} else {
				io.WriteString(w, "FAIL")
				return
			}
			continue
		}
		i++
		g++
	}
	io.WriteString(w, "OK")
}

func main() {
	BorderControl(os.Stdin, os.Stdout)
}
