package borderControl

import (
	"bufio"
	"fmt"
	"io"
)

func BorderControl(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}
