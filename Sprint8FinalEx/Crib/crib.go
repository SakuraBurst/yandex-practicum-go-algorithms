package crib

import (
	"bufio"
	"fmt"
	"io"
)

func Crib(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}
