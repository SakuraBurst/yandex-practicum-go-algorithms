package dFS

import (
	"io"
	"bufio"
	"fmt"
)

func DFS(r io.Reader, w io.Writer){
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}