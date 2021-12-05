package searchTree

import (
	"io"
	"bufio"
	"fmt"
)

func SearchTree(r io.Reader, w io.Writer){
	reader := bufio.NewReader(r)
	//writer := bufio.NewWriter(w)
	fmt.Println(reader)
}