package house buying

import (
	"io"
	"bufio"
	"fmt"
)

func House buying(r io.Reader, w io.Writer){
	reader := bufio.NewReader(r)
	writer := bufio.NewWriter(w)
	fmt.Println(reader, writer)
}