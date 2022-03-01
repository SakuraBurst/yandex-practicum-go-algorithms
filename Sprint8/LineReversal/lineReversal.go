package lineReversal

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func LineReversal(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	fields := strings.Fields(scanner.Text())
	for i, g := 0, len(fields)-1; i < g; i, g = i+1, g-1 {
		fields[i], fields[g] = fields[g], fields[i]
	}
	io.WriteString(w, strings.Join(fields, " "))
}

func main() {
	LineReversal(os.Stdin, os.Stdout)
}
