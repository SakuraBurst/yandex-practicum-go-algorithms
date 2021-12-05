package treeSpliting

import (
	"io"
	"bufio"
	"fmt"
)

func TreeSpliting(r io.Reader, w io.Writer){
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}