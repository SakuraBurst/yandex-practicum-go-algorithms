package substrings

import (
	"io"
	"bufio"
	"fmt"
)

func Substrings(r io.Reader, w io.Writer){
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}