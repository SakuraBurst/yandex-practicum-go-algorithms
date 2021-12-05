package maxDepth

import (
	"io"
	"bufio"
	"fmt"
)

func MaxDepth(r io.Reader, w io.Writer){
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}