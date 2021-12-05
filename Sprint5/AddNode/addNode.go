package addNode

import (
	"io"
	"bufio"
	"fmt"
)

func AddNode(r io.Reader, w io.Writer){
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}