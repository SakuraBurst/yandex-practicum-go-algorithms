package backpack

import (
	"bufio"
	"fmt"
	"io"
)

func Backpack(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}
