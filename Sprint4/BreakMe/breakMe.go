package breakMe

import (
	"io"
	"bufio"
	"fmt"
)

func BreakMe(r io.Reader, w io.Writer){
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}