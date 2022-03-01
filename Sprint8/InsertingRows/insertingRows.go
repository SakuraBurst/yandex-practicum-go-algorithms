package insertingRows

import (
	"bufio"
	"fmt"
	"io"
)

func InsertingRows(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}
