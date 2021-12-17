package deleteNode

import (
	"io"
	"bufio"
	"fmt"
)

func DeleteNode(r io.Reader, w io.Writer){
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}