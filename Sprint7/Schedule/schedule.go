package schedule

import (
	"bufio"
	"fmt"
	"io"
)

func Schedule(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}
