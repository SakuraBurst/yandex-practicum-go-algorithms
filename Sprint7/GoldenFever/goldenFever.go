package goldenFever

import (
	"bufio"
	"fmt"
	"io"
)

func GoldenFever(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}
