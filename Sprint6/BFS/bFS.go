package bFS

import (
	"io"
	"bufio"
	"fmt"
)

func BFS(r io.Reader, w io.Writer){
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}