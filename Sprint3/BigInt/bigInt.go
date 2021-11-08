package bigInt

import (
	"io"
	"bufio"
	"fmt"
)

func BigInt(r io.Reader, w io.Writer){
	reader := bufio.NewReader(r)
	writer := bufio.NewWriter(w)
	fmt.Println(reader, writer)
}