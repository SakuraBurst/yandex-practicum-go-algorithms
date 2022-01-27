package atm

import (
	"bufio"
	"fmt"
	"io"
)

func Atm(r io.Reader, w io.Writer) {
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}
