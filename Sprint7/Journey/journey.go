package journey

import (
	"bufio"
	"fmt"
	"io"
)

func Journey(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}
