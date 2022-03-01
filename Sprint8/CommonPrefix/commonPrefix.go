package commonPrefix

import (
	"bufio"
	"fmt"
	"io"
)

func CommonPrefix(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}
