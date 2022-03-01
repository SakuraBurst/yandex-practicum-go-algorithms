package shiftSearch

import (
	"bufio"
	"fmt"
	"io"
)

func ShiftSearch(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}
