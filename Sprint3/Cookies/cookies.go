package cookies

import (
	"io"
	"bufio"
	"fmt"
)

func Cookies(r io.Reader, w io.Writer){
	reader := bufio.NewReader(r)
	writer := bufio.NewWriter(w)
	fmt.Println(reader, writer)
}